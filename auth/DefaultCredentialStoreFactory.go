package auth

import (
	"github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/pip-services3-go/pip-services3-components-go/build"
)

var MemoryCredentialStoreDescriptor = refer.NewDescriptor("pip-services", "credential-store", "memory", "*", "1.0")

func NewCredentialStoreFactory() *build.Factory {
	factory := build.NewFactory()

	factory.RegisterType(MemoryCredentialStoreDescriptor, NewEmptyMemoryCredentialStore)

	return factory
}
