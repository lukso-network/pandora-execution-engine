package validatorset

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
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

func (c *Contract) SignalToChange(first bool, receipts types.Receipts, blockNumber int64, simulatedBackend bind.ContractBackend) ([]common.Address, bool, bool) {
	return c.validators.SignalToChange(first, receipts, blockNumber, simulatedBackend)
}

func (c *Contract) FinalizeChange(header *types.Header, state *state.StateDB) error {
	return c.validators.FinalizeChange(header, state)
}

func (c *Contract) GetValidatorsByCaller(chainHeader int64) []common.Address {
	return c.validators.GetValidatorsByCaller(chainHeader)
}

func (c *Contract) CountValidators() int {
	panic("implement me")
}

func (c *Contract) PrepareBackend(blockNumber int64, simulatedBackend bind.ContractBackend) error {
	return c.validators.PrepareBackend(blockNumber, simulatedBackend)
}
