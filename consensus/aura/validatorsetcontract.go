package aura

//go:generate abigen --sol validatorset/ValidatorSet.sol --pkg validatorset --out validatorset/validatorsetcontract.go

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/aura/validatorset"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
)

var SYSTEM_ADDRESS = common.HexToAddress("0xfffffffffffffffffffffffffffffffffffffffe")

// RelaySet is a Go wrapper around an on-chain validator set contract.
type ValidatorSetContract struct {
	address  				common.Address		// contract address
	contract 				*validatorset.ValidatorSet
	backend 				bind.ContractBackend // Smart contract backend
	eventCh 				chan *validatorset.ValidatorSetInitiateChange
	exitCh          		chan struct{}
	auraEngine				*Aura
	isStarted				bool
}

func NewValidatorSetWithSimBackend(contractAddr common.Address, chain *core.BlockChain, chainDb ethdb.Database, config *params.ChainConfig, aura *Aura) (*ValidatorSetContract, error) {
	log.Info("Signal for switch to contract-based validator set")
	simBackend := backends.NewSimulatedBackendWithChain(chain, chainDb, config)

	c, err := validatorset.NewValidatorSet(contractAddr, simBackend)
	if err != nil {
		log.Debug("Getting error when initialize validator set contract", "error", err)
		return nil, err
	}
	contract := &ValidatorSetContract{
		address: contractAddr,
		contract: c,
		backend: simBackend,
		eventCh: make(chan *validatorset.ValidatorSetInitiateChange),
		exitCh: make(chan struct{}),
		auraEngine: aura,
	}
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
func (v *ValidatorSetContract) GetValidators(blockNumber *big.Int) []common.Address {
	validatorList, err := v.contract.ValidatorSetCaller.GetValidators(nil)
	if err != nil {
		return nil
	}
	log.Debug("Calling validator list.", "validatorList", validatorList)
	return validatorList
}

// System call - finalizeChange function
func (v *ValidatorSetContract) FinalizeChange() (*types.Transaction, error) {
	simBackend, ok := v.backend.(*backends.SimulatedBackend)
	if !ok {
		log.Error("Getteing error in simulated backed")
		return nil, nil
	}

	if v.auraEngine.curHeader == nil || v.auraEngine.curStateDB == nil {
		log.Error("Current header and state is nil")
		return nil, nil
	}
	log.Debug("in finalizeChange method", "curHeader", v.auraEngine.curHeader.Number)
	simBackend.PrepareCurrentState(v.auraEngine.curHeader, v.auraEngine.curStateDB)

	opts := &bind.TransactOpts{
		From: SYSTEM_ADDRESS,
	}
	tx, err := v.contract.FinalizeChange(opts)
	if err != nil {
		log.Error("Getting error from method calling", "err", err)
		return nil, err
	}
	return tx, err
}

// Starting validator set contract
func (v *ValidatorSetContract) StartWatchingEvent() {
	// Activated validator set contract
	v.FinalizeChange()
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
			_, err := v.FinalizeChange()
			if err == nil {
				v.auraEngine.validatorList = event.NewSet
				log.Debug("Getting validator list from finalizeMethod", "validatorList", v.auraEngine.validatorList)
			}

		case <-v.exitCh:
			log.Debug("Going to shutdown watching")
			return
		}
	}
	initChangeSub.Unsubscribe()
}