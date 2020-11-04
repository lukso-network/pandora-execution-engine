package aura

import (
	"bytes"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/aura/validatorset"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"sort"
	"testing"
)

type Account struct {
	key  *ecdsa.PrivateKey
	addr common.Address
}
type Accounts []Account

func (a Accounts) Len() int           { return len(a) }
func (a Accounts) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Accounts) Less(i, j int) bool { return bytes.Compare(a[i].addr.Bytes(), a[j].addr.Bytes()) < 0 }

func TestRelaySetDeployment(t *testing.T) {
	// Initialize test accounts
	var accounts Accounts
	for i := 0; i < 3; i++ {
		key, _ := crypto.GenerateKey()
		addr := crypto.PubkeyToAddress(key.PublicKey)
		accounts = append(accounts, Account{key: key, addr: addr})
	}
	sort.Sort(accounts)

	// Deploy registrar contract
	contractBackend := backends.NewSimulatedBackend(core.GenesisAlloc{accounts[0].addr: {Balance: big.NewInt(1000000000)}, accounts[1].addr: {Balance: big.NewInt(1000000000)}, accounts[2].addr: {Balance: big.NewInt(1000000000)}}, 10000000)
	defer contractBackend.Close()

	transactOpts := bind.NewKeyedTransactor(accounts[0].key)

	// 3 trusted signers, threshold 2
	contractAddr, _, _, err := validatorset.DeployOwned(transactOpts, contractBackend)
	if err != nil {
		t.Error("Failed to deploy registrar contract", err)
	}
	contractBackend.Commit()
	t.Log("Getting contract address", contractAddr)
}