package config

import (
	cconfig "github.com/pip-services3-go/pip-services3-commons-go/config"
	crun "github.com/pip-services3-go/pip-services3-commons-go/run"
	mustache "github.com/pip-services3-go/pip-services3-expressions-go/mustache"
)

// Abstract config reader that supports configuration parameterization.
// Configuration parameters
//   parameters: this entire section is used as template parameters
type ConfigReader struct {
	parameters *cconfig.ConfigParams
}

// Creates a new instance of the config reader.
// Returns *ConfigReader
func NewConfigReader() *ConfigReader {
	return &ConfigReader{
		parameters: cconfig.NewEmptyConfigParams(),
	}
}

// Configures component by passing configuration parameters.
// Parameters:
//  - config *config.ConfigParams
//  configuration parameters to be set.
func (c *ConfigReader) Configure(config *cconfig.ConfigParams) {
	parameters := config.GetSection("parameters")
	if parameters.Len() > 0 {
		c.parameters = parameters
	}
}

// Parameterized configuration template given as string with dynamic parameters.
// The method uses Mustache template engine implemented in expressions module
// Parameters:
//   - config string
//   a string with configuration template to be parameterized
//   - parameters *config.ConfigParams
//   dynamic parameters to inject into the template
// Returns string, error
// a parameterized configuration string abd error.
func (c *ConfigReader) Parameterize(config string, parameters *cconfig.ConfigParams) (string, error) {
	if parameters == nil {
		parameters = cconfig.NewEmptyConfigParams()
	}

	parameters = c.parameters.Override(parameters)

	context := parameters.Value()

	mustacheTemplate, err := mustache.NewMustacheTemplateFromString(config)
	if err != nil {
		return "", err
	}

	result, err := mustacheTemplate.EvaluateWithVariables(context)
	return result, err
}

// AddChangeListener - Adds a listener that will be notified when configuration is changed
func (c *ConfigReader) AddChangeListener(listener crun.INotifiable) {
	// Do nothing...
}

// RemoveChangeListener - Remove a previously added change listener.
func (c *ConfigReader) RemoveChangeListener(listener crun.INotifiable) {
	// Do nothing...
}
