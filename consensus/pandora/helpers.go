package pandora

import (
	"math/big"
	"math/bits"

	"github.com/status-im/keycard-go/hexutils"

	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/pkg/errors"
	bls_common "github.com/prysmaticlabs/prysm/shared/bls/common"
	"github.com/prysmaticlabs/prysm/shared/bls/herumi"
)

// extraDataWithoutBLSSig
func extraDataWithoutBLSSig(rlpExtraData []byte) (*ExtraData, error) {
	extraData := new(ExtraData)
	if err := rlp.DecodeBytes(rlpExtraData, extraData); err != nil {
		return nil, err
	}
	return extraData, nil
}

// prepareShardingInfo
func prepareShardingInfo(header *types.Header, sealHash common.Hash) [4]string {
	var shardingInfo [4]string
	rlpHeader, _ := rlp.EncodeToBytes(header)

	shardingInfo[0] = sealHash.Hex()
	shardingInfo[1] = header.ReceiptHash.Hex()
	shardingInfo[2] = hexutil.Encode(rlpHeader)
	shardingInfo[3] = hexutil.Encode(header.Number.Bytes())

	return shardingInfo
}

// getDummyHeader
func getDummyHeader() *types.Header {
	return &types.Header{
		ParentHash:  common.HexToHash("3244474eb97faefc26df91a8c3d0f2a8f859855ba87b76b1cc6044cca29add40"),
		UncleHash:   common.HexToHash("1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
		Coinbase:    common.HexToAddress("b46d14ef42ac9bb01303ba1842ea784e2460c7e7"),
		Root:        common.HexToHash("03906b0760f3bec421d8a71c44273a5994c5f0e35b8b8d9e2112dc95a182aae6"),
		TxHash:      common.HexToHash("56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"),
		ReceiptHash: common.HexToHash("56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"),
		Difficulty:  big.NewInt(1),
		Number:      big.NewInt(4),
		GasLimit:    4718380,
		GasUsed:     0,
		Time:        1626171089,
		Extra:       common.FromHex("0xf866c30a800ab860a899054e1dd5ada5f5174edc532ffa39662cbfc90470233028096d7e41a3263114572cb7d0493ba213becec37f43145d041e0bfbaaf4bf8c2a7aeaebdd0d7fd6c326831b986a9802bf5e9ad1f180553ae0af77334cd4eb606ed71b0dc7db424e"),
		MixDigest:   common.HexToHash("a899054e1dd5ada5f5174edc532ffa39662cbfc90470233028096d7e41a32631"),
		Nonce:       types.BlockNonce{0x0000000000000000},
	}
}

func Mul64(a, b uint64) (uint64, error) {
	overflows, val := bits.Mul64(a, b)
	if overflows > 0 {
		return 0, errors.New("multiplication overflows")
	}
	return val, nil
}

func (p *Pandora) StartSlot(epoch uint64) (uint64, error) {
	slot, err := Mul64(p.config.SlotsPerEpoch, epoch)
	if err != nil {
		return slot, errors.Errorf("start slot calculation overflows: %v", err)
	}
	return slot, nil
}

func (pandoraExtraDataSealed *ExtraDataSealed) FromHeader(header *types.Header) {
	err := rlp.DecodeBytes(header.Extra, pandoraExtraDataSealed)

	if nil != err {
		panic(err.Error())
	}
}

func (pandoraExtraDataSealed *ExtraDataSealed) FromExtraDataAndSignature(
	pandoraExtraData ExtraData,
	signature bls_common.Signature,
) {
	var blsSignatureBytes BlsSignatureBytes
	signatureBytes := signature.Marshal()

	if len(signatureBytes) != signatureSize {
		panic("Incompatible bls mode detected")
	}

	copy(blsSignatureBytes[:], signatureBytes[:])
	pandoraExtraDataSealed.ExtraData = pandoraExtraData
	pandoraExtraDataSealed.BlsSignatureBytes = &blsSignatureBytes
}

func (p *Pandora) verifyHeaderWorker(chain consensus.ChainHeaderReader, headers []*types.Header, index int) error {
	var parent *types.Header
	if index == 0 {
		parent = chain.GetHeader(headers[0].ParentHash, headers[0].Number.Uint64()-1)
	} else if headers[index-1].Hash() == headers[index].ParentHash {
		parent = headers[index-1]
	}
	if parent == nil {
		return consensus.ErrUnknownAncestor
	}
	return p.verifyHeader(chain, headers[index], parent)
}

func (p *Pandora) verifyHeader(chain consensus.ChainHeaderReader, header, parent *types.Header) error {
	if header.Time <= parent.Time {
		return errOlderBlockTime
	}
	// Verify the block's difficulty based on its timestamp and parent's difficulty
	expected := p.CalcDifficulty(chain, header.Time, parent)
	if expected.Cmp(header.Difficulty) != 0 {
		return fmt.Errorf("invalid difficulty: have %v, want %v", header.Difficulty, expected)
	}
	// Verify that the gas limit is <= 2^63-1
	cap := uint64(0x7fffffffffffffff)
	if header.GasLimit > cap {
		return fmt.Errorf("invalid gasLimit: have %v, max %v", header.GasLimit, cap)
	}
	// Verify that the gasUsed is <= gasLimit
	if header.GasUsed > header.GasLimit {
		return fmt.Errorf("invalid gasUsed: have %d, gasLimit %d", header.GasUsed, header.GasLimit)
	}

	// Verify that the gas limit remains within allowed bounds
	diff := int64(parent.GasLimit) - int64(header.GasLimit)
	if diff < 0 {
		diff *= -1
	}
	limit := parent.GasLimit / params.GasLimitBoundDivisor

	if uint64(diff) >= limit || header.GasLimit < params.MinGasLimit {
		return fmt.Errorf("invalid gas limit: have %d, want %d += %d", header.GasLimit, parent.GasLimit, limit)
	}
	// Verify that the block number is parent's +1
	if diff := new(big.Int).Sub(header.Number, parent.Number); diff.Cmp(big.NewInt(1)) != 0 {
		return consensus.ErrInvalidNumber
	}

	// verify bls signature
	//if err := p.VerifyBLSSignature(header); err != nil {
	//	return err
	//}

	return nil
}

func (p *Pandora) VerifyBLSSignature(header *types.Header) error {
	// decode the extraData byte
	extraDataWithBLSSig := new(ExtraDataSealed)
	if err := rlp.DecodeBytes(header.Extra, extraDataWithBLSSig); err != nil {
		log.Error("Failed to decode extraData with signature", "err", err)
		return err
	}
	// extract the extraData
	extractedSlot := extraDataWithBLSSig.Slot
	extractedEpoch := extraDataWithBLSSig.Epoch
	extractedIndex := extraDataWithBLSSig.Turn

	//curEpochInfo, err := p.epochInfoCache.get(extractedEpoch)
	curEpochInfo := p.getEpochInfo(extractedEpoch)
	if curEpochInfo == nil {
		log.Error("Epoch info not found in cache", "slot", extractedSlot, "epoch", extractedEpoch)
		return consensus.ErrEpochNotFound
	}

	blsSignatureBytes := extraDataWithBLSSig.BlsSignatureBytes
	log.Debug("Incoming header's extraData info", "slot", extractedSlot, "epoch",
		extractedEpoch, "turn", extractedIndex, "blsSig", common.Bytes2Hex(blsSignatureBytes[:]))

	signature, err := herumi.SignatureFromBytes(blsSignatureBytes[:])
	if err != nil {
		log.Error("Failed retrieve signature from extraData", "err", err)
		return err
	}
	validatorPubKey := curEpochInfo.ValidatorList[extractedIndex]
	sealHash := p.SealHash(header)
	log.Debug("In verifyBlsSignature", "header extra data", common.Bytes2Hex(header.Extra), "header block Number", header.Number.Uint64(), "sealHash", sealHash, "sealHash (signature msg) in bytes", sealHash[:], "validatorPublicKey", hexutils.BytesToHex(validatorPubKey.Marshal()), "extractedIndex", extractedIndex)

	if !signature.Verify(validatorPubKey, sealHash[:]) {
		sealHashffff := hexutil.Encode(sealHash[:])
		publicKeyBytes := hexutil.Encode(validatorPubKey.Marshal())
		log.Error("Failed to verify bls signature", "err", errSigFailedToVerify, "asss", sealHashffff, "pub", publicKeyBytes)
		return errSigFailedToVerify
	}
	return nil
}

// getEpochInfo
func (p *Pandora) getEpochInfo(epoch uint64) *EpochInfo {
	p.epochInfosMu.RLock()
	defer p.epochInfosMu.RUnlock()

	return p.epochInfos[epoch]
}

// setEpochInfo
func (p *Pandora) setEpochInfo(epoch uint64, epochInfo *EpochInfo) {
	p.epochInfosMu.Lock()
	defer p.epochInfosMu.Unlock()

	log.Debug("store new epoch info into map", "epoch", epoch, "epochInfo", fmt.Sprintf("%+v", epochInfo))
	p.epochInfos[epoch] = epochInfo
}
