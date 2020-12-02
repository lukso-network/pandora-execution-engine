package validatorset

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"math/big"
	"reflect"
	"strings"
	"testing"
)

// ownedSet contract is contract which implements validator set contract.
// validator set contract is an interface so for deployment we use owned set contract.
const ownedSetAbi = `[
	{
		"constant": true,
		"inputs": [],
		"name": "recentBlocks",
		"outputs": [
			{
				"name": "",
				"type": "uint256"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "getPending",
		"outputs": [
			{
				"name": "",
				"type": "address[]"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "_new",
				"type": "address"
			}
		],
		"name": "setOwner",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "_validator",
				"type": "address"
			}
		],
		"name": "removeValidator",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "_validator",
				"type": "address"
			}
		],
		"name": "addValidator",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [],
		"name": "finalizeChange",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "_recentBlocks",
				"type": "uint256"
			}
		],
		"name": "setRecentBlocks",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "owner",
		"outputs": [
			{
				"name": "",
				"type": "address"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "finalized",
		"outputs": [
			{
				"name": "",
				"type": "bool"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "getValidators",
		"outputs": [
			{
				"name": "",
				"type": "address[]"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "_validator",
				"type": "address"
			},
			{
				"name": "_blockNumber",
				"type": "uint256"
			},
			{
				"name": "_proof",
				"type": "bytes"
			}
		],
		"name": "reportMalicious",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "systemAddress",
		"outputs": [
			{
				"name": "",
				"type": "address"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "_validator",
				"type": "address"
			},
			{
				"name": "_blockNumber",
				"type": "uint256"
			}
		],
		"name": "reportBenign",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"name": "_initial",
				"type": "address[]"
			}
		],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "constructor"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"name": "currentSet",
				"type": "address[]"
			}
		],
		"name": "ChangeFinalized",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"name": "reporter",
				"type": "address"
			},
			{
				"indexed": true,
				"name": "reported",
				"type": "address"
			},
			{
				"indexed": true,
				"name": "malicious",
				"type": "bool"
			}
		],
		"name": "Report",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"name": "old",
				"type": "address"
			},
			{
				"indexed": true,
				"name": "current",
				"type": "address"
			}
		],
		"name": "NewOwner",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"name": "_parentHash",
				"type": "bytes32"
			},
			{
				"indexed": false,
				"name": "_newSet",
				"type": "address[]"
			}
		],
		"name": "InitiateChange",
		"type": "event"
	}
]`
const ownedSetBin = `0x6080604052336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060146001553480156200005657600080fd5b506040516200197e3803806200197e83398101806040528101908080518201929190505050806000816003908051906020019062000096929190620001f9565b50600090505b815181101562000186576001600460008484815181101515620000bb57fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160006101000a81548160ff021916908315150217905550806004600084848151811015156200012d57fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001018190555080806001019150506200009c565b600360029080546200019a92919062000288565b50505073fffffffffffffffffffffffffffffffffffffffe600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505062000325565b82805482825590600052602060002090810192821562000275579160200282015b82811115620002745782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550916020019190600101906200021a565b5b509050620002849190620002df565b5090565b828054828255906000526020600020908101928215620002cc5760005260206000209182015b82811115620002cb578254825591600101919060010190620002ae565b5b509050620002db9190620002df565b5090565b6200032291905b808211156200031e57600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550600101620002e6565b5090565b90565b61164980620003356000396000f3006080604052600436106100c5576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630c857805146100ca57806311ae9ed2146100f557806313af40351461016157806340a141ff146101a45780634d238c8e146101e7578063752862111461022a5780638c90ac07146102415780638da5cb5b1461026e578063b3f05b97146102c5578063b7ab4db5146102f4578063c476dd4014610360578063d3e848f1146103c5578063d69f13bb1461041c575b600080fd5b3480156100d657600080fd5b506100df610469565b6040518082815260200191505060405180910390f35b34801561010157600080fd5b5061010a61046f565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561014d578082015181840152602081019050610132565b505050509050019250505060405180910390f35b34801561016d57600080fd5b506101a2600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506104fd565b005b3480156101b057600080fd5b506101e5600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610616565b005b3480156101f357600080fd5b50610228600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506109bf565b005b34801561023657600080fd5b5061023f610b90565b005b34801561024d57600080fd5b5061026c60048036038101908080359060200190929190505050610bf6565b005b34801561027a57600080fd5b50610283610c5b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156102d157600080fd5b506102da610c80565b604051808215151515815260200191505060405180910390f35b34801561030057600080fd5b50610309610c93565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561034c578082015181840152602081019050610331565b505050509050019250505060405180910390f35b34801561036c57600080fd5b506103c3600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001908201803590602001919091929391929390505050610d21565b005b3480156103d157600080fd5b506103da610d65565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561042857600080fd5b50610467600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610d8b565b005b60015481565b606060038054806020026020016040519081016040528092919081815260200182805480156104f357602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116104a9575b5050505050905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561055857600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f70aea8d848e8a90fb7661b227dc522eb6395c3dac71b63cb59edd5c9899b236460405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561067357600080fd5b81600080600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff169150600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010154905081801561071f575060028054905081105b801561078f57508273ffffffffffffffffffffffffffffffffffffffff1660028281548110151561074c57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b151561079a57600080fd5b600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010154935060036001600380549050038154811015156107f657fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660038581548110151561083057fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550836004600060038781548110151561088c57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010181905550600360016003805490500381548110151561090f57fe5b9060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905560038054809190600190036109519190611537565b50600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600080820160006101000a81549060ff0219169055600182016000905550506109b8610d9a565b5050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610a1a57600080fd5b80600460008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff16151515610a7757600080fd5b6001600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160006101000a81548160ff021916908315150217905550600380549050600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001018190555060038290806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610b8c610d9a565b5050565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610bec57600080fd5b610bf4610dd9565b565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610c5157600080fd5b8060018190555050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600060149054906101000a900460ff1681565b60606002805480602002602001604051908101604052809291908181526020018280548015610d1757602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610ccd575b5050505050905090565b610d5f33858585858080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050610ed1565b50505050565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b610d963383836111a9565b5050565b600060149054906101000a900460ff161515610db557600080fd5b60008060146101000a81548160ff021916908315150217905550610dd7611480565b565b600060149054906101000a900460ff16151515610df557600080fd5b60036002908054610e07929190611563565b506001600060146101000a81548160ff0219169083151502179055507f8564cd629b15f47dc310d45bcbfc9bcf5420b0d51bf0659a16c67f91d2763253600260405180806020018281038252838181548152602001915080548015610ec157602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610e77575b50509250505060405180910390a1565b83600080600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff169150600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101549050818015610f7d575060028054905081105b8015610fed57508273ffffffffffffffffffffffffffffffffffffffff16600282815481101515610faa57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b1515610ff857600080fd5b85600080600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff169150600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015490508180156110a4575060028054905081105b801561111457508273ffffffffffffffffffffffffffffffffffffffff166002828154811015156110d157fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b151561111f57600080fd5b876001548101431115801561113357504381105b151561113e57600080fd5b600115158a73ffffffffffffffffffffffffffffffffffffffff168c73ffffffffffffffffffffffffffffffffffffffff167f32c78b6140c46745a46e88cd883707d70dbd2f06d13dd76fe5f499c01290da4f60405160405180910390a45050505050505050505050565b82600080600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff169150600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101549050818015611255575060028054905081105b80156112c557508273ffffffffffffffffffffffffffffffffffffffff1660028281548110151561128257fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b15156112d057600080fd5b84600080600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff169150600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010154905081801561137c575060028054905081105b80156113ec57508273ffffffffffffffffffffffffffffffffffffffff166002828154811015156113a957fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b15156113f757600080fd5b866001548101431115801561140b57504381105b151561141657600080fd5b600015158973ffffffffffffffffffffffffffffffffffffffff168b73ffffffffffffffffffffffffffffffffffffffff167f32c78b6140c46745a46e88cd883707d70dbd2f06d13dd76fe5f499c01290da4f60405160405180910390a450505050505050505050565b6001430340600019167f55252fa6eee4741b4e24a74a70e9c11fd2c2281df8d6ea13126ff845f7825c8960036040518080602001828103825283818154815260200191508054801561152757602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116114dd575b50509250505060405180910390a2565b81548183558181111561155e5781836000526020600020918201910161155d91906115b5565b5b505050565b8280548282559060005260206000209081019282156115a45760005260206000209182015b828111156115a3578254825591600101919060010190611588565b5b5090506115b191906115da565b5090565b6115d791905b808211156115d35760008160009055506001016115bb565b5090565b90565b61161a91905b8082111561161657600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055506001016115e0565b5090565b905600a165627a7a7230582080085f4c0e2873f17778383eeec85f5b76acb020ec57fe11eb0374494293495d0029`

