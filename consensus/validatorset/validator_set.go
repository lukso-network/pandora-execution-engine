package validatorset

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
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
	SignalToChange(first bool, receipts types.Receipts, header *types.Header, chain *core.BlockChain, chainDb ethdb.Database) ([]common.Address, bool, bool)

	FinalizeChange(header *types.Header, state *state.StateDB) error

	GetValidatorsByCaller(blockNumber *big.Int) []common.Address

	CountValidators() int

	PrepareBackend(header *types.Header, chain *core.BlockChain, chainDb ethdb.Database) error
}