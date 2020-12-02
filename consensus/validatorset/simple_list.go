package validatorset

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

// SimpleList represents the static validator set
type SimpleList struct {
	validators []common.Address
}

// Creates instance of SimpleList
func NewSimpleList(validators []common.Address) *SimpleList {
	return &SimpleList{
		validators: validators,
	}
}

func (simpleList *SimpleList) SignalToChange(first bool, receipts types.Receipts, blockNumber int64, simulatedBackend bind.ContractBackend) ([]common.Address, bool, bool) {
	return simpleList.validators, false, false
}

func (simpleList *SimpleList) FinalizeChange(header *types.Header, state *state.StateDB) error {
	return nil
}

func (simpleList *SimpleList) GetValidatorsByCaller(blockNumber int64) []common.Address {
	log.Debug("Set of validators obtained from simpleList", "validators", simpleList.validators)
	return simpleList.validators
}

func (simpleList *SimpleList) PrepareBackend(blockNumber int64, simulatedBackend bind.ContractBackend) error {
	return nil
}


