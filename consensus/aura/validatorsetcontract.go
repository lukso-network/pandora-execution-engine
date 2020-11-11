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
	"reflect"
)

var SYSTEM_ADDRESS = common.HexToAddress("0xfffffffffffffffffffffffffffffffffffffffe")

// RelaySet is a Go wrapper around an on-chain validator set contract.
type ValidatorSetContract struct {
	address  				common.Address		// contract address
	contract 				*validatorset.ValidatorSet
	backend 				bind.ContractBackend // Smart contract backend
	eventCh 				chan *validatorset.ValidatorSetInitiateChange
	exitCh          		chan struct{}
	isStarted				bool
	pendingValidatorList	[]common.Address
	currentValidatorList	[]common.Address
	successCall				bool
}

func NewValidatorSetWithSimBackend(contractAddr common.Address, chain *core.BlockChain,
	chainDb ethdb.Database, config *params.ChainConfig) (*ValidatorSetContract, error) {

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
	return validatorList
}

// CheckAndFinalizeChange method will check and call FinalizeChange method so that pending list gets finalized
// in validator set contract
func (v *ValidatorSetContract) CheckAndFinalizeChange(header *types.Header, stateDB *state.StateDB) ([]common.Address, error) {
	if len(v.pendingValidatorList) == 0 { return nil, nil }
	if reflect.DeepEqual(v.pendingValidatorList, v.currentValidatorList) { return nil, nil }
	_, err := v.FinalizeChange(header, stateDB)
	if err != nil {
		log.Error("Getting error from method calling", "err", err)
		return nil, err
	}

	log.Debug("Now updated validator list", "pending", v.pendingValidatorList)
	return v.pendingValidatorList, nil
}

// System call - finalizeChange function
func (v *ValidatorSetContract) FinalizeChange(header *types.Header, stateDB *state.StateDB) (*types.Transaction, error) {
	simBackend, ok := v.backend.(*backends.SimulatedBackend)
	if !ok {
		log.Error("invalid simulated backed.", "invalidSimBackend", v.backend)
		return nil, nil
	}
	simBackend.PrepareCurrentState(header, stateDB)
	opts := &bind.TransactOpts{
		From: SYSTEM_ADDRESS,
	}
	tx, err := v.contract.FinalizeChange(opts)
	return tx, err
}

// Starting validator set contract
func (v *ValidatorSetContract) StartWatchingEvent() {
	// Activated validator set contract
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
		case <-v.exitCh:
			log.Debug("Going to shutdown event watching")
			return
		}
	}
	initChangeSub.Unsubscribe()
}