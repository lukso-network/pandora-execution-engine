package pandora

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/silesiacoin/bls/herumi"
)

const (
	// staleThreshold is the maximum depth of the acceptable stale but valid ethash solution.
	staleThreshold = 7
)

func (pan *Pandora) Seal(chain consensus.ChainHeaderReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error {
	// it does nothing. It just send sealing info to pandora running loop
	pan.newSealRequestCh <- &sealTask{block: block, results: results}
	return nil
}

func (pan *Pandora) submitWork(nonce types.BlockNonce, sealHash common.Hash, blsSignatureBytes *BlsSignatureBytes) bool {
	if pan.currentBlock == nil {
		log.Error("No block found while submitting work", "sealhash", sealHash)
		return false
	}

	// Make sure the work submitted is present
	block := pan.works[sealHash]
	if block == nil {
		log.Warn("Work submitted but none pending", "sealHash", sealHash,
			"blockNumber", pan.currentBlock.NumberU64())
		return false
	}
	// Verify the correctness of submitted result.
	header := block.Header()
	extraDataWithSignature := new(ExtraDataSealed)
	blsSignature, err := herumi.SignatureFromBytes(blsSignatureBytes[:])
	if nil != err {
		log.Error("error while forming signature from bytes", "err", err,
			"methodName", "Seal", "blockNumber", header.Number)
		return false
	}

	pandoraExtraData := new(ExtraData)
	err = rlp.DecodeBytes(header.Extra, pandoraExtraData)
	if nil != err {
		log.Error("rlp decode failed while converting pandora Extra data", "error", err,
			"blockNumber", header.Number)
		return false
	}

	extraDataWithSignature.FromExtraDataAndSignature(*pandoraExtraData, blsSignature)
	header.Extra, err = rlp.EncodeToBytes(extraDataWithSignature)
	if nil != err {
		log.Error("Invalid extraData in header", "sealHash", sealHash, "err", err,
			"slot", pandoraExtraData.Slot, "blockNumber", header.Number)
		return false
	}

	start := time.Now()

	if err := pan.verifyBLSSignature(header); err != nil {
		log.Warn("Invalid bls signature submitted from validator",
			"sealHash", sealHash, "elapsed", common.PrettyDuration(time.Since(start)),
			"err", err, "slot", pandoraExtraData.Slot, "blockNumber", header.Number)
		return false
	}

	// Make sure the result channel is assigned.
	if pan.results == nil {
		log.Error("Pandora result channel is empty, submitted mining result is rejected")
		return false
	}
	log.Debug("Verified correct sharding info", "sealHash", sealHash,
		"elapsed", common.PrettyDuration(time.Since(start)), "slot", pandoraExtraData.Slot, "blockNumber", header.Number)

	// Solutions seems to be valid, return to the miner and notify acceptance.
	solution := block.WithSeal(header)

	// The submitted solution is within the scope of acceptance.
	if solution.NumberU64()+staleThreshold > pan.currentBlock.NumberU64() {
		select {
		case pan.results <- solution:
			log.Debug("Sharding block submitted is acceptable", "number", solution.NumberU64(),
				"sealHash", sealHash, "hash", solution.Hash(), "slot", pandoraExtraData.Slot)
			return true
		default:
			log.Warn("Sealing result is not read by worker", "mode", "remote", "sealHash", sealHash)
			return false
		}
	}
	// The submitted block is too old to accept, drop it.
	log.Warn("Sharding block submitted is too old", "number", solution.NumberU64(),
		"sealHash", sealHash, "hash", solution.Hash(), "slot", pandoraExtraData.Slot)
	return false
}
