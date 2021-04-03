package info

import (
	"github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/pip-services3-go/pip-services3-components-go/build"
)

/*
Creates information components by their descriptors.
*/

var contextInfoDescriptor = refer.NewDescriptor("pip-services", "context-info", "default", "*", "1.0")
var containerInfoDescriptor = refer.NewDescriptor("pip-services", "container-info", "default", "*", "1.0")

// Create a new instance of the factory.
// Returns *build.Factory

func NewDefaultInfoFactory() *build.Factory {
	factory := build.NewFactory()

	factory.RegisterType(contextInfoDescriptor, NewContextInfo)
	factory.RegisterType(containerInfoDescriptor, NewContextInfo)

	return factory
}
