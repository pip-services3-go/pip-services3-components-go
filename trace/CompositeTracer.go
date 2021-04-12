package trace

import (
	"github.com/pip-services3-go/pip-services3-commons-go/refer"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
)

//Aggregates all tracers from component references under a single component.
//It allows to record traces and conveniently send them to multiple destinations.
//### References ###
//- \*:tracer:\*:\*:1.0     (optional) [[ITracer]] components to pass operation traces
//See [[ITracer]]
//### Example ###
//    type MyComponent struct {
//         tracer CompositeTracer
//		}
//       func NewMyComponent() *MyComponent{
//			return &MyComponent{
//				tracer: NewCompositeTracer(nil);
//          }
//       }
//        func (c* MyComponent) SetReferences(references IReferences) {
//            c.tracer.SetReferences(references)
//            ...
//        }

//        public MyMethod(correlatonId string) {
//            timing := c.tracer.BeginTrace(correlationId, "mycomponent", "mymethod");
//
//                ...
//                timing.EndTrace();
//            if err != nil {
//                timing.EndFailure(err);
//            }
//        }
//

type CompositeTracer struct {
	Tracers []ITracer
}

//Creates a new instance of the tracer.
//- references 	references to locate the component dependencies.

func NewCompositeTracer(references cref.IReferences) *CompositeTracer {
	c := &CompositeTracer{}
	if references != nil {
		c.SetReferences(references)
	}
	return c
}

//Sets references to dependent components.
//- references 	references to locate the component dependencies.
func (c *CompositeTracer) SetReferences(references cref.IReferences) {

	if c.Tracers == nil {
		c.Tracers = []ITracer{}
	}

	tracers := references.GetOptional(refer.NewDescriptor("*", "tracer", "*", "*", "*"))
	for _, l := range tracers {
		if l == c {
			continue
		}

		tracer, ok := l.(ITracer)
		if ok {
			c.Tracers = append(c.Tracers, tracer)
		}
	}

}

//Records an operation trace with its name and duration
//- correlationId     (optional) transaction id to trace execution through call chain.
//- component         a name of called component
//- operation         a name of the executed operation.
//- duration          execution duration in milliseconds.
func (c *CompositeTracer) Trace(correlationId string, component string, operation string, duration int64) {
	for _, tracer := range c.Tracers {
		tracer.Trace(correlationId, component, operation, duration)
	}
}

//Records an operation failure with its name, duration and error
//- correlationId     (optional) transaction id to trace execution through call chain.
//- component         a name of called component
//- operation         a name of the executed operation.
//- error             an error object associated with this trace.
//- duration          execution duration in milliseconds.
func (c *CompositeTracer) Failure(correlationId string, component string, operation string, err error, duration int64) {
	for _, tracer := range c.Tracers {
		tracer.Failure(correlationId, component, operation, err, duration)
	}
}

//Begings recording an operation trace
//- correlationId     (optional) transaction id to trace execution through call chain.
//- component         a name of called component
//- operation         a name of the executed operation.
//@returns                 a trace timing object.

func (c *CompositeTracer) BeginTrace(correlationId string, component string, operation string) *TraceTiming {
	return NewTraceTiming(correlationId, component, operation, c)
}
