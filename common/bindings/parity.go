package bindings

import (
	"encoding/binary"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/params"
	"math"
	"math/big"
)

// ParityChainSpec is the chain specification format used by Parity.
type ParityChainSpec struct {
	Name   string `json:"name"`
	Engine struct {
		Ethash *Ethash `json:"ethash,omitempty"`
		AuthorityRound *AuthorityRound `json:"authorityRound,omitempty"`
	} `json:"engine"`

	Params struct {
		MaximumExtraDataSize hexutil.Uint64 `json:"maximumExtraDataSize"`
		MinGasLimit          hexutil.Uint64 `json:"minGasLimit"`
		GasLimitBoundDivisor hexutil.Uint64 `json:"gasLimitBoundDivisor"`
		NetworkID            hexutil.Uint64 `json:"networkID"`
		MaxCodeSize          *big.Int       `json:"maxCodeSize"`
		EIP155Transition     *big.Int       `json:"eip155Transition, omitempty"`
		EIP98Transition      *big.Float     `json:"eip98Transition, omitempty"`
		EIP140Transition     *big.Int       `json:"eip140Transition, omitempty"`
		EIP211Transition     *big.Int       `json:"eip211Transition, omitempty"`
		EIP214Transition     *big.Int       `json:"eip214Transition, omitempty"`
		EIP658Transition     *big.Int       `json:"eip658Transition, omitempty"`
	} `json:"params"`

	Genesis struct {
		Seal struct {
			Ethereum struct {
				Nonce   hexutil.Bytes `json:"nonce"`
				MixHash hexutil.Bytes `json:"mixHash"`
			} `json:"ethereum"`
		} `json:"seal"`

		Difficulty *hexutil.Big   `json:"difficulty"`
		Author     common.Address `json:"author"`
		Timestamp  hexutil.Uint64 `json:"timestamp"`
		ParentHash common.Hash    `json:"parentHash"`
		ExtraData  hexutil.Bytes  `json:"extraData"`
		GasLimit   hexutil.Uint64 `json:"gasLimit"`
	} `json:"genesis"`

	Nodes    []string                                   `json:"nodes"`
	Accounts map[common.Address]*parityChainSpecAccount `json:"accounts"`
}

type Ethash struct {
	Params struct {
		MinimumDifficulty      *hexutil.Big `json:"minimumDifficulty"`
		DifficultyBoundDivisor *hexutil.Big `json:"difficultyBoundDivisor"`
		DurationLimit          *hexutil.Big `json:"durationLimit"`
		BlockReward            *hexutil.Big `json:"blockReward"`
		HomesteadTransition    *big.Int     `json:"homesteadTransition, omitempty"`
		EIP150Transition       *big.Int     `json:"eip150Transition, omitempty"`
		EIP160Transition       *big.Int     `json:"eip160Transition, omitempty"`
		EIP161abcTransition    *big.Int     `json:"eip161abcTransition, omitempty"`
		EIP161dTransition      *big.Int     `json:"eip161dTransition, omitempty"`
		EIP649Reward           *hexutil.Big `json:"eip649Reward, omitempty"`
		EIP100bTransition      *big.Int     `json:"eip100bTransition, omitempty"`
		EIP649Transition       *big.Int     `json:"eip649Transition, omitempty"`
	} `json:"params"`
}

type AuthorityRound struct {
	Params struct {
		StepDuration uint64 `json:"stepDuration, omitempty"`
		Validators   struct {
			List []common.Address `json:"list, omitempty"`
		} `json:"validators, omitempty"`
	} `json:"params, omitempty"`
}

// parityChainSpecAccount is the prefunded genesis account and/or precompiled
// contract definition.
type parityChainSpecAccount struct {
	Balance *hexutil.Big            `json:"balance"`
	Nonce   uint64                  `json:"nonce,omitempty"`
	Builtin *parityChainSpecBuiltin `json:"builtin,omitempty"`
}

// parityChainSpecBuiltin is the precompiled contract definition.
type parityChainSpecBuiltin struct {
	Name       string                  `json:"name,omitempty"`
	ActivateAt uint64                  `json:"activate_at,omitempty"`
	Pricing    *parityChainSpecPricing `json:"pricing,omitempty"`
}

