package pokecache

import (
	"fmt"
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
		fmt.Printf("Got %s from cache\n", key)
		return entry.val, true
	}
	return nil, false
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{createdAt: time.Now(), val: val}
	fmt.Printf("Added new %s to cache\n", key)
}

func NewCache(baseTime time.Duration) *Cache {
	fmt.Println("New Cache has been initialized")
	cache := Cache{entries: map[string]cacheEntry{}, mu: sync.Mutex{}}
	go reapLoop(baseTime, &cache)
	return &cache
}

func reapLoop(baseTime time.Duration, cache *Cache) {
	fmt.Println("cache cleaner initialized")
	ticker := time.NewTicker(baseTime)
	defer ticker.Stop()
	for range ticker.C {
		cache.mu.Lock()
		for key, entry := range cache.entries {
			if time.Since(entry.createdAt) >= baseTime {
				fmt.Printf("Cache deleted at %s\n", key)
				delete(cache.entries, key)
			}
		}
		cache.mu.Unlock()
	}
}
