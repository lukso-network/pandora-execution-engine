package pandora

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	bls_common "github.com/prysmaticlabs/prysm/shared/bls/common"
)

const signatureSize = 96

type BlsSignatureBytes [signatureSize]byte

type TransactionPool interface {
	AddLocal(tx *types.Transaction) error
	Has(hash common.Hash) bool
}

// mineResult wraps the pow solution parameters for the specified block.
type shardingResult struct {
	nonce   types.BlockNonce
	hash    common.Hash
	blsSeal *BlsSignatureBytes

	errc chan error
}

type SealHashUpdate struct {
	PreviousHash common.Hash
	UpdatedHash  common.Hash
}

type ExtraDataSealed struct {
	ExtraData
	BlsSignatureBytes *BlsSignatureBytes
}

// EpochInfo
type EpochInfo struct {
	Epoch            uint64
	ValidatorList    [32]bls_common.PublicKey
	EpochTimeStart   uint64
	SlotTimeDuration time.Duration
}

// Reorg holds reorg related information. Based on this info orchestrator can revert pandora blocks
type Reorg struct {
	VanParentHash []byte `json:"van_parent_hash"`
	PanParentHash []byte `json:"pan_parent_hash"`
	NewSlot       uint64 `json:"new_slot"`
}

type EpochInfoPayload struct {
	Epoch            uint64        `json:"epoch"`         // Epoch number
	ValidatorList    [32]string    `json:"validatorList"` // Validators public key list for specific epoch
	EpochTimeStart   uint64        `json:"epochTimeStart"`
	SlotTimeDuration time.Duration `json:"slotTimeDuration"`
	ReorgInfo        *Reorg        `json:"reorg_info"`
	FinalizedSlot    uint64        `json:"finalizedSlot"`
}

// ExtraData
type ExtraData struct {
	Slot  uint64
	Epoch uint64
	Turn  uint64
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

func (ei *EpochInfo) copy() *EpochInfo {
	return &EpochInfo{
		ei.Epoch,
		ei.ValidatorList,
		ei.EpochTimeStart,
		ei.SlotTimeDuration,
	}
}

// copyEpochInfo
func copyEpochInfo(ei *EpochInfo) *EpochInfo {
	cpy := *ei
	return &cpy
}
