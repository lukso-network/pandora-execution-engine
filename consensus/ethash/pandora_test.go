package ethash

import (
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
)

func Test_GetMinimalConsensus(t *testing.T) {
	ethash := &Ethash{
		mci: newlru("epochSet", 12, NewMinimalConsensusInfo),
	}
	defer func(ethash *Ethash) {
		err := ethash.Close()
		if err != nil {
			t.Error("error while closing ethash", "error", err)
		}
	}(ethash)
	genesisEpoch := NewMinimalConsensusInfo(0).(*MinimalEpochConsensusInfo)
	genesisEpoch.EpochTimeStart = time.Unix(1616008343, 0)
	ethash.mci.cache.Add(0, genesisEpoch)

	header := &types.Header{
		Time: uint64(1616008390),
	}
	_, err := ethash.getMinimalConsensus(header)
	if err != nil {
		t.Error("error while getting minimal consensus", "error", err)
	}
}
