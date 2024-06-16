package api

import (
	"log"
	"sync"
	"time"
)

type CacheEntry[V any] struct {
	Data     V
	CachedAt time.Time
}

func (entry CacheEntry[V]) IsExpired(renewTime time.Duration) bool {
	return time.Since(entry.CachedAt) >= renewTime
}

type Cache[V any] struct {
	Cache      map[string]CacheEntry[V]
	mu         sync.RWMutex
	CacheRenew time.Duration
}

func (cache *Cache[V]) IsCacheHit(steamid string) bool {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	entry, ok := cache.Cache[steamid]
	if !ok {
		return false
	}

	expired := entry.IsExpired(cache.CacheRenew)
	if !expired {
		entry.CachedAt = time.Now().UTC()
		cache.Cache[steamid] = entry
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

func (cache *Cache[V]) ReadCache(steamid string) V {
	cache.mu.RLock()
	defer cache.mu.RUnlock()

	return cache.Cache[steamid].Data
}

func (cache *Cache[V]) CleanExpiredEntries() int {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	removed := 0

	for id, entry := range cache.Cache {
		if entry.IsExpired(cache.CacheRenew) {
			log.Println("Removing expired entry for id:", id)
			delete(cache.Cache, id)
			removed++
		}
	}

	return removed
}

type Cleaner[V any] struct {
	Name     string
	ticker   *time.Ticker
	done     chan bool
	Cache    *Cache[V]
	Interval time.Duration
}

func (cleaner *Cleaner[V]) Start() {
	cleaner.ticker = time.NewTicker(cleaner.Interval)

	go func() {
		for {
			select {
			case <-cleaner.done:
				return
			case <-cleaner.ticker.C:
				log.Printf("Running cleaner for cache '%s'\n", cleaner.Name)
				removed := cleaner.Cache.CleanExpiredEntries()
				log.Printf("Removed %d expired entries from cache '%s'\n", removed, cleaner.Name)
			}
		}
	}()
	log.Printf("Started cleaner for cache '%s' with interval %v\n", cleaner.Name, cleaner.Interval)
}

func (cleaner *Cleaner[V]) Stop() {
	cleaner.ticker.Stop()
	cleaner.done <- true
	log.Printf("Stopping cleaner for cache '%s'\n", cleaner.Name)
}
