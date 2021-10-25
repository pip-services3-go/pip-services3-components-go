package state

import (
	"github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/pip-services3-go/pip-services3-components-go/build"
)

// Creates IStateStore components by their descriptors.
// See Factory
// See IStateStore
// See MemoryStateStore
// See NullStateStore
func NewDefaultStateStoreFactory() *build.Factory {
	factory := build.NewFactory()
	nullStateStoreDescriptor := refer.NewDescriptor("pip-services", "state-store", "null", "*", "1.0")
	memoryStateStoreDescriptor := refer.NewDescriptor("pip-services", "state-store", "memory", "*", "1.0")

	factory.RegisterType(nullStateStoreDescriptor, NewEmptyNullStateStore)
	factory.RegisterType(memoryStateStoreDescriptor, NewEmptyMemoryStateStore)

	return factory
}
