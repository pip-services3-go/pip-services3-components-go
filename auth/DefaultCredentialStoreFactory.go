package auth

import (
	"github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/pip-services3-go/pip-services3-components-go/build"
)

/*
Creates ICredentialStore components by their descriptors.
*/
var MemoryCredentialStoreDescriptor = refer.NewDescriptor("pip-services", "credential-store", "memory", "*", "1.0")

// Create a new instance of the factory.
// Returns *build.Factory
func NewDefaultCredentialStoreFactory() *build.Factory {
	factory := build.NewFactory()

	factory.RegisterType(MemoryCredentialStoreDescriptor, NewEmptyMemoryCredentialStore)

	return factory
}
