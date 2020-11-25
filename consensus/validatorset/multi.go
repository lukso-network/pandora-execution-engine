package validatorset

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"reflect"
)

type Multi struct {
	sets map[uint64]ValidatorSet
}


func NewMulti(setMap map[uint64]ValidatorSet) *Multi {
	return &Multi{
		sets: setMap,
	}
}

func (multi *Multi) correctSet(parentBlockNum *big.Int) (ValidatorSet, uint64) {
	prevBlockNumber := uint64(0)
	prevValidator := multi.sets[0]

	for key, value := range multi.sets {
		if big.NewInt(int64(key)).Cmp(parentBlockNum) >= 0 {
			log.Trace("Multi ValidatorSet retrieved for block", "blockHeight", key)
			break
		}
		prevBlockNumber = key
		prevValidator = value
	}
	log.Error("constructor validation ensures that there is at least one validator set for block 0; block 0 is less than any uint;")
	return prevValidator, prevBlockNumber
}

func (multi *Multi) SignalToChange(first bool, logs []*types.Log, header *types.Header) ([]common.Address, bool) {
	validator, setBlockNumber := multi.correctSet(header.Number)
	log.Debug("getting a validator set", "type", reflect.TypeOf(validator))

	first = big.NewInt(int64(setBlockNumber)).Cmp(header.Number) == 0
	return validator.SignalToChange(first, logs, header)
}

func (multi *Multi) FinalizeChange(header *types.Header, state *state.StateDB) error {
	validator, _ := multi.correctSet(header.Number)
	return validator.FinalizeChange(header, state)
}

func (multi *Multi) GetValidatorsByCaller(blockNumber *big.Int) []common.Address {
	validator, _ := multi.correctSet(blockNumber)
	log.Debug("validator set", "type", reflect.TypeOf(validator))
	return validator.GetValidatorsByCaller(blockNumber)
}

func (multi *Multi) CountValidators() int {
	panic("implement me")
}

func (multi *Multi) PrepareBackend(header *types.Header, chain *core.BlockChain, chainDb ethdb.Database) error {
	validator, _ := multi.correctSet(header.Number)
	return validator.PrepareBackend(header, chain, chainDb)
}