package validatorset

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	validatorset "github.com/ethereum/go-ethereum/consensus/validatorset/res"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	lru "github.com/hashicorp/golang-lru"
	"sort"
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

func (vsc *ValidatorSafeContract) parseInitiateChangeEvent(receipts types.Receipts) ([]common.Address, bool) {
	for _, receipt := range receipts {
		for _, txlog := range receipt.Logs {
			// checks the transaction's from address
			if txlog.Address == vsc.contractAddress {
				log.Debug("detected epoch change event bloom")

				var logValue types.Log
				logValue = *txlog
				event, err := vsc.contract.ParseInitiateChange(logValue)
				if err != nil { return nil, false }

				log.Info("Signal for transition within contract", "New List", event.NewSet)
				return event.NewSet, true
			}
		}
	}

	return nil, false
}

func (vsc *ValidatorSafeContract) SignalToChange(first bool, receipts types.Receipts, blockNumber int64, simulatedBackend bind.ContractBackend) ([]common.Address, bool, bool) {
	if first {
		log.Debug("signalling transition to fresh contract.")
		if err := vsc.PrepareBackend(blockNumber, simulatedBackend); err != nil {
			log.Error("error when preparing backend for contract", "error", err)
			return nil, false, true
		}
		validators := vsc.GetValidatorsByCaller(blockNumber)
		// sort when get the validator list so that aura engine can select validator list deterministically
		sort.Slice(validators, func(i, j int) bool {
			return validators[i].Hex() < validators[j].Hex()
		})

		log.Info("Signal for switch to contract-based validator set.")
		log.Debug("Initial contract validators", "validatorSet", validators)

		return validators, true, true
	}

	validators, hasSignal := vsc.parseInitiateChangeEvent(receipts)
	// sort when get the validator list so that aura engine can select validator list deterministically
	sort.Slice(validators, func(i, j int) bool {
		return validators[i].Hex() < validators[j].Hex()
	})
	return validators, hasSignal, false
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

func (vsc *ValidatorSafeContract) GetValidatorsByCaller(blockNumber int64) []common.Address {
	//if validators, ok := vsc.validators.Get(blockNumber); ok {
	//	log.Debug("Set of validators obtained from lru cache", "validators", vsc.validators)
	//	return validators.([]common.Address)
	//}
	validators, err := vsc.contract.GetValidators(nil)
	if err == nil {
		log.Debug("Set of validators obtained from contract", "validators", validators)
		// sort when get the validator list so that aura engine can select validator list deterministically
		sort.Slice(validators, func(i, j int) bool {
			return validators[i].Hex() < validators[j].Hex()
		})
		vsc.validators.Add(blockNumber, validators)
		return validators
	}
	return nil
}

func (vsc *ValidatorSafeContract) CountValidators() int {
	panic("implement me")
}

func (vsc *ValidatorSafeContract) PrepareBackend(blockNumber int64, simulatedBackend bind.ContractBackend) error {
	if vsc.backend == nil {
		log.Debug("Preparing simulated backend for contract", "blockNumber", blockNumber)
		contract, err := validatorset.NewValidatorSet(vsc.contractAddress, simulatedBackend)

		if err != nil {
			return err
		}

		vsc.contract = contract
		vsc.backend = simulatedBackend
	}
	return nil
}
