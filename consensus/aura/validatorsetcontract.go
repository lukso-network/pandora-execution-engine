package aura

//go:generate abigen --sol validatorset/ValidatorSet.sol --pkg validatorset --out validatorset/validatorsetcontract.go

import (
"github.com/ethereum/go-ethereum/accounts/abi/bind"
"github.com/ethereum/go-ethereum/common"
"github.com/ethereum/go-ethereum/consensus/aura/validatorset"
	"math/big"
)

// RelaySet is a Go wrapper around an on-chain validator set contract.
type ValidatorSetContract struct {
	address  		common.Address
	contract 		*validatorset.ValidatorSet
	backend 		bind.ContractBackend // Smart contract backend
	ValidatorList 	[]common.Address
}

// NewValidatorSet binds ValidatorSet contract and returns a registrar instance.
func NewValidatorSet(contractAddr common.Address, backend bind.ContractBackend) (*ValidatorSetContract, error) {
	c, err := validatorset.NewValidatorSet(contractAddr, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetContract{
		address: contractAddr,
		contract: c}, nil
}

// ContractAddr returns the address of contract.
func (validatorSetContract *ValidatorSetContract) ContractAddr() common.Address {
	return validatorSetContract.address
}

// Contract returns the underlying contract instance.
func (validatorSetContract *ValidatorSetContract) Contract() *validatorset.ValidatorSet {
	return validatorSetContract.contract
}

// Contract returns all the validator list.
func (validatorSetContract *ValidatorSetContract) GetValidators(blockNumber *big.Int) []common.Address {
	callOpts := &bind.CallOpts{BlockNumber: blockNumber}
	validatorList, error := validatorSetContract.contract.ValidatorSetCaller.GetValidators(callOpts)
	if error == nil {
		validatorSetContract.ValidatorList = validatorList
		return validatorList
	}
	return validatorSetContract.ValidatorList
}