package connect

import (
	"github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/pip-services3-go/pip-services3-components-go/build"
)

type DefaultDiscoveryFactory struct {
	descriptor                refer.Descriptor
	memoryDiscoveryDescriptor refer.Descriptor
	build.Factory
}

func NewDefaultDiscoveryFactory() (defaultFactory *DefaultDiscoveryFactory) {
	defaultFactory = &DefaultDiscoveryFactory{
		Factory:                   *build.NewFactory(),
		descriptor:                *refer.NewDescriptor("pip-services", "factory", "discovery", "default", "1.0"),
		memoryDiscoveryDescriptor: *refer.NewDescriptor("pip-services", "discovery", "memory", "*", "1.0"),
	}
	defaultFactory.RegisterType(defaultFactory.GetMemoryDiscoveryDescriptor(), MemoryDiscovery{})
	return
}

func (ddf *DefaultDiscoveryFactory) GetDescriptor() refer.Descriptor {
	return ddf.descriptor
}

func (ddf *DefaultDiscoveryFactory) GetMemoryDiscoveryDescriptor() refer.Descriptor {
	return ddf.memoryDiscoveryDescriptor
}
