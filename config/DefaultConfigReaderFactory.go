package config

import (
	"github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/pip-services3-go/pip-services3-components-go/build"
)

/*
Creates IConfigReader components by their descriptors.
*/
var memoryConfigReaderDescriptor = refer.NewDescriptor("pip-services", "config-reader", "memory", "*", "1.0")
var jsonConfigReaderDescriptor = refer.NewDescriptor("pip-services", "config-reader", "json", "*", "1.0")
var yamlConfigReaderDescriptor = refer.NewDescriptor("pip-services", "config-reader", "yaml", "*", "1.0")

//Create a new instance of the factory.
//Returns *build.Factory
func NewDefaultConfigReaderFactory() *build.Factory {
	factory := build.NewFactory()

	factory.RegisterType(memoryConfigReaderDescriptor, NewEmptyMemoryConfigReader)
	factory.RegisterType(jsonConfigReaderDescriptor, NewEmptyJsonConfigReader)
	factory.RegisterType(yamlConfigReaderDescriptor, NewEmptyYamlConfigReader)

	return factory
}
