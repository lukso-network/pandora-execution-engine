package aura

//go:generate abigen --sol validatorset/ValidatorSet.sol --pkg validatorset --out validatorset/validator_contract.go

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/aura/validatorset"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"sort"
)

// Use for calling validator set contract method
var SYSTEM_ADDRESS = common.HexToAddress("0xfffffffffffffffffffffffffffffffffffffffe")

// ValidatorSetContract implements to interact with validator set contract
type ValidatorSetContract struct {
	address  				common.Address		// contract address
	contract 				*validatorset.ValidatorSet
	backend 				bind.ContractBackend // Smart contract backend
	eventCh 				chan *validatorset.ValidatorSetInitiateChange
	exitCh          		chan struct{}
}

// NewValidatorSetWithSimBackend implements simulated backend to initiate validator set contract
func NewValidatorSetWithSimBackend(contractAddr common.Address, backend bind.ContractBackend) (*ValidatorSetContract, error) {
	c, err := validatorset.NewValidatorSet(contractAddr, backend)
	if err != nil {
		return nil, err
	}

	return &ValidatorSetContract{
		address: contractAddr,
		contract: c,
		backend: backend,
		eventCh: make(chan *validatorset.ValidatorSetInitiateChange),
		exitCh: make(chan struct{}),
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
func (v *ValidatorSetContract) GetValidators(blockNumber *big.Int ) []common.Address {
	var opts *bind.CallOpts
	if blockNumber != nil {
		opts = &bind.CallOpts{
			BlockNumber: blockNumber,
		}
	}
	// get validator list from smart contract
	validatorList, err := v.contract.ValidatorSetCaller.GetValidators(opts)
	if err != nil {
		return nil
	}

	// sort when get the validator list so that aura engine can select validator list deterministically
	sort.Slice(validatorList, func(i, j int) bool {
		return validatorList[i].Hex() < validatorList[j].Hex()
	})
	return validatorList
}

// FinalizeChange changes state of validator set contract. this method implements
// to call finalizeChange method of smart contract
func (v *ValidatorSetContract) FinalizeChange(header *types.Header, state *state.StateDB) (*types.Transaction, error) {
	// prepare current stateDB for changing state trie
	v.backend.PrepareCurrentState(header, state)

	opts := &bind.TransactOpts{
		From: SYSTEM_ADDRESS,
	}
	tx, err := v.contract.FinalizeChange(opts)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Starting event watching
func (v *ValidatorSetContract) StartWatchingEvent() {
	go v.watchingInitiateChangeLoop()
}

// StopWatchingEvent stop event watching
func (v *ValidatorSetContract) StopWatchingEvent() {
	close(v.exitCh)
}

// watchingInitiateChangeLoop implements to watch on event of validator set contract
func (v *ValidatorSetContract) watchingInitiateChangeLoop() {
	opts := &bind.WatchOpts{
		Start: nil,
		Context: nil,
	}
	initChangeSub, err := v.contract.WatchInitiateChange(opts, v.eventCh, nil)
	if err != nil {
		log.Error("Getting error when start watching on event", "err", err)
		return
	}
	for {
		select {
		case event := <-v.eventCh:
			// Short circuit when receiving empty result.
			if event == nil {
				continue
			}
		case <-v.exitCh:
			return
		}
	}
	initChangeSub.Unsubscribe()
}