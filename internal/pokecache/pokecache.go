package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mu      sync.Mutex
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, exists := c.entries[key]
	if exists {
		return entry.val, true
	}
	return nil, false
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{createdAt: time.Now(), val: val}
}

func NewCache(baseTime time.Duration) *Cache {
	cache := Cache{entries: map[string]cacheEntry{}, mu: sync.Mutex{}}
	go reapLoop(baseTime, &cache)
	return &cache
}

func reapLoop(baseTime time.Duration, cache *Cache) {
	ticker := time.NewTicker(baseTime)
	defer ticker.Stop()
	for range ticker.C {
		cache.mu.Lock()
		for key, entry := range cache.entries {
			if time.Since(entry.createdAt) >= baseTime {
				delete(cache.entries, key)
			}
		}
		cache.mu.Unlock()
	}
}
