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

func (c *Contract) SignalToChange(first bool, receipts types.Receipts, header *types.Header, chain *core.BlockChain, chainDb ethdb.Database) ([]common.Address, bool, bool) {
	return c.validators.SignalToChange(first, receipts, header, chain, chainDb)
}

func (c *Contract) FinalizeChange(header *types.Header, state *state.StateDB) error {
	return c.validators.FinalizeChange(header, state)
}

func (c *Contract) GetValidatorsByCaller(chainHeader *big.Int) []common.Address {
	return c.validators.GetValidatorsByCaller(chainHeader)
}

func (c *Contract) CountValidators() int {
	panic("implement me")
}

func (c *Contract) PrepareBackend(header *types.Header, chain *core.BlockChain, chainDb ethdb.Database) error {
	return c.validators.PrepareBackend(header, chain, chainDb)
}
