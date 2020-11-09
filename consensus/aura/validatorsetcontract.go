package aura

//go:generate abigen --sol validatorset/ValidatorSet.sol --pkg validatorset --out validatorset/validatorsetcontract.go

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/aura/validatorset"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
)

var SYSTEM_ADDRESS = common.HexToAddress("0xfffffffffffffffffffffffffffffffffffffffe")

// RelaySet is a Go wrapper around an on-chain validator set contract.
type ValidatorSetContract struct {
	address  		common.Address		// contract address
	contract 		*validatorset.ValidatorSet
	backend 		bind.ContractBackend // Smart contract backend
	ValidatorList 	[]common.Address
}

func NewValidatorSetWithSimBackend(contractAddr common.Address, chain *core.BlockChain, chainDb ethdb.Database, config *params.ChainConfig) (*ValidatorSetContract, error) {
	simBackend := backends.NewSimulatedBackendWithChain(chain, chainDb, config)
	log.Debug("Getting simulated backed", "backend", simBackend)

	c, err := validatorset.NewValidatorSet(contractAddr, simBackend)
	if err != nil {
		log.Debug("Getting error when initialize validator set contract", "error", err)
		return nil, err
	}
	return &ValidatorSetContract{
		address: contractAddr,
		contract: c,
		backend: simBackend,
	}, nil
}

// ContractAddr returns the address of contract.
func (v *ValidatorSetContract) ContractAddr() common.Address {
	return v.address
}

// Contract returns the underlying contract instance.
func (v *ValidatorSetContract) Contract() *validatorset.ValidatorSet {
	return v.contract
}

// Contract returns all the validator list.
func (v *ValidatorSetContract) GetValidators(blockNumber *big.Int) []common.Address {
	validatorList, err := v.contract.ValidatorSetCaller.GetValidators(nil)
	if err != nil {
		return nil
	}
	log.Debug("Calling validator list.", "validatorList", validatorList)
	v.ValidatorList = validatorList
	return v.ValidatorList
}

// System call - finalizeChange function
func (v *ValidatorSetContract) FinalizeChange(header *types.Header, stateDB *state.StateDB) (*types.Transaction, error) {
	simBackend, ok := v.backend.(*backends.SimulatedBackend)
	if !ok {
		log.Error("Getteing error in simulated backed")
		return nil, nil
	}
	simBackend.PrepareCurrentState(header, stateDB)

	opts := &bind.TransactOpts{
		From: SYSTEM_ADDRESS,
	}
	tx, err := v.contract.FinalizeChange(opts)
	if err != nil {
		return nil, err
	}
	return tx, err
}