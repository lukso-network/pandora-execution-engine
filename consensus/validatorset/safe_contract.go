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
	"github.com/ethereum/go-ethereum/params"
	lru "github.com/hashicorp/golang-lru"
)

//go:generate abigen --sol res/ValidatorSet.sol --pkg validatorset --out res/validator_contract.go

/// A validator contract with reporting.
type ValidatorSafeContract struct {
	contractAddress 	common.Address
	validators 	    	*lru.ARCCache
	backend 			bind.ContractBackend
	contract 			*validatorset.ValidatorSet
}

func NewValidatorSafeContract(contractAddr common.Address, chain *core.BlockChain,
	chainDb ethdb.Database, chainConfig *params.ChainConfig) *ValidatorSafeContract {

	simulatedBackend := backends.NewSimulatedBackendWithChain(chain, chainDb, chainConfig)
	c, err := validatorset.NewValidatorSet(contractAddr, simulatedBackend)
	if err != nil {
		return nil
	}
	validators, _ := lru.NewARC(1024)

	return &ValidatorSafeContract{
		contractAddress: contractAddr,
		validators: 	 validators,
		backend: 		 simulatedBackend,
		contract: 		 c,
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

func (vsc *ValidatorSafeContract) SignalToChange(first bool, logs []*types.Log, header *types.Header) bool {
	if first {
		log.Info("signalling transition to fresh contract.")
		return true
	}
	_, hasSignal := vsc.parseInitiateChangeEvent(logs)
	return hasSignal
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

func (vsc *ValidatorSafeContract) GetValidatorsByCaller(header *types.Header) []common.Address {
	if validators, ok := vsc.validators.Get(header.ParentHash); ok {
		log.Trace("Set of validators obtained", "validators", validators)
		return validators.([]common.Address)
	}
	if validators, err := vsc.contract.GetValidators(nil); err != nil {
		log.Trace("Set of validators obtained", "validators", validators)
		vsc.validators.Add(header.ParentHash, validators)
		return validators
	}
	return nil
}

func (vsc *ValidatorSafeContract) CountValidators() {
	panic("implement me")
}
