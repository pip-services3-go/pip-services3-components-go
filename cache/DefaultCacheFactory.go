package cache

import (
	"github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/pip-services3-go/pip-services3-components-go/build"
)

/*
Creates ICache components by their descriptors.
*/
var nullCacheDescriptor = refer.NewDescriptor("pip-services", "cache", "null", "*", "1.0")
var memoryCacheDescriptor = refer.NewDescriptor("pip-services", "cache", "memory", "*", "1.0")

// Create a new instance of the factory.
// Returns *build.Factory
func NewDefaultCacheFactory() *build.Factory {
	factory := build.NewFactory()

	factory.RegisterType(nullCacheDescriptor, NewNullCache)
	factory.RegisterType(memoryCacheDescriptor, NewMemoryCache)

	return factory
}
