package lock

import (
	"time"

	"github.com/pip-services3-go/pip-services3-commons-go/config"
	"github.com/pip-services3-go/pip-services3-commons-go/errors"
)

type Lock struct {
	retryTimeout int64
	locker       ILock
}

func InheritLock(locker ILock) *Lock {
	return &Lock{
		retryTimeout: 100,
		locker:       locker,
	}
}

func (c *Lock) Configure(config *config.ConfigParams) {
	c.retryTimeout = config.GetAsLongWithDefault("options.retry_timeout", c.retryTimeout)
}

func (c *Lock) AcquireLock(correlationId string,
	key string, ttl int64, timeout int64) error {

	expireTime := time.Now().Add(time.Duration(timeout) * time.Millisecond)

	// Repeat until time expires
	for time.Now().Before(expireTime) {
		// Try to get lock first
		locked, err := c.locker.TryAcquireLock(correlationId, key, ttl)
		if locked || err != nil {
			return err
		}

		// Sleep
		time.Sleep(time.Duration(c.retryTimeout) * time.Millisecond)
	}

	// Throw exception
	err := errors.NewConflictError(
		correlationId,
		"LOCK_TIMEOUT",
		"Acquiring lock "+key+" failed on timeout",
	).WithDetails("key", key)

	return err
}
