package lock

type ILock interface {
	TryAcquireLock(correlationId string, key string, ttl int64) (bool, error)

	AcquireLock(correlationId string, key string, ttl int64, timeout int64) error

	ReleaseLock(correlationId string, key string) error
}
