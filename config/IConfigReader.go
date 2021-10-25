package config

import (
	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	crun "github.com/pip-services3-go/pip-services3-commons-go/run"
)

/*
Interface for configuration readers that retrieve configuration from various sources and make it available for other components.

Some IConfigReader implementations may support configuration parameterization. The parameterization allows to use configuration as a template and inject there dynamic values. The values may come from application command like arguments or environment variables.
*/
type IConfigReader interface {
	// Reads configuration and parameterize it with given values.
	ReadConfig(correlationId string, parameters *cconf.ConfigParams) (*cconf.ConfigParams, error)

	//Adds a listener that will be notified when configuration is changed
	AddChangeListener(listener crun.INotifiable)

	// Remove a previously added change listener.
	RemoveChangeListener(listener crun.INotifiable)
}
