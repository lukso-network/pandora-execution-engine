package pandora

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"time"
)

const signatureSize = 96

type BlsSignatureBytes [signatureSize]byte

// mineResult wraps the pow solution parameters for the specified block.
type shardingResult struct {
	nonce   types.BlockNonce
	hash    common.Hash
	blsSeal *BlsSignatureBytes

	errc chan error
}

// EpochInfo
type EpochInfo struct {
	Epoch            uint64        `json:"epoch"`         // Epoch number
	ValidatorList    [32]string    `json:"validatorList"` // Validators public key list for specific epoch
	EpochTimeStart   uint64        `json:"epochTimeStart"`
	SlotTimeDuration time.Duration `json:"slotTimeDuration"`
}

// ExtraData
type ExtraData struct {
	Slot  uint64
	Epoch uint64
	Turn  uint64
}

// ExtraDataWithBLSSig
type ExtraDataWithBLSSig struct {
	ExtraData
}

// sealWork wraps a seal work package for remote sealer.
type shardingInfoReq struct {
	slot        uint64
	epoch       uint64
	blockNumber uint64
	parentHash  common.Hash

	errc chan error
	res  chan<- [4]string //
}


// sealTask wraps a seal block with relative result channel
type sealTask struct {
	block   *types.Block
	results chan<- *types.Block
}