package config

import (
	cconfig "github.com/pip-services3-go/pip-services3-commons-go/config"
	"github.com/aymerick/raymond"
)

type MemoryConfigReader struct {
	config *cconfig.ConfigParams
}

func NewEmptyMemoryConfigReader() *MemoryConfigReader {
	return &MemoryConfigReader {
		config: cconfig.NewEmptyConfigParams(),
	}
}

func NewMemoryConfigReader(config *cconfig.ConfigParams) *MemoryConfigReader {
	return &MemoryConfigReader {
		config: config,
	}
}

func (c *MemoryConfigReader) Configure(config *cconfig.ConfigParams) {
	c.config = config
}

func (c *MemoryConfigReader) ReadConfig(correlationId string,
	parameters *cconfig.ConfigParams) (*cconfig.ConfigParams, error) {

	if parameters != nil {
		template := c.config.String()
		context := parameters.Value()
		config, err := raymond.Render(template, context)
		result := cconfig.NewConfigParamsFromString(config)
		return result, err
	} else {
		result := cconfig.NewConfigParamsFromValue(c.config.Value())
		return result, nil
	}
}
