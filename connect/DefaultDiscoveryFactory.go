package connect

import (
	"github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/pip-services3-go/pip-services3-components-go/build"
)

var MemoryDiscoveryDescriptor = refer.NewDescriptor("pip-services", "discovery", "memory", "*", "1.0")

func NewDefaultDiscoveryFactory() *build.Factory {
	factory := build.NewFactory()

	factory.RegisterType(MemoryDiscoveryDescriptor, NewEmptyMemoryDiscovery)

	return factory
}
