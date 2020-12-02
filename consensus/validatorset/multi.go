package validatorset

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"sort"
)

type Multi struct {
	sets 			map[int]ValidatorSet
	blockNumbers	[]int
}


func NewMulti(setMap map[int]ValidatorSet) *Multi {
	setNumbers := make([]int, 0)
	for key, _ := range setMap {
		setNumbers = append(setNumbers, key)
	}
	sort.Slice(setNumbers, func(i, j int) bool {
		return setNumbers[i] > setNumbers[j]
	})

	return &Multi{
		sets: setMap,
		blockNumbers: setNumbers,
	}
}

// get correct set by block number, along with block number at which
// this set was activated.
func (multi *Multi) correctSet(blockNumber *big.Int) (ValidatorSet, int64) {
	if len(multi.sets) == 0 {
		panic("constructor validation ensures that there is at least one validator set for block 0")
	}
	setNum := 0
	for _, setNumber := range multi.blockNumbers {
		if blockNumber.Cmp(big.NewInt(int64(setNumber))) >= 0 {
			setNum = setNumber
			break
		}
	}
	log.Debug("Multi ValidatorSet retrieved for block", "blockHeight", setNum)
	return multi.sets[setNum], int64(setNum)
}

func (multi *Multi) SignalToChange(first bool, receipts types.Receipts, blockNumber int64, simulatedBackend bind.ContractBackend) ([]common.Address, bool, bool) {
	validator, setBlockNumber := multi.correctSet(big.NewInt(blockNumber))

	// block number is same as set block height in which transition happens
	first = big.NewInt(setBlockNumber).Cmp(big.NewInt(blockNumber)) == 0

	return validator.SignalToChange(first, receipts, blockNumber, simulatedBackend)
}

func (multi *Multi) FinalizeChange(header *types.Header, state *state.StateDB) error {
	validator, _ := multi.correctSet(header.Number)
	return validator.FinalizeChange(header, state)
}

func (multi *Multi) GetValidatorsByCaller(blockNumber int64) []common.Address {
	validator, setBlockNumber := multi.correctSet(big.NewInt(blockNumber))
	first := big.NewInt(setBlockNumber).Cmp(big.NewInt(blockNumber)) == 0

	if first {
		log.Debug("Getting validator list at transition height", "set", validator, "blockNumber", blockNumber)
		validator, _ := multi.correctSet(big.NewInt(blockNumber - 1))
		return validator.GetValidatorsByCaller(blockNumber)
	}

	return validator.GetValidatorsByCaller(blockNumber)
}

func (multi *Multi) PrepareBackend(blockNumber int64, simulatedBackend bind.ContractBackend) error {
	validator, _ := multi.correctSet(big.NewInt(blockNumber))
	return validator.PrepareBackend(blockNumber, simulatedBackend)
}