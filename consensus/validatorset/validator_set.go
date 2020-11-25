package validatorset

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
)

var (
	EmptyAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")
	SYSTEM_ADDRESS = common.HexToAddress("0xfffffffffffffffffffffffffffffffffffffffe")
)

type ValidatorSet interface {
	SignalToChange(first bool, logs []*types.Log, header *types.Header) bool
	FinalizeChange(header *types.Header, state *state.StateDB) error
	GetValidatorsByCaller(header *types.Header) []common.Address
	CountValidators()
}

func NewValidatorSet(multiMap map[uint64]ValidatorSet, validatorSpec *params.ValidatorSet,
	chain *core.BlockChain, chainDb ethdb.Database, chainConfig *params.ChainConfig) ValidatorSet {

	if validatorSpec.List != nil {
		return NewSimpleList(validatorSpec.List)
	} else if validatorSpec.SafeContract != EmptyAddress {
		return NewValidatorSafeContract(validatorSpec.SafeContract, chain, chainDb, chainConfig)
	} else if validatorSpec.Contract != EmptyAddress {
		return NewContract(validatorSpec.Contract, chain, chainDb, chainConfig)
	} else {
		log.Debug("parsing multi validator")
		for key, value := range validatorSpec.Multi {
			multiMap[key] = NewValidatorSet(multiMap, &value, chain, chainDb, chainConfig)
		}
		return NewMulti(multiMap)
	}
}