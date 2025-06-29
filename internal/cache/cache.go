package cache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mutex *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cache: make(map[string]cacheEntry),
		mutex: &sync.Mutex{},
	}

	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, value []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	_, exists := c.cache[key]
	if exists {
		return
	}

	c.cache[key] = cacheEntry{createdAt: time.Now(), val: value}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	data, exists := c.cache[key]
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
		for key, val := range c.cache {
			age := time.Since(val.createdAt)
			if age > interval {
				delete(c.cache, key)
			}
		}
	}
}
