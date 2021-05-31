/*
* Developed by: Md. Muhaimin Shah Pahalovi
* Generated: 5/31/21
* This file is generated to support Lukso pandora module.
* Purpose: orchestrator API needs request of some types and also replies response on some types. All the types of
	orchestraotr is mentioned here.
*/
package orcclient_pandora

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
