package validatorset

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

var (
	EmptyAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")
	SYSTEM_ADDRESS = common.HexToAddress("0xfffffffffffffffffffffffffffffffffffffffe")
)

func NewValidatorSet(multiMap map[int]ValidatorSet, validatorSpec *params.ValidatorSet) ValidatorSet {

	if validatorSpec.List != nil { return NewSimpleList(validatorSpec.List)
	} else if validatorSpec.SafeContract != EmptyAddress {
		return NewValidatorSafeContract(validatorSpec.SafeContract)
	} else if validatorSpec.Contract != EmptyAddress {
		return NewContract(validatorSpec.Contract)
	} else {
		for key, value := range validatorSpec.Multi {
			multiMap[int(key)] = NewValidatorSet(multiMap, &value)
		}
		return NewMulti(multiMap)
	}
}


type ValidatorSet interface {
	SignalToChange(first bool, receipts types.Receipts, blockNumber int64, simulatedBackend bind.ContractBackend) ([]common.Address, bool, bool)

	FinalizeChange(header *types.Header, state *state.StateDB) error

	GetValidatorsByCaller(blockNumber int64) []common.Address

	CountValidators() int

	PrepareBackend(blockNumber int64, simulatedBackend bind.ContractBackend) error
}