package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct{
	mu sync.Mutex
	data map[string]cacheEntry
	interval time.Duration
} 

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		data:     make(map[string]cacheEntry),
		interval: interval,
	}

	go c.reapLoop()

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = cacheEntry {
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.data[key] 
	if !ok{
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticket := time.NewTicker(c.interval)

	for range ticket.C {
		c.reap()
	}
}

func (c *Cache) reap() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()

	for key, entry := range c.data {
		if entry.createdAt.Before(now.Add(-c.interval)) {
			delete(c.data, key)
		}
	}
}
