package validatorset

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

var (
	EmptyAddress  = common.HexToAddress("0x0000000000000000000000000000000000000000")
	SystemAddress = common.HexToAddress("0xfffffffffffffffffffffffffffffffffffffffe")
)

// Creates a validator set from spec. Parse validator set spec and creates validatorSet.
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

// A validator set.
type ValidatorSet interface {
	// Whether the given block signals the end of an epoch, but change won't take effect
	// until finality.
	//
	// Engine should set `first` only if the header is genesis. Multiplexing validator
	// sets can set `first` to internal changes.
	SignalToChange(first bool, receipts types.Receipts, blockNumber int64, simulatedBackend bind.ContractBackend) ([]common.Address, bool, bool)

	// Signalling that a new epoch has begun.
	//
	// All calls here will be from the `SYSTEM_ADDRESS`: 2^160 - 2
	// and will have an effect on the block's state.
	FinalizeChange(header *types.Header, state *state.StateDB) error

	// Draws validator list from static list or from validator set contract.
	GetValidatorsByCaller(blockNumber int64) []common.Address

	// Prepare simulated backend for interacting with smart contract
	PrepareBackend(blockNumber int64, simulatedBackend bind.ContractBackend) error
}