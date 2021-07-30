package pandora

import (
	"math/big"
	"math/bits"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/pkg/errors"
	common2 "github.com/silesiacoin/bls/common"
)

// copyEpochInfo
func copyEpochInfo(ei *EpochInfo) *EpochInfo {
	cpy := *ei
	return &cpy
}

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

func (pandoraExtraDataSealed *ExtraDataWithBLSSig) FromHeader(header *types.Header) {
	err := rlp.DecodeBytes(header.Extra, pandoraExtraDataSealed)

	if nil != err {
		panic(err.Error())
	}
}

func (pandoraExtraDataSealed *ExtraDataWithBLSSig) FromExtraDataAndSignature(
	pandoraExtraData ExtraData,
	signature common2.Signature,
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
