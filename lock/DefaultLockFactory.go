package lock

import (
	"github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/pip-services3-go/pip-services3-components-go/build"
)

var NullLockDescriptor = refer.NewDescriptor("pip-services", "lock", "null", "*", "1.0")
var MemoryLockDescriptor = refer.NewDescriptor("pip-services", "lock", "memory", "*", "1.0")

func NewDefaultLockFactory() *build.Factory {
	factory := build.NewFactory()

	factory.RegisterType(NullLockDescriptor, NewNullLock)
	factory.RegisterType(MemoryLockDescriptor, NewMemoryLock)

	return factory
}
