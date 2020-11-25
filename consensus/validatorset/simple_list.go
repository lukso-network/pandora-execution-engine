package validatorset

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
)

type SimpleList struct {
	validators []common.Address
}

func NewSimpleList(validators []common.Address) *SimpleList {
	return &SimpleList{
		validators: validators,
	}
}

func (simpleList *SimpleList) SignalToChange(first bool, logs []*types.Log, header *types.Header) bool {
	panic("implement me")
}

func (simpleList *SimpleList) FinalizeChange(header *types.Header, state *state.StateDB) error {
	panic("implement me")
}

func (simpleList *SimpleList) GetValidatorsByCaller(header *types.Header) []common.Address {
	panic("implement me")
}

func (simpleList *SimpleList) CountValidators() {
	panic("implement me")
}


