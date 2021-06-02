package pandora_orcclient

import "github.com/ethereum/go-ethereum/common"

type Status int

type BlockHash struct {
	Slot uint64
	Hash common.Hash
}

type BlockStatus struct {
	BlockHash
	Status Status
}
