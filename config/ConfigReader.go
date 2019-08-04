package config

import (
	cconfig "github.com/pip-services3-go/pip-services3-commons-go/config"
	"github.com/aymerick/raymond"
)

type ConfigReader struct {
	parameters *cconfig.ConfigParams
}

func NewConfigReader() *ConfigReader {
	return &ConfigReader {
		parameters: cconfig.NewEmptyConfigParams(),
	}
}

func (c *ConfigReader) Configure(config *cconfig.ConfigParams) {
	parameters := config.GetSection("parameters")
	if parameters.Len() > 0 {
		c.parameters = parameters
	}
}

func (c *ConfigReader) Parameterize(config string, parameters *cconfig.ConfigParams) (string, error) {
	parameters = c.parameters.Override(parameters)

	context := parameters.Value()
	result, err := raymond.Render(config, context)
	return result, err
}