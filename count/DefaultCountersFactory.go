package count

import (
	"github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/pip-services3-go/pip-services3-components-go/build"
)

var NullCountersDescriptor = refer.NewDescriptor("pip-services", "counters", "null", "*", "1.0")
var LogCountersDescriptor = refer.NewDescriptor("pip-services", "counters", "log", "*", "1.0")
var CompositeCountersDescriptor = refer.NewDescriptor("pip-services", "counters", "composite", "*", "1.0")

func NewDefaultCountersFactory() *build.Factory {
	factory := build.NewFactory()

	factory.RegisterType(NullCountersDescriptor, NewNullCounters)
	factory.RegisterType(LogCountersDescriptor, NewLogCounters)
	factory.RegisterType(CompositeCountersDescriptor, NewCompositeCounters)

	return factory
}
