package trace

//Interface for tracer components that capture operation traces.
type ITracer interface {

	//Records an operation trace with its name and duration
	//
	//- correlationId     (optional) transaction id to trace execution through call chain.
	//- component         a name of called component
	//- operation         a name of the executed operation.
	//- duration          execution duration in milliseconds.
	Trace(correlationId string, component string, operation string, duration int64)

	//Records an operation failure with its name, duration and error
	//- correlationId     (optional) transaction id to trace execution through call chain.
	//- component         a name of called component
	//- operation         a name of the executed operation.
	//- error             an error object associated with this trace.
	//- duration          execution duration in milliseconds.
	Failure(correlationId string, component string, operation string, err error, duration int64)

	//Begings recording an operation trace
	//- correlationId     (optional) transaction id to trace execution through call chain.
	//- component         a name of called component
	//- operation         a name of the executed operation.
	//Returns                 a trace timing object.
	BeginTrace(correlationId string, component string, operation string) *TraceTiming
}
