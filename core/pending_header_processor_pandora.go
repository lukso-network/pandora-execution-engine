package core

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	lru "github.com/hashicorp/golang-lru"
)

/*
* Generated: May - 05 - 2021
* This file is generated to support Lukso pandora module
* Purpose: In response of https://github.com/lukso-network/pandora-execution-engine/issues/57 we need to have a pending
* in memory database. Which will hold the headers when they are locally validated but not validated by orchestrator node.
* Insert Headers operation will halt until the header is validated by orchestrator.
 */

// PandoraPendingHeaderContainer will hold temporary headers in a in memory db.
type PandoraPendingHeaderContainer struct {
	headerContainer *lru.Cache // in-memory database which will hold headers temporarily
	pndHeaderFeed   event.Feed // announce new arrival of pending header
}

// WriteAndNotifyHeader writes header in database and also notify to the subscribers
func (container *PandoraPendingHeaderContainer) WriteAndNotifyHeader(header *types.Header) {
	// first send notification
	container.NotifyHeader(header)

	// then write into database
	container.WriteHeader(header)
}

func (container *PandoraPendingHeaderContainer) NotifyHeader(header *types.Header) {
	container.pndHeaderFeed.Send(PendingHeaderEvent{Headers: []*types.Header{header}})
}

// NewPandoraPendingHeaderContainer will return a fully initiated in-memory header container
func NewPandoraPendingHeaderContainer() *PandoraPendingHeaderContainer {
	// an arbitrary size of cache which will contain downloaded headers that are needed to be verified by orchestrator
	cache, err := lru.New(10)
	if err != nil {
		panic(err)
	}
	return &PandoraPendingHeaderContainer{
		headerContainer: cache,
	}
}

// WriteHeader dump a single header in the header container
func (container *PandoraPendingHeaderContainer) WriteHeader(header *types.Header) {
	// write the header into db
	container.headerContainer.Add(header.Hash(), header)
}

// DeleteHeader deletes a single header from the container
func (container *PandoraPendingHeaderContainer) DeleteHeader(header *types.Header) {
	container.headerContainer.Remove(header.Hash())
}

// ReadAllHeaders reads all the headers from the memory
func (container *PandoraPendingHeaderContainer) ReadAllHeaders() []*types.Header {

	// first retrieve the hashes of the headers
	values := container.headerContainer.Keys()
	var headers []*types.Header
	for _, data := range values {
		if headerData, ok := container.headerContainer.Get(data); ok {
			headers = append(headers, headerData.(*types.Header))
		}
	}

	return headers
}
