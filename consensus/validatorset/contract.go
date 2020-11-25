package validatorset

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/params"
)

type Contract struct {
	contractAddress 	common.Address
	validators 			*ValidatorSafeContract
}


func NewContract(contractAddr common.Address, chain *core.BlockChain,
	chainDb ethdb.Database, chainConfig *params.ChainConfig) *Contract {
	return &Contract{
		contractAddress: contractAddr,
		validators: NewValidatorSafeContract(contractAddr, chain, chainDb, chainConfig),
	}
}

func (c *Contract) SignalToChange(first bool, logs []*types.Log, header *types.Header) bool {
	return c.validators.SignalToChange(first, logs, header)
}

func (c *Contract) FinalizeChange(header *types.Header, state *state.StateDB) error {
	return c.validators.FinalizeChange(header, state)
}

func (c *Contract) GetValidatorsByCaller(header *types.Header) []common.Address {
	return c.validators.GetValidatorsByCaller(header)
}

func (c *Contract) CountValidators() {
	panic("implement me")
}
