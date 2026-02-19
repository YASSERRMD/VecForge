package db

import "time"

type Cache interface {
	Get(key string) ([]Hit, bool)
	Set(key string, hits []Hit, ttl time.Duration)
	Delete(key string)
	Clear()
}

type MemoryCache struct {
	data map[string]cacheEntry
}

type cacheEntry struct {
	hits    []Hit
	expires time.Time
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{data: make(map[string]cacheEntry)}
}

func (c *MemoryCache) Get(key string) ([]Hit, bool) {
	entry, ok := c.data[key]
	if !ok || time.Now().After(entry.expires) {
		return nil, false
	}
	return entry.hits, true
}

func (c *MemoryCache) Set(key string, hits []Hit, ttl time.Duration) {
	c.data[key] = cacheEntry{hits: hits, expires: time.Now().Add(ttl)}
}

func (c *MemoryCache) Delete(key string) {
	delete(c.data, key)
}

func (c *MemoryCache) Clear() {
	c.data = make(map[string]cacheEntry)
}
