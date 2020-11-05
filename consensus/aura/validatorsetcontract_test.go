package aura

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"strings"
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

var testKey, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")

const abiJSONRelaySet = `[
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
    "inputs": [],
    "name": "finalizeChange",
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
    "constant": false,
    "inputs": [
      {
        "name": "_parentHash",
        "type": "bytes32"
      },
      {
        "name": "_newSet",
        "type": "address[]"
      }
    ],
    "name": "initiateChange",
    "outputs": [],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": true,
    "inputs": [],
    "name": "relayedSet",
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
        "name": "_relayedSet",
        "type": "address"
      }
    ],
    "name": "setRelayed",
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
    "inputs": [],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "constructor"
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
    "name": "NewRelayed",
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
  }
]`
const abiBinRelaySet = `0x6080604052336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555034801561005057600080fd5b5073fffffffffffffffffffffffffffffffffffffffe600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610ba5806100b56000396000f3006080604052600436106100a4576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806313af4035146100a957806375286211146100ec5780638da5cb5b14610103578063a0285c011461015a578063ae3783d6146101a3578063b7ab4db5146101fa578063bd96567714610266578063c476dd40146102a9578063d3e848f11461030e578063d69f13bb14610365575b600080fd5b3480156100b557600080fd5b506100ea600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506103b2565b005b3480156100f857600080fd5b506101016104cb565b005b34801561010f57600080fd5b506101186105c7565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561016657600080fd5b506101a160048036038101908080356000191690602001909291908035906020019082018035906020019190919293919293905050506105ec565b005b3480156101af57600080fd5b506101b86106a6565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561020657600080fd5b5061020f6106cc565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015610252578082015181840152602081019050610237565b505050509050019250505060405180910390f35b34801561027257600080fd5b506102a7600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506107ea565b005b3480156102b557600080fd5b5061030c600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001908201803590602001919091929391929390505050610905565b005b34801561031a57600080fd5b50610323610a3e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561037157600080fd5b506103b0600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610a64565b005b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561040d57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f70aea8d848e8a90fb7661b227dc522eb6395c3dac71b63cb59edd5c9899b236460405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561052757600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663752862116040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b1580156105ad57600080fd5b505af11580156105c1573d6000803e3d6000fd5b50505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561064857600080fd5b82600019167f55252fa6eee4741b4e24a74a70e9c11fd2c2281df8d6ea13126ff845f7825c898383604051808060200182810382528484828181526020019250602002808284378201915050935050505060405180910390a2505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6060600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b7ab4db56040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b15801561075457600080fd5b505af1158015610768573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250602081101561079257600080fd5b8101908080516401000000008111156107aa57600080fd5b828101905060208101848111156107c057600080fd5b81518560208202830111640100000000821117156107dd57600080fd5b5050929190505050905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561084557600080fd5b8073ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f4fea88aaf04c303804bb211ecc32a00ac8e5f0656bb854cad8a4a2e438256b7460405160405180910390a380600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635518edf833868686866040518663ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200184815260200180602001828103825284848281815260200192508082843782019150509650505050505050600060405180830381600087803b158015610a2057600080fd5b505af1158015610a34573d6000803e3d6000fd5b5050505050505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f981dfca3384846040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b158015610b5d57600080fd5b505af1158015610b71573d6000803e3d6000fd5b5050505050505600a165627a7a7230582006595124d53ed3ddda829a262c3e89584fa1ab6b358824d6e19bd1fa38c652840029`
const deployedCodeRelaySet = `6080604052600436106100a4576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806313af4035146100a957806375286211146100ec5780638da5cb5b14610103578063a0285c011461015a578063ae3783d6146101a3578063b7ab4db5146101fa578063bd96567714610266578063c476dd40146102a9578063d3e848f11461030e578063d69f13bb14610365575b600080fd5b3480156100b557600080fd5b506100ea600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506103b2565b005b3480156100f857600080fd5b506101016104cb565b005b34801561010f57600080fd5b506101186105c7565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561016657600080fd5b506101a160048036038101908080356000191690602001909291908035906020019082018035906020019190919293919293905050506105ec565b005b3480156101af57600080fd5b506101b86106a6565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561020657600080fd5b5061020f6106cc565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015610252578082015181840152602081019050610237565b505050509050019250505060405180910390f35b34801561027257600080fd5b506102a7600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506107ea565b005b3480156102b557600080fd5b5061030c600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001908201803590602001919091929391929390505050610905565b005b34801561031a57600080fd5b50610323610a3e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561037157600080fd5b506103b0600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610a64565b005b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561040d57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f70aea8d848e8a90fb7661b227dc522eb6395c3dac71b63cb59edd5c9899b236460405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561052757600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663752862116040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b1580156105ad57600080fd5b505af11580156105c1573d6000803e3d6000fd5b50505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561064857600080fd5b82600019167f55252fa6eee4741b4e24a74a70e9c11fd2c2281df8d6ea13126ff845f7825c898383604051808060200182810382528484828181526020019250602002808284378201915050935050505060405180910390a2505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6060600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b7ab4db56040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b15801561075457600080fd5b505af1158015610768573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250602081101561079257600080fd5b8101908080516401000000008111156107aa57600080fd5b828101905060208101848111156107c057600080fd5b81518560208202830111640100000000821117156107dd57600080fd5b5050929190505050905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561084557600080fd5b8073ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f4fea88aaf04c303804bb211ecc32a00ac8e5f0656bb854cad8a4a2e438256b7460405160405180910390a380600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635518edf833868686866040518663ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200184815260200180602001828103825284848281815260200192508082843782019150509650505050505050600060405180830381600087803b158015610a2057600080fd5b505af1158015610a34573d6000803e3d6000fd5b5050505050505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f981dfca3384846040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b158015610b5d57600080fd5b505af1158015610b71573d6000803e3d6000fd5b5050505050505600a165627a7a7230582006595124d53ed3ddda829a262c3e89584fa1ab6b358824d6e19bd1fa38c652840029`

const ValidatorSetABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"finalizeChange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"validator\",\"type\":\"address\"},{\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"reportMalicious\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"validator\",\"type\":\"address\"},{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"reportBenign\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_parentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_newSet\",\"type\":\"address[]\"}],\"name\":\"InitiateChange\",\"type\":\"event\"}]"
const abiJSONRelayedOwnedSet = `[
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
		"inputs": [
			{
				"name": "_reporter",
				"type": "address"
			},
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
		"name": "relayReportMalicious",
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
		"name": "relaySet",
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
				"name": "_relaySet",
				"type": "address"
			}
		],
		"name": "setRelay",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "_reporter",
				"type": "address"
			},
			{
				"name": "_validator",
				"type": "address"
			},
			{
				"name": "_blockNumber",
				"type": "uint256"
			}
		],
		"name": "relayReportBenign",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"name": "_relaySet",
				"type": "address"
			},
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
	}
]`
const abiBinRelayedOwnedSet = `0x6080604052336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060146001553480156200005657600080fd5b5060405162001bd938038062001bd983398101806040528101908080519060200190929190805182019291905050508060008160039080519060200190620000a0929190620001f0565b50600090505b815181101562000190576001600460008484815181101515620000c557fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160006101000a81548160ff021916908315150217905550806004600084848151811015156200013757fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055508080600101915050620000a6565b60036002908054620001a49291906200027f565b50505081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050506200031c565b8280548282559060005260206000209081019282156200026c579160200282015b828111156200026b5782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509160200191906001019062000211565b5b5090506200027b9190620002d6565b5090565b828054828255906000526020600020908101928215620002c35760005260206000209182015b82811115620002c2578254825591600101919060010190620002a5565b5b509050620002d29190620002d6565b5090565b6200031991905b808211156200031557600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550600101620002dd565b5090565b90565b6118ad806200032c6000396000f3006080604052600436106100d0576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630c857805146100d557806311ae9ed21461010057806313af40351461016c57806340a141ff146101af5780634d238c8e146101f25780635518edf81461023557806375286211146102ba5780638c90ac07146102d15780638da5cb5b146102fe578063a6940b0714610355578063b3f05b97146103ac578063b7ab4db5146103db578063c805f68b14610447578063f981dfca1461048a575b600080fd5b3480156100e157600080fd5b506100ea6104f7565b6040518082815260200191505060405180910390f35b34801561010c57600080fd5b506101156104fd565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561015857808201518184015260208101905061013d565b505050509050019250505060405180910390f35b34801561017857600080fd5b506101ad600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061058b565b005b3480156101bb57600080fd5b506101f0600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506106a4565b005b3480156101fe57600080fd5b50610233600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610a4d565b005b34801561024157600080fd5b506102b8600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001908201803590602001919091929391929390505050610c1e565b005b3480156102c657600080fd5b506102cf610cbf565b005b3480156102dd57600080fd5b506102fc60048036038101908080359060200190929190505050610d25565b005b34801561030a57600080fd5b50610313610d8a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561036157600080fd5b5061036a610daf565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156103b857600080fd5b506103c1610dd5565b604051808215151515815260200191505060405180910390f35b3480156103e757600080fd5b506103f0610de8565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015610433578082015181840152602081019050610418565b505050509050019250505060405180910390f35b34801561045357600080fd5b50610488600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610e76565b005b34801561049657600080fd5b506104f5600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610f15565b005b60015481565b6060600380548060200260200160405190810160405280929190818152602001828054801561058157602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610537575b5050505050905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156105e657600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f70aea8d848e8a90fb7661b227dc522eb6395c3dac71b63cb59edd5c9899b236460405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561070157600080fd5b81600080600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff169150600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015490508180156107ad575060028054905081105b801561081d57508273ffffffffffffffffffffffffffffffffffffffff166002828154811015156107da57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b151561082857600080fd5b600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101549350600360016003805490500381548110151561088457fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166003858154811015156108be57fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550836004600060038781548110151561091a57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010181905550600360016003805490500381548110151561099d57fe5b9060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905560038054809190600190036109df919061179b565b50600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600080820160006101000a81549060ff021916905560018201600090555050610a46610f81565b5050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610aa857600080fd5b80600460008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff16151515610b0557600080fd5b6001600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160006101000a81548160ff021916908315150217905550600380549050600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001018190555060038290806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610c1a610f81565b5050565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610c7a57600080fd5b610cb885858585858080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050610fc0565b5050505050565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610d1b57600080fd5b610d23611298565b565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610d8057600080fd5b8060018190555050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600060149054906101000a900460ff1681565b60606002805480602002602001604051908101604052809291908181526020018280548015610e6c57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610e22575b5050505050905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610ed157600080fd5b80600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610f7157600080fd5b610f7c838383611390565b505050565b600060149054906101000a900460ff161515610f9c57600080fd5b60008060146101000a81548160ff021916908315150217905550610fbe611667565b565b83600080600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff169150600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010154905081801561106c575060028054905081105b80156110dc57508273ffffffffffffffffffffffffffffffffffffffff1660028281548110151561109957fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b15156110e757600080fd5b85600080600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff169150600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101549050818015611193575060028054905081105b801561120357508273ffffffffffffffffffffffffffffffffffffffff166002828154811015156111c057fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b151561120e57600080fd5b876001548101431115801561122257504381105b151561122d57600080fd5b600115158a73ffffffffffffffffffffffffffffffffffffffff168c73ffffffffffffffffffffffffffffffffffffffff167f32c78b6140c46745a46e88cd883707d70dbd2f06d13dd76fe5f499c01290da4f60405160405180910390a45050505050505050505050565b600060149054906101000a900460ff161515156112b457600080fd5b600360029080546112c69291906117c7565b506001600060146101000a81548160ff0219169083151502179055507f8564cd629b15f47dc310d45bcbfc9bcf5420b0d51bf0659a16c67f91d276325360026040518080602001828103825283818154815260200191508054801561138057602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611336575b50509250505060405180910390a1565b82600080600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff169150600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010154905081801561143c575060028054905081105b80156114ac57508273ffffffffffffffffffffffffffffffffffffffff1660028281548110151561146957fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b15156114b757600080fd5b84600080600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff169150600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101549050818015611563575060028054905081105b80156115d357508273ffffffffffffffffffffffffffffffffffffffff1660028281548110151561159057fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b15156115de57600080fd5b86600154810143111580156115f257504381105b15156115fd57600080fd5b600015158973ffffffffffffffffffffffffffffffffffffffff168b73ffffffffffffffffffffffffffffffffffffffff167f32c78b6140c46745a46e88cd883707d70dbd2f06d13dd76fe5f499c01290da4f60405160405180910390a450505050505050505050565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a0285c01600143034060036040518363ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180836000191660001916815260200180602001828103825283818154815260200191508054801561176057602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611716575b50509350505050600060405180830381600087803b15801561178157600080fd5b505af1158015611795573d6000803e3d6000fd5b50505050565b8154818355818111156117c2578183600052602060002091820191016117c19190611819565b5b505050565b8280548282559060005260206000209081019282156118085760005260206000209182015b828111156118075782548255916001019190600101906117ec565b5b509050611815919061183e565b5090565b61183b91905b8082111561183757600081600090555060010161181f565b5090565b90565b61187e91905b8082111561187a57600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550600101611844565b5090565b905600a165627a7a7230582014fd93432d09bbfad462c8fedae2fae75c36043b2d1fbdbe38cb6a60ffaccbf50029`
const deployedCodeRelayedOwnedSet = ``


