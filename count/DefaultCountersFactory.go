package count

import (
	"github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/pip-services3-go/pip-services3-components-go/build"
)

/*
Creates ICounters components by their descriptors.
*/
var nullCountersDescriptor = refer.NewDescriptor("pip-services", "counters", "null", "*", "1.0")
var logCountersDescriptor = refer.NewDescriptor("pip-services", "counters", "log", "*", "1.0")
var compositeCountersDescriptor = refer.NewDescriptor("pip-services", "counters", "composite", "*", "1.0")

// Create a new instance of the factory.
// Returns *build.Factory
func NewDefaultCountersFactory() *build.Factory {
	factory := build.NewFactory()

	factory.RegisterType(nullCountersDescriptor, NewNullCounters)
	factory.RegisterType(logCountersDescriptor, NewLogCounters)
	factory.RegisterType(compositeCountersDescriptor, NewCompositeCounters)

	return factory
}