var (
	contractAddress 	common.Address					// contract address which is used to interact with contract
	simulatedBackend 	*backends.SimulatedBackend		// contract backend
	initialValidators	[]common.Address				// initial validators which are set on contract initially
)

// init method deploy validator set contract and prepare all addresses which are interact with
// contract. Set deployed validator set contract address.
func init() {

	validatorKey1, _ := crypto.GenerateKey()
	validator1Addr := crypto.PubkeyToAddress(validatorKey1.PublicKey)
	validatorKey2, _ := crypto.GenerateKey()
	validator2Addr := crypto.PubkeyToAddress(validatorKey2.PublicKey)

	initialValidators = append(initialValidators, validator1Addr, validator2Addr)

	deployerKey, _ := crypto.GenerateKey()
	deployerAddr := crypto.PubkeyToAddress(deployerKey.PublicKey)
	opts := bind.NewKeyedTransactor(deployerKey)

	sim := backends.NewSimulatedBackend(
		core.GenesisAlloc{
			deployerAddr: {Balance: big.NewInt(params.Ether)},
			validator1Addr: {Balance: big.NewInt(params.Ether)},
			validator2Addr: {Balance: big.NewInt(params.Ether)},
		}, 10000000)
	defer sim.Close()

	parsed, _ := abi.JSON(strings.NewReader(ownedSetAbi))
	contractAddr, _, _, _ := bind.DeployContract(opts, parsed, common.FromHex(ownedSetBin), sim, initialValidators)
	sim.Commit()

	contractAddress = contractAddr
	simulatedBackend = sim
}