func simTestBackend(testAddr common.Address) *backends.SimulatedBackend {
	return backends.NewSimulatedBackend(
		core.GenesisAlloc{
			testAddr: {Balance: big.NewInt(10000000000)},
		}, 10000000,
	)
}

// Helper function to deploy contract
func deployTestContract(t *testing.T, abiJSON string, abiBin string, simBackend *backends.SimulatedBackend) (common.Address, *bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		t.Errorf("could not get code at test addr: %v", err)
	}

	auth := bind.NewKeyedTransactor(testKey)
	contractAddr, _, contract, err := bind.DeployContract(auth, parsed, common.FromHex(abiBin), simBackend)
	simBackend.Commit()
	return contractAddr, contract, err
}


func TestValidatorContractDeployment(t *testing.T) {
	deployerAddr := crypto.PubkeyToAddress(testKey.PublicKey)
	simBackend := simTestBackend(deployerAddr)
	defer simBackend.Close()

	relaySetContractAddr, relaysetContract, err := deployTestContract(t, abiJSONRelaySet, abiBinRelaySet, simBackend)
	if err != nil {
		t.Errorf("could not deploy err: %v contract: %v", err, relaysetContract)
	}

	bgCtx := context.Background()
	codeRelaySet, err := simBackend.CodeAt(bgCtx, relaySetContractAddr,nil)
	hexCodeRelaySet := hex.EncodeToString(codeRelaySet)
	if deployedCodeRelaySet != hexCodeRelaySet {
		t.Errorf("Deployed code mismatch. actualDeployedCode: %v expectedDeployedCode: %v", hexCodeRelaySet, deployedCodeRelaySet)
	}
}

func TestGetValidatorList(t *testing.T) {
	deployerAddr := crypto.PubkeyToAddress(testKey.PublicKey)
	simBackend := simTestBackend(deployerAddr)
	defer simBackend.Close()

	relaySetContractAddr, relaysetContract, err := deployTestContract(t, abiJSONRelaySet, abiBinRelaySet, simBackend)
	if err != nil {
		t.Errorf("could not deploy err: %v contract: %v", err, relaysetContract)
	}

	validatorSetContract, err := NewValidatorSet(relaySetContractAddr, simBackend)
	if err != nil {
		t.Errorf("could not initiate validator set contract err: %v contractAddress: %v", err, relaySetContractAddr)
	}

	validatorList := validatorSetContract.GetValidators(nil)
	if len(validatorList) != 0 {
		t.Errorf("wrong result found for validator list. err: %v contractAddress: %v", err, relaySetContractAddr)
	}

}

