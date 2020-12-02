package validatorset

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
)

// Validator set maintained in a contract, updated using `getValidators` method.
// It can also report validators for misbehaviour with two levels: `reportMalicious` and `reportBenign`.
// A validator contract with reporting.
type Contract struct {
	contractAddress 	common.Address
	validators 			*ValidatorSafeContract
}

// TODO- Need to implement report_malicious and report_benign method
// Creates instance of Contract
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

func (c *Contract) PrepareBackend(blockNumber int64, simulatedBackend bind.ContractBackend) error {
	return c.validators.PrepareBackend(blockNumber, simulatedBackend)
}