// parityChainSpecPricing represents the different pricing models that builtin
// contracts might advertise using.
type parityChainSpecPricing struct {
	Linear       *parityChainSpecLinearPricing       `json:"linear,omitempty"`
	ModExp       *parityChainSpecModExpPricing       `json:"modexp,omitempty"`
	AltBnPairing *parityChainSpecAltBnPairingPricing `json:"alt_bn128_pairing,omitempty"`
}

type parityChainSpecLinearPricing struct {
	Base uint64 `json:"base"`
	Word uint64 `json:"word"`
}

type parityChainSpecModExpPricing struct {
	Divisor uint64 `json:"divisor"`
}

type parityChainSpecAltBnPairingPricing struct {
	Base uint64 `json:"base"`
	Pair uint64 `json:"pair"`
}

// newParityChainSpec converts a go-ethereum genesis block into a Parity specific
// chain specification format.
func NewParityChainSpec(network string, genesis *core.Genesis, bootnodes []string) (*ParityChainSpec, error) {
	// Only ethash is currently supported between go-ethereum and Parity
	if genesis.Config.Ethash == nil && nil == genesis.Config.Aura {
		return nil, errors.New("unsupported consensus engine")
	}
	// Reconstruct the chain spec in Parity's format
	spec := &ParityChainSpec{
		Name:  network,
		Nodes: bootnodes,
	}

	if nil != genesis.Config.Ethash {
		spec.Engine.Ethash = &Ethash{}
		spec.Engine.Ethash.Params.MinimumDifficulty = (*hexutil.Big)(params.MinimumDifficulty)
		spec.Engine.Ethash.Params.DifficultyBoundDivisor = (*hexutil.Big)(params.DifficultyBoundDivisor)
		spec.Engine.Ethash.Params.DurationLimit = (*hexutil.Big)(params.DurationLimit)
		spec.Engine.Ethash.Params.BlockReward = (*hexutil.Big)(ethash.FrontierBlockReward)
		spec.Engine.Ethash.Params.HomesteadTransition = genesis.Config.HomesteadBlock
		spec.Engine.Ethash.Params.EIP150Transition = genesis.Config.EIP150Block
		spec.Engine.Ethash.Params.EIP160Transition = genesis.Config.EIP155Block
		spec.Engine.Ethash.Params.EIP161abcTransition = genesis.Config.EIP158Block
		spec.Engine.Ethash.Params.EIP161dTransition = genesis.Config.EIP158Block
		spec.Engine.Ethash.Params.EIP649Reward = (*hexutil.Big)(ethash.ByzantiumBlockReward)
		spec.Engine.Ethash.Params.EIP100bTransition = genesis.Config.ByzantiumBlock
		spec.Engine.Ethash.Params.EIP649Transition = genesis.Config.ByzantiumBlock
	}

	if nil != genesis.Config.Aura {
		spec.Engine.AuthorityRound = &AuthorityRound{}
		authorityRoundEngine := spec.Engine.AuthorityRound
		authorityRoundEngine.Params.Validators.List = genesis.Config.Aura.Authorities
		authorityRoundEngine.Params.StepDuration = genesis.Config.Aura.Period
	}

	spec.Params.MaximumExtraDataSize = (hexutil.Uint64)(params.MaximumExtraDataSize)
	spec.Params.MinGasLimit = (hexutil.Uint64)(params.MinGasLimit)
	spec.Params.GasLimitBoundDivisor = (hexutil.Uint64)(params.GasLimitBoundDivisor)
	spec.Params.NetworkID = (hexutil.Uint64)(genesis.Config.ChainID.Uint64())
	spec.Params.MaxCodeSize = big.NewInt(params.MaxCodeSize)
	if nil != genesis.Config.EIP155Block {
		spec.Params.EIP155Transition = genesis.Config.EIP155Block
	}
	if nil != genesis.Config.Ethash {
		spec.Params.EIP98Transition = big.NewFloat(math.MaxUint64)
	}
	if nil != genesis.Config.ByzantiumBlock {
		spec.Params.EIP140Transition = genesis.Config.ByzantiumBlock
		spec.Params.EIP211Transition = genesis.Config.ByzantiumBlock
		spec.Params.EIP214Transition = genesis.Config.ByzantiumBlock
		spec.Params.EIP658Transition = genesis.Config.ByzantiumBlock
	}

	spec.Genesis.Seal.Ethereum.Nonce = (hexutil.Bytes)(make([]byte, 8))
	binary.LittleEndian.PutUint64(spec.Genesis.Seal.Ethereum.Nonce[:], genesis.Nonce)

	spec.Genesis.Seal.Ethereum.MixHash = (hexutil.Bytes)(genesis.Mixhash[:])
	spec.Genesis.Difficulty = (*hexutil.Big)(genesis.Difficulty)
	spec.Genesis.Author = genesis.Coinbase
	spec.Genesis.Timestamp = (hexutil.Uint64)(genesis.Timestamp)
	spec.Genesis.ParentHash = genesis.ParentHash
	spec.Genesis.ExtraData = hexutil.Bytes{}

	if nil != genesis.ExtraData {
		spec.Genesis.ExtraData = genesis.ExtraData
	}

	spec.Genesis.GasLimit = (hexutil.Uint64)(genesis.GasLimit)

	spec.Accounts = make(map[common.Address]*parityChainSpecAccount)
	for address, account := range genesis.Alloc {
		spec.Accounts[address] = &parityChainSpecAccount{
			Balance: (*hexutil.Big)(account.Balance),
			Nonce:   account.Nonce,
		}
	}
	spec.Accounts[common.BytesToAddress([]byte{1})].Builtin = &parityChainSpecBuiltin{
		Name: "ecrecover", Pricing: &parityChainSpecPricing{Linear: &parityChainSpecLinearPricing{Base: 3000}},
	}
	spec.Accounts[common.BytesToAddress([]byte{2})].Builtin = &parityChainSpecBuiltin{
		Name: "sha256", Pricing: &parityChainSpecPricing{Linear: &parityChainSpecLinearPricing{Base: 60, Word: 12}},
	}
	spec.Accounts[common.BytesToAddress([]byte{3})].Builtin = &parityChainSpecBuiltin{
		Name: "ripemd160", Pricing: &parityChainSpecPricing{Linear: &parityChainSpecLinearPricing{Base: 600, Word: 120}},
	}
	spec.Accounts[common.BytesToAddress([]byte{4})].Builtin = &parityChainSpecBuiltin{
		Name: "identity", Pricing: &parityChainSpecPricing{Linear: &parityChainSpecLinearPricing{Base: 15, Word: 3}},
	}
	if genesis.Config.ByzantiumBlock != nil {
		spec.Accounts[common.BytesToAddress([]byte{5})].Builtin = &parityChainSpecBuiltin{
			Name: "modexp", ActivateAt: genesis.Config.ByzantiumBlock.Uint64(), Pricing: &parityChainSpecPricing{ModExp: &parityChainSpecModExpPricing{Divisor: 20}},
		}
		spec.Accounts[common.BytesToAddress([]byte{6})].Builtin = &parityChainSpecBuiltin{
			Name: "alt_bn128_add", ActivateAt: genesis.Config.ByzantiumBlock.Uint64(), Pricing: &parityChainSpecPricing{Linear: &parityChainSpecLinearPricing{Base: 500}},
		}
		spec.Accounts[common.BytesToAddress([]byte{7})].Builtin = &parityChainSpecBuiltin{
			Name: "alt_bn128_mul", ActivateAt: genesis.Config.ByzantiumBlock.Uint64(), Pricing: &parityChainSpecPricing{Linear: &parityChainSpecLinearPricing{Base: 40000}},
		}
		spec.Accounts[common.BytesToAddress([]byte{8})].Builtin = &parityChainSpecBuiltin{
			Name: "alt_bn128_pairing", ActivateAt: genesis.Config.ByzantiumBlock.Uint64(), Pricing: &parityChainSpecPricing{AltBnPairing: &parityChainSpecAltBnPairingPricing{Base: 100000, Pair: 80000}},
		}
	}

	return spec, nil
}
