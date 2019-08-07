package test

import (
	"github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/pip-services3-go/pip-services3-components-go/build"
)

var ShutdownDescriptor = refer.NewDescriptor("pip-services", "shutdown", "*", "*", "1.0")

func NewDefaultTestFactory() *build.Factory {
	factory := build.NewFactory()

	factory.RegisterType(ShutdownDescriptor, NewShutdown)

	return factory
}
