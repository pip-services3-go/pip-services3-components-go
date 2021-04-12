package lock

import (
	"github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/pip-services3-go/pip-services3-components-go/build"
)

/*
Creates ILock components by their descriptors.
*/
var nullLockDescriptor = refer.NewDescriptor("pip-services", "lock", "null", "*", "1.0")
var memoryLockDescriptor = refer.NewDescriptor("pip-services", "lock", "memory", "*", "1.0")

// Create a new instance of the factory.
// Returns *build.Factory
func NewDefaultLockFactory() *build.Factory {
	factory := build.NewFactory()

	factory.RegisterType(nullLockDescriptor, NewNullLock)
	factory.RegisterType(memoryLockDescriptor, NewMemoryLock)

	return factory
}
