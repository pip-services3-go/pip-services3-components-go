package cache

import (
	"github.com/pip-services-go/pip-services-commons-go/refer"
	"github.com/pip-services-go/pip-services-components-go/build"
)

var NullCacheDescriptor = refer.NewDescriptor("pip-services", "cache", "null", "*", "1.0")
var MemoryCacheDescriptor = refer.NewDescriptor("pip-services", "cache", "memory", "*", "1.0")

func NewDefaultCacheFactory() *build.Factory {
	factory := build.NewFactory()

	factory.RegisterType(NullCacheDescriptor, NewNullCache)
	factory.RegisterType(MemoryCacheDescriptor, NewMemoryCache)

	return factory
}
