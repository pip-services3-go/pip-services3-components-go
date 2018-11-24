package connect

type IDiscovery interface {
	Register(correlationId string, key string, connection *ConnectionParams) (interface{}, error)
	ResolveOne(correlationId string, key string) (*ConnectionParams, error)
	ResolveAll(correlationId string, key string) ([]*ConnectionParams, error)
}
