package monnify

import (
	"sync"
	"time"
)

type Cache struct {
	data  map[string]CacheEntry
	mutex sync.RWMutex
}

type CacheEntry struct {
	value      string
	expiration time.Time
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]CacheEntry),
	}
}

func (c *Cache) Set(key string, value string, ttl time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	expiration := time.Now().Add(ttl)
	c.data[key] = CacheEntry{value, expiration}
}

func (c *Cache) Get(key string) (string, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	cacheEntry, ok := c.data[key]
	if !ok {
		return "", false
	}

	if time.Now().After(cacheEntry.expiration) {
		delete(c.data, key)
		return "", false
	}

	return cacheEntry.value, true
}

func (c *Cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.data, key)
}
