package test

import (
	"github.com/pip-services3-go/pip-services3-commons-go/v3/refer"
	"github.com/pip-services3-go/pip-services3-components-go/v3/build"
)

var ShutdownDescriptor = refer.NewDescriptor("pip-services", "shutdown", "*", "*", "1.0")

func NewDefaultTestFactory() *build.Factory {
	factory := build.NewFactory()

	factory.RegisterType(ShutdownDescriptor, NewShutdown)

	return factory
}
