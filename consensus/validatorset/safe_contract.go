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

// A validator contract with reporting.
type ValidatorSafeContract struct {
	contractAddress 	common.Address			// safe contract address
	validators 	    	*lru.ARCCache			// validators cache reduces db call to get validator set from contract
	backend 			bind.ContractBackend	// backend interacts with validator set contract
	contract 			*validatorset.ValidatorSet	// validator set contract instance
}

// Creates ValidatorSafeContract instance
func NewValidatorSafeContract(contractAddr common.Address) *ValidatorSafeContract {
	validators, _ := lru.NewARC(1024)
	return &ValidatorSafeContract{
		contractAddress: contractAddr,
		validators: 	 validators,
		backend: 		 nil,
		contract: 		 nil,
	}
}

// Whether the contract address matches with tx log address.
//
// The expected log should have 3 topics:
//   1. ETHABI-encoded log name.
//   2. the block's parent hash.
//   3. the "nonce": n for the nth transition in history.
//
// We can only search for the first 2, since we don't have the third
// just yet.
//
// The log data is an array of all new validator addresses.
// check receipts for log event. bloom should be `expected_bloom` for the
// header the receipts correspond to.
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
			// this will not happen unless simulatedBackend is wrong
			panic(err)
		}
		// getting initial validators from contract
		validators, _ := vsc.contract.GetValidators(nil)

		// sort when get the validator list so that aura engine can select validator list deterministically
		sort.Slice(validators, func(i, j int) bool {
			return validators[i].Hex() < validators[j].Hex()
		})

		log.Info("Signal for switch to contract-based validator set.")
		log.Debug("Initial contract validators", "validatorSet", validators)

		return validators, first, true
	}

	log.Debug("Expected topics. topic=[0x55252fa6eee4741b4e24a74a70e9c11fd2c2281df8d6ea13126ff845f7825c89]", "number=", blockNumber)
	// check receipts for log event. bloom should be `expected_bloom` for the
	// header the receipts correspond to.
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
		From: SystemAddress,
	}
	_, err := vsc.contract.FinalizeChange(opts)
	if err != nil {
		return err
	}
	return nil
}

func (vsc *ValidatorSafeContract) GetValidatorsByCaller(blockNumber int64) []common.Address {
	// first check whether validator set is already in cache or not.
	// If exists in cache, then return directly from cache
	if validators, ok := vsc.validators.Get(blockNumber); ok {
		log.Debug("Set of validators obtained from lru cache", "validators", vsc.validators)
		return validators.([]common.Address)
	}

	// getting validator set from smart contract.
	validators, err := vsc.contract.GetValidators(nil)

	// error comes when validator set contract address is wrong or simulated
	// backend is wrong. Generally this will not happen
	if err != nil {
		panic(err)
	}

	log.Debug("Set of validators obtained from contract", "validators", validators)
	// sort when get the validator list so that aura engine can select validator list deterministically
	sort.Slice(validators, func(i, j int) bool {
		return validators[i].Hex() < validators[j].Hex()
	})
	// added to cache for reducing db call to get validator set
	vsc.validators.Add(blockNumber, validators)

	return validators
}


func (vsc *ValidatorSafeContract) PrepareBackend(blockNumber int64, simulatedBackend bind.ContractBackend) error {
	// if contract backend is not configured then set in safe contract
	if vsc.backend == nil {
		log.Debug("Setting up contract caller. at :", "blockNumber", blockNumber)
		contract, err := validatorset.NewValidatorSet(vsc.contractAddress, simulatedBackend)

		if err != nil {
			return err
		}

		vsc.contract = contract
		vsc.backend = simulatedBackend
	}
	return nil
}
