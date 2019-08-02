package connect

import (
	"github.com/pip-services3-go/pip-services3-commons-go/config"
	"github.com/pip-services3-go/pip-services3-commons-go/refer"
)

type ConnectionResolver struct {
	connections []*ConnectionParams
	references  refer.IReferences
}

func NewEmptyConnectionResolver() *ConnectionResolver {
	return &ConnectionResolver{
		connections: []*ConnectionParams{},
		references:  nil,
	}
}

func NewConnectionResolver(config *config.ConfigParams,
	references refer.IReferences) *ConnectionResolver {
	c := &ConnectionResolver{
		connections: []*ConnectionParams{},
		references:  references,
	}

	if config != nil {
		c.Configure(config)
	}

	return c
}

func (c *ConnectionResolver) Configure(config *config.ConfigParams) {
	connections := NewManyConnectionParamsFromConfig(config)

	for _, connection := range connections {
		c.connections = append(c.connections, connection)
	}
}

func (c *ConnectionResolver) SetReferences(references refer.IReferences) {
	c.references = references
}

func (c *ConnectionResolver) GetAll() []*ConnectionParams {
	return c.connections
}

func (c *ConnectionResolver) Add(connection *ConnectionParams) {
	c.connections = append(c.connections, connection)
}

func (c *ConnectionResolver) resolveInDiscovery(correlationId string,
	connection *ConnectionParams) (result *ConnectionParams, err error) {

	if !connection.UseDiscovery() {
		return connection, nil
	}

	key := connection.DiscoveryKey()
	if c.references == nil {
		return nil, nil
	}

	discoveryDescriptor := refer.NewDescriptor("*", "discovery", "*", "*", "*")
	components := c.references.GetOptional(discoveryDescriptor)
	if len(components) == 0 {
		err := refer.NewReferenceError(correlationId, discoveryDescriptor)
		return nil, err
	}

	for _, component := range components {
		discovery, _ := component.(IDiscovery)
		if discovery != nil {
			connection, err = discovery.ResolveOne(correlationId, key)
			if connection != nil || err != nil {
				return connection, err
			}
		}
	}

	return nil, nil
}

func (c *ConnectionResolver) Resolve(correlationId string) (*ConnectionParams, error) {
	if len(c.connections) == 0 {
		return nil, nil
	}

	resolveConnections := []*ConnectionParams{}

	for _, connection := range c.connections {
		if !connection.UseDiscovery() {
			return connection, nil
		}

		resolveConnections = append(resolveConnections, connection)
	}

	for _, connection := range resolveConnections {
		c, err := c.resolveInDiscovery(correlationId, connection)
		if c != nil || err != nil {
			return c, err
		}
	}

	return nil, nil
}

func (c *ConnectionResolver) resolveAllInDiscovery(correlationId string,
	connection *ConnectionParams) (result []*ConnectionParams, err error) {

	if !connection.UseDiscovery() {
		return []*ConnectionParams{connection}, nil
	}

	key := connection.DiscoveryKey()
	if c.references == nil {
		return nil, nil
	}

	discoveryDescriptor := refer.NewDescriptor("*", "discovery", "*", "*", "*")
	components := c.references.GetOptional(discoveryDescriptor)
	if len(components) == 0 {
		err := refer.NewReferenceError(correlationId, discoveryDescriptor)
		return nil, err
	}

	resolvedConnections := []*ConnectionParams{}

	for _, component := range components {
		discovery, _ := component.(IDiscovery)
		if discovery != nil {
			connections, err := discovery.ResolveAll(correlationId, key)
			if err != nil {
				return nil, err
			}
			if connections != nil {
				for _, c := range connections {
					resolvedConnections = append(resolvedConnections, c)
				}
			}
		}
	}

	return resolvedConnections, nil
}

func (c *ConnectionResolver) ResolveAll(correlationId string) ([]*ConnectionParams, error) {
	resolvedConnections := []*ConnectionParams{}
	resolveConnections := []*ConnectionParams{}

	for _, connection := range c.connections {
		if !connection.UseDiscovery() {
			resolvedConnections = append(resolvedConnections, connection)
		} else {
			resolveConnections = append(resolveConnections, connection)
		}
	}

	for _, connection := range resolveConnections {
		connections, err := c.resolveAllInDiscovery(correlationId, connection)
		if err != nil {
			return nil, err
		}
		for _, c := range connections {
			resolvedConnections = append(resolvedConnections, c)
		}
	}

	return resolvedConnections, nil
}

func (c *ConnectionResolver) registerInDiscovery(correlationId string,
	connection *ConnectionParams) (result bool, err error) {

	if !connection.UseDiscovery() {
		return false, nil
	}

	key := connection.DiscoveryKey()
	if c.references == nil {
		return false, nil
	}

	discoveryDescriptor := refer.NewDescriptor("*", "discovery", "*", "*", "*")
	components := c.references.GetOptional(discoveryDescriptor)
	if len(components) == 0 {
		err := refer.NewReferenceError(correlationId, discoveryDescriptor)
		return false, err
	}

	registered := false

	for _, component := range components {
		discovery, _ := component.(IDiscovery)
		if discovery != nil {
			_, err = discovery.Register(correlationId, key, connection)
			if err != nil {
				return false, err
			}
			registered = true
		}
	}

	return registered, nil
}

func (c *ConnectionResolver) Register(correlationId string, connection *ConnectionParams) error {
	registered, err := c.registerInDiscovery(correlationId, connection)
	if registered {
		c.connections = append(c.connections, connection)
	}
	return err
}
