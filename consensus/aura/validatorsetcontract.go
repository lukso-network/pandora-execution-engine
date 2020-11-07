package aura

//go:generate abigen --sol validatorset/ValidatorSet.sol --pkg validatorset --out validatorset/validatorsetcontract.go

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/aura/validatorset"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
)

var SYSTEM_ADDRESS = common.HexToAddress("0xfffffffffffffffffffffffffffffffffffffffe")

// RelaySet is a Go wrapper around an on-chain validator set contract.
type ValidatorSetContract struct {
	address  		common.Address		// contract address
	contract 		*validatorset.ValidatorSet
	backend 		bind.ContractBackend // Smart contract backend
	ValidatorList 	[]common.Address
	signTxFn 		SignTxFn       // Signer function to authorize hashes with
	signer 			common.Address // Ethereum address of the signing key
	chainID 		*big.Int
}

// newRPCClient creates a rpc client with specified node URL.
func newRPCClient(url string) *rpc.Client {
	client, err := rpc.Dial(url)
	if err != nil {
		log.Debug("Failed to connect to Ethereum node: %v", err)
	}
	return client
}

// NewValidatorSet binds ValidatorSet contract and returns a registrar instance.
func NewValidatorSet(contractAddr common.Address) (*ValidatorSetContract, error) {
	backend := ethclient.NewClient(newRPCClient("http://localhost:8545"))
	c, err := validatorset.NewValidatorSet(contractAddr, backend)
	if err != nil {
		return nil, err
	}
	log.Debug("Validator contract is getting initiated ", "contract", c)
	return &ValidatorSetContract{
		address: contractAddr,
		contract: c}, nil
}

func NewValidatorSetWithSimBackend(contractAddr common.Address, backend bind.ContractBackend) (*ValidatorSetContract, error) {
	c, err := validatorset.NewValidatorSet(contractAddr, backend)
	if err != nil {
		log.Debug("Getting error when initialize validator set contract", "error", err)
		return nil, err
	}
	log.Debug("Validator contract is getting initiated ", "contract", c)
	return &ValidatorSetContract{
		address: contractAddr,
		contract: c}, nil
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
	log.Debug("Getting block number to call method", "blockNumber", blockNumber, "signer", v.signer)
	validatorList, err := v.contract.ValidatorSetCaller.GetValidators(nil)
	if err != nil {
		log.Debug("Getting error while calling GetValidators method", "error", err)
		return nil
	}
	log.Debug("Calling validator list.", "validatorList", validatorList)
	v.ValidatorList = validatorList
	return v.ValidatorList
}

// System call - finalizeChange function
func (v *ValidatorSetContract) FinalizeChange(header *types.Header) (*types.Transaction, error) {
	log.Debug("Calling FinalizeChange method", "systemAddr", SYSTEM_ADDRESS, "signerAddr", v.signer, "ChainId", v.chainID)

	opts := &bind.TransactOpts{
		From: SYSTEM_ADDRESS,
		//Nonce: new(big.Int).SetUint64(header.Nonce.Uint64()),
		Signer: func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return v.signTxFn(accounts.Account{Address: v.signer}, tx, v.chainID)
		},
		//Value: nil,
		//GasPrice: big.NewInt(1),
		GasLimit: 9000000,
	}
	tx, err := v.contract.FinalizeChange(opts)
	if err != nil {
		log.Debug("Error occur while transact to finalizeChange method", "err", err)
		return nil, err
	}
	log.Debug("Getting transaction for finalizeChange method", "tx", tx, "err", err)
	return tx, err
}