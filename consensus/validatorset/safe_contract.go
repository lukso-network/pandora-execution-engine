package validatorset

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	validatorset "github.com/ethereum/go-ethereum/consensus/validatorset/res"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	lru "github.com/hashicorp/golang-lru"
	"math/big"
)

//go:generate abigen --sol res/ValidatorSet.sol --pkg validatorset --out res/validator_contract.go

/// A validator contract with reporting.
type ValidatorSafeContract struct {
	contractAddress 	common.Address
	validators 	    	*lru.ARCCache
	backend 			bind.ContractBackend
	contract 			*validatorset.ValidatorSet
}

func NewValidatorSafeContract(contractAddr common.Address) *ValidatorSafeContract {
	validators, _ := lru.NewARC(1024)
	return &ValidatorSafeContract{
		contractAddress: contractAddr,
		validators: 	 validators,
		backend: 		 nil,
		contract: 		 nil,
	}
}

func (vsc *ValidatorSafeContract) parseInitiateChangeEvent(logs []*types.Log) ([]common.Address, bool) {
	for _, txlog := range logs {
		// checks the transaction's from address
		if txlog.Address == vsc.contractAddress {
			log.Trace("detected epoch change event bloom")

			var logValue types.Log
			logValue = *txlog
			event, err := vsc.contract.ParseInitiateChange(logValue)
			if err != nil { return nil, false }

			log.Info("Signal for transition within contract", "New List", event.NewSet)
			return event.NewSet, true
		}
	}
	return nil, false
}

func (vsc *ValidatorSafeContract) SignalToChange(first bool, logs []*types.Log, header *types.Header) ([]common.Address, bool) {
	if first {
		log.Info("signalling transition to fresh contract.")
		validators := vsc.GetValidatorsByCaller(header.Number)

		log.Info("Signal for switch to contract-based validator set.")
		log.Trace("Initial contract validators", "validatorSet", validators)

		return validators, true
	}
	validators, hasSignal := vsc.parseInitiateChangeEvent(logs)
	return validators, hasSignal
}

func (vsc *ValidatorSafeContract) FinalizeChange(header *types.Header, state *state.StateDB) error {
	// prepare current stateDB for changing state trie
	vsc.backend.PrepareCurrentState(header, state)
	opts := &bind.TransactOpts{
		From: SYSTEM_ADDRESS,
	}
	_, err := vsc.contract.FinalizeChange(opts)
	if err != nil {
		return err
	}
	return nil
}

func (vsc *ValidatorSafeContract) GetValidatorsByCaller(blockNumber *big.Int) []common.Address {
	if validators, ok := vsc.validators.Get(blockNumber); ok {
		log.Trace("Set of validators obtained", "validators", validators)
		return validators.([]common.Address)
	}
	if validators, err := vsc.contract.GetValidators(nil); err != nil {
		log.Trace("Set of validators obtained", "validators", validators)
		vsc.validators.Add(blockNumber, validators)
		return validators
	}
	return nil
}

func (vsc *ValidatorSafeContract) CountValidators() int {
	panic("implement me")
}

func (vsc *ValidatorSafeContract) PrepareBackend(header *types.Header, chain *core.BlockChain, chainDb ethdb.Database) error {
	if vsc.backend == nil {
		log.Trace("Preparing simulated backend for contract", "blockNumber", header.Number)

		simulatedBackend := backends.NewSimulatedBackendWithChain(chain, chainDb, chain.Config())
		contract, err := validatorset.NewValidatorSet(vsc.contractAddress, simulatedBackend)

		if err != nil {
			return err
		}

		vsc.contract = contract
		vsc.backend = simulatedBackend
	}
	return nil
}
