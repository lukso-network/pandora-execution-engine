package pandora

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
)

func (pan *Pandora) Author(header *types.Header) (common.Address, error) {
	return header.Coinbase, nil
}

// SealHash returns the hash of a block prior to it being sealed.
func (p *Pandora) SealHash(header *types.Header) (hash common.Hash) {
	hasher := sha3.NewLegacyKeccak256()

	extraData := header.Extra
	extraDataLen := len(extraData)

	// Bls signature is 96 bytes long and will be inserted at the bottom of the extraData field
	if extraDataLen > signatureSize {
		//extraData = extraData[:extraDataLen-signatureSize]
		pandoraExtraData := new(ExtraDataWithBLSSig)
		pandoraExtraData.FromHeader(header)
		headerExtra := new(ExtraData)
		headerExtra.Epoch = pandoraExtraData.Epoch
		headerExtra.Turn = pandoraExtraData.Turn
		headerExtra.Slot = pandoraExtraData.Slot
		extraData, _ = rlp.EncodeToBytes(headerExtra)
	}

	rlp.Encode(hasher, []interface{}{
		header.ParentHash,
		header.UncleHash,
		header.Coinbase,
		header.Root,
		header.TxHash,
		header.ReceiptHash,
		header.Bloom,
		header.Difficulty,
		header.Number,
		header.GasLimit,
		header.GasUsed,
		header.Time,
		extraData,
	})
	hasher.Sum(hash[:0])
	return hash
}