func TestNewValidatorSet(t *testing.T)  {
	validatorSetJSON, err := ioutil.ReadFile("res/validatorset.json")
	if err != nil {
		t.Fatalf("could not read file: %v", err)
	}

	var authority params.ValidatorSet
	err = json.Unmarshal(validatorSetJSON, &authority)
	if err != nil {
		t.Errorf("could not get code at test addr: %v", err)
	}

	multiMap := make(map[int]ValidatorSet)
	validators := NewValidatorSet(multiMap, &authority)
	validatorSet := validators.GetValidatorsByCaller(20)

	assert.Equal(t, validatorSet[0], common.HexToAddress("0x45d9d2cd449a754c494264e1809c50e34d64562b"))
}

// Test weather contract deployed successfully and get initial validator list from
// contract which is passed in constructor and also checks weather or not the NewValidatorSetWithSimBackend
// setup the configuration correctly.
func TestValidatorContract_GetValidatorList(t *testing.T) {
	validatorSetContract := NewValidatorSafeContract(contractAddress)
	validatorSetContract.PrepareBackend(0, simulatedBackend)

	validatorList := validatorSetContract.GetValidatorsByCaller(0)
	if reflect.DeepEqual(validatorList, initialValidators) == false {
		t.Errorf("mismatch in initial validator list. expected: %v actual: %v", initialValidators, validatorSetContract)
	}
}

// Test scenario - deploy contract and call finalizeChange and if the call successful then
// the transaction is not reverted.
func TestVaidatorContract_FinalizeChange(t *testing.T) {
	validatorSetContract := NewValidatorSafeContract(contractAddress)
	validatorSetContract.PrepareBackend(0, simulatedBackend)

	if err  := validatorSetContract.FinalizeChange(nil, nil); err != nil {
		t.Errorf("got error when call finalizeChange method. expected: %v actual: %v", nil, err)
	}
}

