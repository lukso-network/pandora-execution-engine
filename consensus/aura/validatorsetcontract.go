package aura

//go:generate abigen --sol validatorset/ValidatorSet.sol --pkg validatorset --out validatorset/validatorsetcontract.go

import (
"github.com/ethereum/go-ethereum/accounts/abi/bind"
"github.com/ethereum/go-ethereum/common"
"github.com/ethereum/go-ethereum/consensus/aura/validatorset"
)

// RelaySet is a Go wrapper around an on-chain validator set contract.
type ValidatorSetContract struct {
	address  common.Address
	contract *validatorset.ValidatorSet
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