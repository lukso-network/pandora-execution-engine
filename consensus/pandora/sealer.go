package pandora

import (
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/types"
)

func (pan *Pandora) Seal(chain consensus.ChainHeaderReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error {
	// it does nothing. It just send sealing info to pandora running loop
	pan.newSealRequestCh <- &sealTask{block: block, results: results}
	return nil
}
