package trace

//Dummy implementation of tracer that doesn't do anything.
//It can be used in testing or in situations when tracing is required
//but shall be disabled.
//See ITracer

type NullTracer struct {
}

//Creates a new instance of the tracer.
func NewNullTracer() *NullTracer {
	return &NullTracer{}
}

//Records an operation trace with its name and duration
//- correlationId     (optional) transaction id to trace execution through call chain.
//- component         a name of called component
//- operation         a name of the executed operation.
//- duration          execution duration in milliseconds.
func (c *NullTracer) Trace(correlationId string, component string, operation string, duration int64) {
	// Do nothing...
}

//Records an operation failure with its name, duration and error

//- correlationId     (optional) transaction id to trace execution through call chain.
//- component         a name of called component
//- operation         a name of the executed operation.
//- error             an error object associated with this trace.
//- duration          execution duration in milliseconds.

func (c *NullTracer) Failure(correlationId string, component string, operation string, err error, duration int64) {
	// Do nothing...
}

//Begings recording an operation trace
//- correlationId     (optional) transaction id to trace execution through call chain.
//- component         a name of called component
//- operation         a name of the executed operation.
//Returns                 a trace timing object.
func (c *NullTracer) BeginTrace(correlationId string, component string, operation string) *TraceTiming {
	return NewTraceTiming(correlationId, component, operation, c)
}
