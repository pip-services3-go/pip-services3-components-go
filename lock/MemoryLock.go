package lock

import (
	"sync"
	"time"
)

type MemoryLock struct {
	Lock
	mux   sync.Mutex
	locks map[string]time.Time
}

func NewMemoryLock() *MemoryLock {
	c := &MemoryLock{
		locks: map[string]time.Time{},
	}
	c.Lock = *InheritLock(c)

	return c
}

func (c *MemoryLock) TryAcquireLock(correlationId string,
	key string, ttl int64) (bool, error) {

	c.mux.Lock()
	defer c.mux.Unlock()

	expireTime, ok := c.locks[key]
	if ok {
		if expireTime.After(time.Now()) {
			return false, nil
		}
	}

	expireTime = time.Now().Add(time.Duration(ttl) * time.Millisecond)
	c.locks[key] = expireTime

	return true, nil
}

func (c *MemoryLock) ReleaseLock(correlationId string,
	key string) error {

	c.mux.Lock()
	defer c.mux.Unlock()

	delete(c.locks, key)

	return nil
}
