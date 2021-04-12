package trace

import (
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	cbuild "github.com/pip-services3-go/pip-services3-components-go/build"
)

//Creates [[ITracer]] components by their descriptors.
//See [[Factory]]
//See [[NullTracer]]
//See [[ConsoleTracer]]
//See [[CompositeTracer]]

type DefaultTracerFactory struct {
	cbuild.Factory
	NullTracerDescriptor      *cref.Descriptor
	LogTracerDescriptor       *cref.Descriptor
	CompositeTracerDescriptor *cref.Descriptor
}

//Create a new instance of the factory.
func NewDefaultTracerFactory() *DefaultTracerFactory {
	c := &DefaultTracerFactory{
		Factory:                   *cbuild.NewFactory(),
		NullTracerDescriptor:      cref.NewDescriptor("pip-services", "tracer", "null", "*", "1.0"),
		LogTracerDescriptor:       cref.NewDescriptor("pip-services", "tracer", "log", "*", "1.0"),
		CompositeTracerDescriptor: cref.NewDescriptor("pip-services", "tracer", "composite", "*", "1.0"),
	}

	c.RegisterType(c.NullTracerDescriptor, NewNullTracer)
	c.RegisterType(c.LogTracerDescriptor, NewLogTracer)
	c.RegisterType(c.CompositeTracerDescriptor, NewCompositeTracer)

	return c
}
