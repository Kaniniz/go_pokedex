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
	interval time.Duration //int64
	mu       *sync.RWMutex
	cache    map[string]cacheEntry
}

func (c Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
	return
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.cache[key]
	if ok {
		return value.val, ok
	}
	return nil, false
}

func (c Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for true {
		<-ticker.C
		c.mu.Lock()
		for key, entry := range c.cache {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.cache, key)
			}
		}
		c.mu.Unlock()
	}
}

func NewCache(inter time.Duration) Cache {
	cache := Cache{
		interval: inter,
		mu:       &sync.RWMutex{},
		cache:    map[string]cacheEntry{},
	}

	go cache.reapLoop()
	return cache
}
