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

func (con_p ConnectionParams) UseDiscovery() bool {
	return con_p.GetAsNullableString("discovery_key") != nil
}

func (con_p ConnectionParams) GetDiscoveryKey() string {
	return con_p.GetAsString("discovery_key")
}

func (con_p *ConnectionParams) SetDiscoveryKey(value string) {
	con_p.Put("discovery_key", value)
}

func (con_p ConnectionParams) GetProtocol(defaultValue string) string {
	return con_p.GetAsStringWithDefault("protocol", defaultValue)
}

func (con_p *ConnectionParams) SetProtocol(value string) {
	con_p.Put("protocol", value)
}

func (con_p ConnectionParams) GetHost() string {
	host := con_p.GetAsNullableString("host")
	if host != nil {
		return *host
	}
	return *con_p.GetAsNullableString("ip")
}

func (con_p *ConnectionParams) SetHost(value string) {
	con_p.Put("host", value)
}

func (con_p ConnectionParams) GetPort() int {
	return con_p.GetAsInteger("port")
}

func (con_p *ConnectionParams) SetPort(value int) {
	con_p.Put("port", value)
}

func (con_p ConnectionParams) GetUri() string {
	return con_p.GetAsString("uri")
}

func (con_p *ConnectionParams) SetUri(value string) {
	con_p.Put("uri", value)
}








