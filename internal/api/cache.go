package api

import (
	"sync"
	"time"
)

type CacheEntry[V any] struct {
	Data     V
	CachedAt time.Time
}

type Cache[V any] struct {
	Cache      map[string]CacheEntry[V]
	mu         sync.Mutex
	CacheRenew time.Duration
}

func (cache *Cache[V]) IsCacheHit(steamid string) bool {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	entry, ok := cache.Cache[steamid]
	if !ok {
		return false
	}

	expired := time.Since(entry.CachedAt) >= cache.CacheRenew
	if !expired {
		entry.CachedAt = time.Now().UTC()
	}

	return !expired
}

func (cache *Cache[V]) UpdateCache(steamid string, data V) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	entry := cache.Cache[steamid]
	entry.Data = data
	entry.CachedAt = time.Now().UTC()
	cache.Cache[steamid] = entry
}
