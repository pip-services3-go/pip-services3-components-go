package cache

import (
	"time"
)

type CacheEntry struct {
	key        string
	value      interface{}
	expiration time.Time
}

func NewCacheEntry(key string, value interface{}, timeout int64) *CacheEntry {
	return &CacheEntry{
		key:        key,
		value:      value,
		expiration: time.Now().Add(time.Duration(timeout) * time.Millisecond),
	}
}

func (c *CacheEntry) Key() string {
	return c.key
}

func (c *CacheEntry) Value() interface{} {
	return c.value
}

func (c *CacheEntry) Expiration() time.Time {
	return c.expiration
}

func (c *CacheEntry) SetValue(value interface{}, timeout int64) {
	c.value = value
	c.expiration = time.Now().Add(time.Duration(timeout) * time.Millisecond)
}

func (c *CacheEntry) IsExpired() bool {
	return time.Now().After(c.expiration)
}
