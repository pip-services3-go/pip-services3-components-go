package lock

type ILock interface {
	TryAcquireLock(correlationId string, key string, ttl int) (bool, error)

	AcquireLock(correlationId string, key string, ttl int, timeout int) error

	ReleaseLock(correlationId string, key string) error
}
