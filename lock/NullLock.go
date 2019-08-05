package lock

type NullLock struct{}

func NewNullLock() *NullLock {
	return &NullLock{}
}

func (c *NullLock) TryAcquireLock(correlationId string,
	key string, ttl int) (bool, error) {
	return true, nil
}

func (c *NullLock) AcquireLock(correlationId string,
	key string, ttl int, timeout int) error {
	return nil
}

func (c *NullLock) ReleaseLock(correlationId string,
	key string) error {
	return nil
}
