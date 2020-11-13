package aura

//go:generate abigen --sol validatorset/ValidatorSet.sol --pkg validatorset --out validatorset/validator_contract.go

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/aura/validatorset"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

var SYSTEM_ADDRESS = common.HexToAddress("0xfffffffffffffffffffffffffffffffffffffffe")

// RelaySet is a Go wrapper around an on-chain validator set contract.
type ValidatorSetContract struct {
	address  				common.Address		// contract address
	contract 				*validatorset.ValidatorSet
	backend 				bind.ContractBackend // Smart contract backend
	eventCh 				chan *validatorset.ValidatorSetInitiateChange
	exitCh          		chan struct{}
	pendingValidatorList	[]common.Address
	hasChange				bool
}

func NewValidatorSetWithSimBackend(contractAddr common.Address, backend bind.ContractBackend) (*ValidatorSetContract, error) {
	log.Info("Signal for switch to contract-based validator set")
	c, err := validatorset.NewValidatorSet(contractAddr, backend)

	if err != nil {
		log.Debug("Getting error when initialize validator set contract", "error", err)
		return nil, err
	}

	contract := &ValidatorSetContract{
		address: contractAddr,
		contract: c,
		backend: backend,
		eventCh: make(chan *validatorset.ValidatorSetInitiateChange),
		exitCh: make(chan struct{}),
		pendingValidatorList: nil,
		hasChange: true,
	}
	go contract.watchingInitiateChangeLoop()
	return contract, nil
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
func (v *ValidatorSetContract) GetValidators() []common.Address {
	validatorList, err := v.contract.ValidatorSetCaller.GetValidators(nil)
	if err != nil {
		return nil
	}
	log.Debug("Calling validator list.", "validatorList", validatorList)

	//todo - need to sort validatorList
	return validatorList
}

// System call - finalizeChange function
func (v *ValidatorSetContract) FinalizeChange(header *types.Header, state *state.StateDB) (*types.Transaction, error) {
	if v.hasChange == false {
		log.Debug("no validator list changes in this epoch!")
		return nil, nil
	}
	v.backend.PrepareCurrentState(header, state)
	opts := &bind.TransactOpts{
		From: SYSTEM_ADDRESS,
	}
	tx, err := v.contract.FinalizeChange(opts)
	if err != nil {
		log.Error("getting error from method calling", "err", err)
		return nil, err
	}
	v.hasChange = false
	return tx, nil
}

// Starting validator set contract
func (v *ValidatorSetContract) StartWatchingEvent() {
	go v.watchingInitiateChangeLoop()
}

func (v *ValidatorSetContract) StopWatchingEvent() {
	close(v.exitCh)
}

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
		log.Debug("Watching on initiateChange event : 0x55252fa6eee4741b4e24a74a70e9c11fd2c2281df8d6ea13126ff845f7825c89" )
		select {
		case event := <-v.eventCh:
			log.Debug("getting a new event", "event", event)
			// Short circuit when receiving empty result.
			if event == nil {
				continue
			}
			log.Debug("Getting pending validator list", "newValidatorSet", event.NewSet)
			v.pendingValidatorList = event.NewSet
			v.hasChange = true
		case <-v.exitCh:
			log.Debug("Going to shutdown event watching")
			return
		}
	}
	initChangeSub.Unsubscribe()
}