package pokecache

import (
	"time"
	"sync"
)

type cacheEntry struct {
	createdAt	time.Time
	val			[]byte
}

type Cache struct {
	mu			sync.Mutex
	data 		map[string]cacheEntry
	interval	time.Duration
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	newEntry := cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.data[key] = newEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.data[key]
	return entry.val, ok
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		data:     make(map[string]cacheEntry),
    	interval: interval,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.reap()
	}
}

func (c *Cache) reap() {
	c.mu.Lock()
    defer c.mu.Unlock()
	for key, entry := range c.data {
		age := time.Since(entry.createdAt)
		if age > c.interval {
			delete(c.data, key)
		}
	}
}