package validatorset

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
)

type SimpleList struct {
	validators []common.Address
}

func NewSimpleList(validators []common.Address) *SimpleList {
	return &SimpleList{
		validators: validators,
	}
}

func (simpleList *SimpleList) SignalToChange(first bool, receipts types.Receipts, header *types.Header, simulatedBackend bind.ContractBackend) ([]common.Address, bool, bool) {
	return simpleList.validators, false, false
}

func (simpleList *SimpleList) FinalizeChange(header *types.Header, state *state.StateDB) error {
	panic("implement me")
}

func (simpleList *SimpleList) GetValidatorsByCaller(blockNumber *big.Int) []common.Address {
	log.Trace("Set of validators obtained from simpleList", "validators", simpleList.validators)
	return simpleList.validators
}

func (simpleList *SimpleList) CountValidators() int {
	return len(simpleList.validators)
}

func (simpleList *SimpleList) PrepareBackend(header *types.Header, simulatedBackend bind.ContractBackend) error {
	return nil
}


