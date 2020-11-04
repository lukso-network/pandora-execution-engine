package aura

//go:generate abigen --sol validatorset/RelaySet.sol --pkg validatorset --out validatorset/relaySet.go

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/aura/validatorset"
)

// RelaySet is a Go wrapper around an on-chain validator set contract.
type RelaySet struct {
	address  common.Address
	contract *validatorset.BaseOwnedSet
}

// NewRelaySet binds relaySet contract and returns a registrar instance.
func NewRelaySet(contractAddr common.Address, backend bind.ContractBackend) (*RelaySet, error) {
	c, err := validatorset.NewBaseOwnedSet(contractAddr, backend)
	if err != nil {
		return nil, err
	}
	return &RelaySet{
		address: contractAddr,
		contract: c}, nil
}

// ContractAddr returns the address of contract.
func (relaySet *RelaySet) ContractAddr() common.Address {
	return relaySet.address
}

// Contract returns the underlying contract instance.
func (relaySet *RelaySet) Contract() *validatorset.BaseOwnedSet {
	return relaySet.contract
}

// Lookup searches ChangeFinalizedEvents event
//func (relaySet *RelaySet) LookupBaseOwnedSetChangeFinalizedEvents(blockLogs [][]*types.Log, section uint64, hash common.Hash) {
//	var changeFinlizeEvents []*validatorset.BaseOwnedSetChangeFinalized
//
//	for _, logs := range blockLogs {
//		for _, log := range logs {
//			event, err := relaySet.contract.ParseChangeFinalized(*log)
//			if err != nil {
//				continue
//			}
//			if event.CurrentSet == section && common.Hash(event.CurrentSet) == hash {
//				changeFinlizeEvents = append(changeFinlizeEvents, event)
//			}
//		}
//	}
//	return votes
//}
