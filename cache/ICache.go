package cache

type ICache interface {
	Retrieve(correlationId string, key string) (interface{}, error)
	Store(correlationId string, key string, value interface{}, timeout int64) (interface{}, error)
	Remove(correlationId string, key string) error
}
