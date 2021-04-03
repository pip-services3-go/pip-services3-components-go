package log

import (
	"github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/pip-services3-go/pip-services3-components-go/build"
)

/*
Creates ILogger components by their descriptors.
*/
var nullLoggerDescriptor = refer.NewDescriptor("pip-services", "logger", "null", "*", "1.0")
var consoleLoggerDescriptor = refer.NewDescriptor("pip-services", "logger", "console", "*", "1.0")
var compositeLoggerDescriptor = refer.NewDescriptor("pip-services", "logger", "composite", "*", "1.0")

// Create a new instance of the factory.
// Returns *build.Factory
func NewDefaultLoggerFactory() *build.Factory {
	factory := build.NewFactory()

	factory.RegisterType(nullLoggerDescriptor, NewNullLogger)
	factory.RegisterType(consoleLoggerDescriptor, NewConsoleLogger)
	factory.RegisterType(compositeLoggerDescriptor, NewCompositeLogger)

	return factory
}
