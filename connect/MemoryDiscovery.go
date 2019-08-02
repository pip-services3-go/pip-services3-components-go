package connect

import (
	"github.com/pip-services3-go/pip-services3-commons-go/config"
)

type MemoryDiscovery struct {
	items map[string][]*ConnectionParams
}

func NewEmptyMemoryDiscovery() *MemoryDiscovery {
	return &MemoryDiscovery{
		items: map[string][]*ConnectionParams{},
	}
}

func NewMemoryDiscovery(config *config.ConfigParams) *MemoryDiscovery {
	c := &MemoryDiscovery{
		items: map[string][]*ConnectionParams{},
	}

	if config != nil {
		c.Configure(config)
	}

	return c
}

func (c *MemoryDiscovery) Configure(config *config.ConfigParams) {
	c.ReadConnections(config)
}

func (c *MemoryDiscovery) ReadConnections(config *config.ConfigParams) {
	c.items = map[string][]*ConnectionParams{}

	keys := config.Keys()
	for _, key := range keys {
		value := config.GetAsString(key)
		connection := NewConnectionParamsFromString(value)
		c.items[key] = []*ConnectionParams{connection}
	}
}

func (c *MemoryDiscovery) Register(correlationId string, key string,
	connection *ConnectionParams) (result *ConnectionParams, err error) {

	if connection != nil {
		connections, _ := c.items[key]
		if connections == nil {
			connections = []*ConnectionParams{connection}
			c.items[key] = connections
		} else {
			connections = append(connections, connection)
		}
	}

	return connection, nil
}

func (c *MemoryDiscovery) ResolveOne(correlationId string,
	key string) (result *ConnectionParams, err error) {

	connections, _ := c.ResolveAll(correlationId, key)
	if len(connections) > 0 {
		return connections[0], nil
	}

	return nil, nil
}

func (c *MemoryDiscovery) ResolveAll(correlationId string,
	key string) (result []*ConnectionParams, err error) {
	connections, _ := c.items[key]

	if connections == nil {
		connections = []*ConnectionParams{}
	}

	return connections, nil
}
