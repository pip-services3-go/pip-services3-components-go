package log

import (
	"github.com/pip-services-go/pip-services-commons-go/refer"
	"github.com/pip-services-go/pip-services-components-go/build"
)

var NullLoggerDescriptor = refer.NewDescriptor("pip-services", "logger", "null", "*", "1.0")
var ConsoleLoggerDescriptor = refer.NewDescriptor("pip-services", "logger", "console", "*", "1.0")
var CompositeLoggerDescriptor = refer.NewDescriptor("pip-services", "logger", "composite", "*", "1.0")

func NewDefaultLoggerFactory() *build.Factory {
	factory := build.NewFactory()

	factory.RegisterType(NullLoggerDescriptor, NewNullLogger)
	factory.RegisterType(ConsoleLoggerDescriptor, NewConsoleLogger)

	return factory
}
