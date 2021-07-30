package pandora

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/log"
)

var (
	errPandoraStopped = errors.New("pandora stopped")
)

// API is a user facing RPC API to allow controlling the signer and voting
// mechanisms of the proof-of-authority scheme.
type API struct {
	ctx     context.Context
	cancel  context.CancelFunc
	chain   consensus.ChainHeaderReader
	pandora *Pandora
}

// GetShardingWork returns a work package for external miner.
func (api *API) GetShardingWork(parentHash common.Hash, blockNumber uint64, slotNumber uint64, epoch uint64) ([4]string, error) {
	emptyRes := [4]string{}
	if api.pandora == nil {
		return [4]string{}, errors.New("not supported")
	}

	var (
		shardingInfoCh = make(chan [4]string)
		errorCh        = make(chan error)
	)
	select {
	case api.pandora.fetchShardingInfoCh <- &shardingInfoReq{errc: errorCh, res: shardingInfoCh, slot: slotNumber, epoch: epoch}:
		log.Debug("Try to fetch current header")
	case <-api.ctx.Done():
		return emptyRes, errPandoraStopped
	}
	select {
	case shardingInfo := <-shardingInfoCh:
		log.Debug("Sending current sharding info to validator", "shardingInfo", fmt.Sprintf("%+v", shardingInfo))
		return shardingInfo, nil
	case err := <-errorCh:
		return emptyRes, err
	}
}

// SubmitWorkBLS can be used by external miner to submit their POS solution.
// It returns an indication if the work was accepted.
// Note either an invalid solution, a stale work a non-existent work will return false.
// This submit work contains BLS storing feature.
func (api *API) SubmitWorkBLS(nonce types.BlockNonce, hash common.Hash, hexSignatureString string) bool {
	if api.pandora == nil {
		return false
	}

	signatureBytes := hexutil.MustDecode(hexSignatureString)
	blsSignatureBytes := new(BlsSignatureBytes)
	copy(blsSignatureBytes[:], signatureBytes[:])

	var errc = make(chan error, 1)

	select {
	case api.pandora.submitShardingInfoCh <- &shardingResult{
		nonce:   nonce,
		hash:    hash,
		blsSeal: blsSignatureBytes,
		errc:    errc,
	}:
	case <-api.pandora.ctx.Done():
		return false
	}
	err := <-errc
	return err == nil
}
