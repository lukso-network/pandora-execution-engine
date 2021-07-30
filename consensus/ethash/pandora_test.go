package ethash

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/assert"
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

func Test_Ethash_IsInRecentSlot(t *testing.T) {
	ethash := &Ethash{
		mci: newlru("epochSet", 12, NewMinimalConsensusInfo),
		config: Config{
			// In pandora-vanguard implementation we do not need to increase nonce and mixHash is sealed/calculated on the Vanguard side
			PowMode: ModePandora,
			Log:     log.Root(),
		},
	}

	defer func() {
		MockedTimeNow = 0
	}()

	t.Run("should return error", func(t *testing.T) {
		header := &types.Header{Time: uint64(time.Now().Unix())}
		err, _ := ethash.IsInRecentSlot(header)
		assert.Error(t, err)
	})

	genesisTime := time.Now()
	MockedTimeNow = uint64(genesisTime.Unix())

	genesisInfo := NewMinimalConsensusInfo(0).(*MinimalEpochConsensusInfo)
	genesisInfo.AssignEpochStartFromGenesis(genesisTime)
	err := ethash.InsertMinimalConsensusInfo(0, genesisInfo)
	assert.NoError(t, err)

	t.Run("should return true when using time.Now()", func(t *testing.T) {
		header := &types.Header{Time: MockedTimeNow}
		currentErr, inRecentSlot := ethash.IsInRecentSlot(header)
		assert.NoError(t, currentErr)
		assert.True(t, inRecentSlot)
	})

	t.Run("should return false if slot is higher", func(t *testing.T) {
		header := &types.Header{Time: MockedTimeNow + SlotTimeDuration*5}
		currentErr, inRecentSlot := ethash.IsInRecentSlot(header)
		assert.NoError(t, currentErr)
		assert.False(t, inRecentSlot)
	})

	t.Run("should return false if slot is lower", func(t *testing.T) {
		header := &types.Header{Time: MockedTimeNow - SlotTimeDuration*5}
		currentErr, inRecentSlot := ethash.IsInRecentSlot(header)
		assert.NoError(t, currentErr)
		assert.False(t, inRecentSlot)
	})

	t.Run("should return true if slot is higher than genesis", func(t *testing.T) {
		MockedTimeNow = MockedTimeNow + SlotTimeDuration*3
		header := &types.Header{Time: MockedTimeNow}
		currentErr, inRecentSlot := ethash.IsInRecentSlot(header)
		assert.NoError(t, currentErr)
		assert.True(t, inRecentSlot)
	})
}
