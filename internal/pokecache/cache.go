package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	cacheMap map[string]cacheEntry
	interval time.Duration
	mu       sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(itv time.Duration) *Cache {
	m := make(map[string]cacheEntry)
	newCache := Cache{
		cacheMap: m,
		interval: itv,
	}
	ticker := time.NewTicker(itv)
	go func() {
		for {
			<-ticker.C
			newCache.Reaploop(itv)
		}
	}()

	return &newCache
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cacheMap[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}

	fmt.Printf("Add cache entry '%s', created at: %v\n", key, c.cacheMap[key].createdAt)
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, exists := c.cacheMap[key]

	return entry.val, exists
}

func (c *Cache) Reaploop(internval time.Duration) {
	for key, entry := range c.cacheMap {
		if time.Now().Sub(entry.createdAt) >= internval {
			delete(c.cacheMap, key)
		}
	}
}
