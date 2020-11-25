package validatorset

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
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
	for key, value := range multi.sets {
		if big.NewInt(int64(key)).Cmp(parentBlockNum) >= 0 {
			log.Trace("Multi ValidatorSet retrieved for block", "blockHeight", key)
			return value, key
		}
	}
	log.Error("constructor validation ensures that there is at least one validator set for block 0; block 0 is less than any uint;")
	return nil, 0
}

func (multi *Multi) SignalToChange(first bool, logs []*types.Log, header *types.Header) bool {
	validator, setBlockNumber := multi.correctSet(header.Number)
	log.Debug("getting a validator set", "type", reflect.TypeOf(validator))

	first = big.NewInt(int64(setBlockNumber)).Cmp(header.Number) == 0
	return validator.SignalToChange(first, logs, header)
}

func (multi *Multi) FinalizeChange(header *types.Header, state *state.StateDB) error {
	validator, _ := multi.correctSet(header.Number)
	log.Debug("getting a validator set", "type", reflect.TypeOf(validator))

	return validator.FinalizeChange(header, state)
}

func (multi *Multi) GetValidatorsByCaller(header *types.Header) []common.Address {
	validator, _ := multi.correctSet(header.Number)
	log.Debug("getting a validator set", "type", reflect.TypeOf(validator))

	return validator.GetValidatorsByCaller(header)
}

func (multi *Multi) CountValidators() {
	panic("implement me")
}