package pandora

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
	errc chan error
	res  chan<- [4]string //
}
