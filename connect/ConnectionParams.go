package connect

import (
	"github.com/pip-services3-go/pip-services3-commons-go/config"
)

type ConnectionParams struct {
	config.ConfigParams
}

func NewEmptyConnectionParams() *ConnectionParams {
	return &ConnectionParams{
		ConfigParams: *config.NewEmptyConfigParams(),
	}
}

func NewConnectionParams(values map[string]string) *ConnectionParams {
	return &ConnectionParams{
		ConfigParams: *config.NewConfigParams(values),
	}
}

func NewConnectionParamsFromValue(value interface{}) *ConnectionParams {
	return &ConnectionParams{
		ConfigParams: *config.NewConfigParamsFromValue(value),
	}
}

func NewConnectionParamsFromTuples(tuples ...interface{}) *ConnectionParams {
	return &ConnectionParams{
		ConfigParams: *config.NewConfigParamsFromTuplesArray(tuples),
	}
}

func NewConnectionParamsFromTuplesArray(tuples []interface{}) *ConnectionParams {
	return &ConnectionParams{
		ConfigParams: *config.NewConfigParamsFromTuplesArray(tuples),
	}
}

func NewConnectionParamsFromString(line string) *ConnectionParams {
	return &ConnectionParams{
		ConfigParams: *config.NewConfigParamsFromString(line),
	}
}

func NewConnectionParamsFromMaps(maps ...map[string]string) *ConnectionParams {
	return &ConnectionParams{
		ConfigParams: *config.NewConfigParamsFromMaps(maps...),
	}
}

func NewManyConnectionParamsFromConfig(config *config.ConfigParams) []*ConnectionParams {
	result := []*ConnectionParams{}

	connections := config.GetSection("connections")

	if connections.Len() > 0 {
		for _, section := range connections.GetSectionNames() {
			connection := connections.GetSection(section)
			result = append(result, NewConnectionParams(connection.Value()))
		}
	} else {
		connection := config.GetSection("connection")
		if connection.Len() > 0 {
			result = append(result, NewConnectionParams(connection.Value()))
		}
	}

	return result
}

func NewConnectionParamsFromConfig(config *config.ConfigParams) *ConnectionParams {
	connections := NewManyConnectionParamsFromConfig(config)
	if len(connections) > 0 {
		return connections[0]
	}
	return nil
}

func (c *ConnectionParams) UseDiscovery() bool {
	return c.GetAsString("discovery_key") != ""
}

func (c *ConnectionParams) DiscoveryKey() string {
	return c.GetAsString("discovery_key")
}

func (c *ConnectionParams) SetDiscoveryKey(value string) {
	c.Put("discovery_key", value)
}

func (c *ConnectionParams) Protocol() string {
	return c.GetAsString("protocol")
}

func (c *ConnectionParams) ProtocolWithDefault(defaultValue string) string {
	return c.GetAsStringWithDefault("protocol", defaultValue)
}

func (c *ConnectionParams) SetProtocol(value string) {
	c.Put("protocol", value)
}

func (c *ConnectionParams) Host() string {
	host := c.GetAsString("host")
	if host != "" {
		return host
	}
	return c.GetAsString("ip")
}

func (c *ConnectionParams) SetHost(value string) {
	c.Put("host", value)
}

func (c *ConnectionParams) Port() int {
	return c.GetAsInteger("port")
}

func (c *ConnectionParams) PortWithDefault(defaultValue int) int {
	return c.GetAsIntegerWithDefault("port", defaultValue)
}

func (c *ConnectionParams) SetPort(value int) {
	c.Put("port", value)
}

func (c *ConnectionParams) Uri() string {
	return c.GetAsString("uri")
}

func (c *ConnectionParams) SetUri(value string) {
	c.Put("uri", value)
}
