package connect

type IDiscovery interface {
	Register(correlationId string, key string,
		connection *ConnectionParams) (result *ConnectionParams, err error)

	ResolveOne(correlationId string, key string) (result *ConnectionParams, err error)

	ResolveAll(correlationId string, key string) (result []*ConnectionParams, err error)
}
