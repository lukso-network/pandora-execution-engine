package pandora

import (
	lru "github.com/hashicorp/golang-lru"
	"sync"
)

const inmemoryEpochInfos = 4096 // Number of recent block signatures to keep in memory

type EpochInfoCache struct {
	epochInfos *lru.Cache // Signatures of recent blocks to speed up mining
	lock       sync.RWMutex
}

// NewPanHeaderCache initializes the map and underlying cache.
func NewEpochInfoCache() *EpochInfoCache {
	cache, err := lru.New(inmemoryEpochInfos)
	if err != nil {
		panic(err)
	}
	return &EpochInfoCache{
		epochInfos: cache,
	}
}

// Put
func (c *EpochInfoCache) put(epoch uint64, epochInfo *EpochInfo) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	copiedEpochInfo := copyEpochInfo(epochInfo)
	c.epochInfos.Add(epoch, copiedEpochInfo)
	return nil
}

// Get
func (c *EpochInfoCache) get(epoch uint64) (*EpochInfo, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	item, exists := c.epochInfos.Get(epoch)
	if exists && item != nil {
		epochInfo := item.(*EpochInfo)
		copiedEpochInfo := copyEpochInfo(epochInfo)
		return copiedEpochInfo, nil
	}
	return nil, errInvalidEpochInfo
}
