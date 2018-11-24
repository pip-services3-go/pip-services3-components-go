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

func NewManyConnectionParamsFromConfig(conf *config.ConfigParams) []*ConnectionParams {
	result := make([]*ConnectionParams, 0)
	connections := conf.GetSection("connection")
	if connections.Len() > 0 {
		connectionSections := connections.GetSectionNames()
		for _, conName := range connectionSections {
			//TODO::Need to discuss with Sergey!
			result = append(result, NewConnectionParamsFromValue(connections.GetSection(conName)))
		}
	} else {
		conn := conf.GetSection("connection")
		//TODO::Need to discuss with Sergey!
		result = append(result, NewConnectionParamsFromValue(conn))
	}
	return result
}

func NewConnectionParamsFromConfig(conf *config.ConfigParams) *ConnectionParams {
	connections := NewManyConnectionParamsFromConfig(conf)
	if len(connections) > 0 {
		return connections[0]
	}

	return nil
}

func (conP ConnectionParams) UseDiscovery() bool {
	return conP.GetAsNullableString("discovery_key") != nil
}

func (conP ConnectionParams) GetDiscoveryKey() string {
	return conP.GetAsString("discovery_key")
}

func (conP *ConnectionParams) SetDiscoveryKey(value string) {
	conP.Put("discovery_key", value)
}

func (conP ConnectionParams) GetProtocol(defaultValue string) string {
	return conP.GetAsStringWithDefault("protocol", defaultValue)
}

func (conP *ConnectionParams) SetProtocol(value string) {
	conP.Put("protocol", value)
}

func (conP ConnectionParams) GetHost() string {
	host := conP.GetAsNullableString("host")
	if host != nil {
		return *host
	}
	return *conP.GetAsNullableString("ip")
}

func (conP *ConnectionParams) SetHost(value string) {
	conP.Put("host", value)
}

func (conP ConnectionParams) GetPort() int {
	return conP.GetAsInteger("port")
}

func (conP *ConnectionParams) SetPort(value int) {
	conP.Put("port", value)
}

func (conP ConnectionParams) GetUri() string {
	return conP.GetAsString("uri")
}

func (conP *ConnectionParams) SetUri(value string) {
	conP.Put("uri", value)
}
