package cache

import (
	"sync"
	"time"
)

type Cache struct {
	PokeCache map[string]cacheEntry
	Mutex     *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		PokeCache: make(map[string]cacheEntry),
		Mutex:     &sync.Mutex{},
	}

	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, value []byte) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	_, exists := c.PokeCache[key]
	if exists {
		return
	}

	c.PokeCache[key] = cacheEntry{createdAt: time.Now(), val: value}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	data, exists := c.PokeCache[key]
	if !exists {
		return []byte{}, false
	}

	if len(data.val) == 0 {
		return []byte{}, false
	}

	return data.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		for key, val := range c.PokeCache {
			age := time.Since(val.createdAt)
			if age > interval {
				delete(c.PokeCache, key)
			}
		}
	}
}
