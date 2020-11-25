package validatorset

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"math/big"
)

type Contract struct {
	contractAddress 	common.Address
	validators 			*ValidatorSafeContract
}


func NewContract(contractAddr common.Address) *Contract {
	return &Contract{
		contractAddress: contractAddr,
		validators: NewValidatorSafeContract(contractAddr),
	}
}

func (c *Contract) SignalToChange(first bool, logs []*types.Log, header *types.Header) ([]common.Address, bool) {
	return c.validators.SignalToChange(first, logs, header)
}

func (c *Contract) FinalizeChange(header *types.Header, state *state.StateDB) error {
	return c.validators.FinalizeChange(header, state)
}

func (c *Contract) GetValidatorsByCaller(blockNumber *big.Int) []common.Address {
	return c.validators.GetValidatorsByCaller(blockNumber)
}

func (c *Contract) CountValidators() int {
	panic("implement me")
}

func (c *Contract) PrepareBackend(header *types.Header, chain *core.BlockChain, chainDb ethdb.Database) error {
	return c.validators.PrepareBackend(header, chain, chainDb)
}
