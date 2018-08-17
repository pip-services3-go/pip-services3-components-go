package cache

import (
	"sync"

	"github.com/pip-services-go/pip-services-commons-go/config"
)

type MemoryCache struct {
	cache   map[string]*CacheEntry
	lock    *sync.Mutex
	timeout int64
	maxSize int
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		cache:   map[string]*CacheEntry{},
		lock:    &sync.Mutex{},
		timeout: 60000,
		maxSize: 1000,
	}
}

func NewMemoryCacheFromConfig(cfg *config.ConfigParams) *MemoryCache {
	c := NewMemoryCache()
	c.Configure(cfg)
	return c
}

func (c *MemoryCache) Configure(cfg *config.ConfigParams) {
	c.timeout = cfg.GetAsLongWithDefault("timeout", c.timeout)
	c.maxSize = cfg.GetAsIntegerWithDefault("max_size", c.maxSize)
}

func (c *MemoryCache) Cleanup() {
	var oldest *CacheEntry
	var keysToRemove = []string{}

	c.lock.Lock()
	defer c.lock.Unlock()

	for key, value := range c.cache {
		if value.IsExpired() {
			keysToRemove = append(keysToRemove, key)
		}
		if oldest == nil || oldest.Expiration().After(value.Expiration()) {
			oldest = value
		}
	}

	for _, key := range keysToRemove {
		delete(c.cache, key)
	}

	if len(c.cache) > c.maxSize && oldest != nil {
		delete(c.cache, oldest.Key())
	}
}

func (c *MemoryCache) Retrieve(correlationId string, key string) (interface{}, error) {
	if key == "" {
		panic("Key cannot be empty")
	}

	c.lock.Lock()
	defer c.lock.Unlock()

	entry := c.cache[key]
	if entry != nil {
		if entry.IsExpired() {
			delete(c.cache, key)
			return nil, nil
		}

		return entry.Value(), nil
	}

	return nil, nil
}

func (c *MemoryCache) Store(correlationId string, key string, value interface{}, timeout int64) (interface{}, error) {
	if key == "" {
		panic("Key cannot be empty")
	}

	c.lock.Lock()
	defer c.lock.Unlock()

	entry := c.cache[key]
	if timeout <= 0 {
		timeout = c.timeout
	}

	// if value == nil {
	// 	if entry != nil {
	// 		delete(c.cache, key)
	// 	}
	// 	return nil, nil
	// }

	if entry != nil {
		entry.SetValue(value, timeout)
	} else {
		c.cache[key] = NewCacheEntry(key, value, timeout)
	}

	// cleanup
	if c.maxSize > 0 && len(c.cache) > c.maxSize {
		c.Cleanup()
	}

	return value, nil
}

func (c *MemoryCache) Remove(correlationId string, key string) error {
	if key == "" {
		panic("Key cannot be empty")
	}

	c.lock.Lock()
	defer c.lock.Unlock()

	delete(c.cache, key)

	return nil
}

func (c *MemoryCache) Clear(correlationId string) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.cache = map[string]*CacheEntry{}

	return nil
}
