package config

import (
	cconfig "github.com/pip-services3-go/pip-services3-commons-go/config"
	crun "github.com/pip-services3-go/pip-services3-commons-go/run"
	mustache "github.com/pip-services3-go/pip-services3-expressions-go/mustache"
)

/*
Config reader that stores configuration in memory.

The reader supports parameterization using Mustache template engine implemented in expressions module.

Configuration parameters
The configuration parameters are the configuration template

see
IConfigReader

Example
  config := NewConfigParamsFromTuples(
      "connection.host", "{{SERVICE_HOST}}",
      "connection.port", "{{SERVICE_PORT}}{{^SERVICE_PORT}}8080{{/SERVICE_PORT}}"
  );

  configReader := NewMemoryConfigReader();
  configReader.Configure(config);

  parameters := NewConfigParamsFromValue(process.env);

  res, err := configReader.ReadConfig("123", parameters);
  // Possible result: connection.host=10.1.1.100;connection.port=8080
*/
type MemoryConfigReader struct {
	config *cconfig.ConfigParams
}

// Creates a new instance of config reader.
// Returns *MemoryConfigReader
func NewEmptyMemoryConfigReader() *MemoryConfigReader {
	return &MemoryConfigReader{
		config: cconfig.NewEmptyConfigParams(),
	}
}

// Creates a new instance of config reader.
// Parameters:
//   - config *cconfig.ConfigParams
//   component configuration parameters
// Returns *MemoryConfigReader
func NewMemoryConfigReader(config *cconfig.ConfigParams) *MemoryConfigReader {
	return &MemoryConfigReader{
		config: config,
	}
}

// Configures component by passing configuration parameters.
// Parameters:
//   - config *cconfig.ConfigParams
//   configuration parameters to be set.
func (c *MemoryConfigReader) Configure(config *cconfig.ConfigParams) {
	c.config = config
}

// Reads configuration and parameterize it with given values.
// Parameters:
//   - correlationId string
//   transaction id to trace execution through call chain.
//   - parameters *cconfig.ConfigParams
//   values to parameters the configuration or null to skip parameterization.
// Returns *cconfig.ConfigParams, error
// configuration or error.
func (c *MemoryConfigReader) ReadConfig(correlationId string,
	parameters *cconfig.ConfigParams) (*cconfig.ConfigParams, error) {

	if parameters != nil {
		template := c.config.String()
		context := parameters.Value()

		mustacheTemplate, err := mustache.NewMustacheTemplateFromString(template)
		if err != nil {
			return nil, err
		}

		config, err := mustacheTemplate.EvaluateWithVariables(context)
		if err != nil {
			return nil, err
		}

		result := cconfig.NewConfigParamsFromString(config)
		return result, nil
	} else {
		result := cconfig.NewConfigParamsFromValue(c.config.Value())
		return result, nil
	}
}

// AddChangeListener - Adds a listener that will be notified when configuration is changed
func (c *MemoryConfigReader) AddChangeListener(listener crun.INotifiable) {
	// Do nothing...
}

// RemoveChangeListener - Remove a previously added change listener.
func (c *MemoryConfigReader) RemoveChangeListener(listener crun.INotifiable) {
	// Do nothing...
}
