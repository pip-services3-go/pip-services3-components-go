package cache

type NullCache struct{}

func NewNullCache() *NullCache {
	return &NullCache{}
}

func (c *NullCache) Retrieve(correlationId string, key string) (interface{}, error) {
	return nil, nil
}

func (c *NullCache) Store(correlationId string, key string, value interface{}, timeout int64) (interface{}, error) {
	return value, nil
}

func (c *NullCache) Remove(correlationId string, key string) error {
	return nil
}
