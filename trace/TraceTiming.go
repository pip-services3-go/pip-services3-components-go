package trace

import "time"

//Timing object returned by {ITracer.BeginTrace} to end timing
//of execution block and record the associated trace.
//
//### Example ###
//
//    var timing = tracer.BeginTrace("mymethod.exec_time");
//
//        ...
//        timing.EndTrace();
//    if err != nil {
//        timing.EndFailure(err);
//    }
//

type TraceTiming struct {
	start         int64
	trace         ITracer
	correlationId string
	component     string
	operation     string
}

//Creates a new instance of the timing callback object.
//
//- correlationId     (optional) transaction id to trace execution through call chain.
//- component 	an associated component name
//- operation 	an associated operation name
//- callback 		a callback that shall be called when endTiming is called.
func NewTraceTiming(correlationId string, component string, operation string, tracer ITracer) *TraceTiming {
	return &TraceTiming{
		correlationId: correlationId,
		component:     component,
		operation:     operation,
		trace:         tracer,
		start:         time.Now().UTC().UnixNano(),
	}
}

//Ends timing of an execution block, calculates elapsed time
//and records the associated trace.
func (c *TraceTiming) EndTrace() {
	if c.trace != nil {
		elapsed := time.Now().UTC().UnixNano() - c.start
		c.trace.Trace(c.correlationId, c.component, c.operation, elapsed/int64(time.Millisecond))
	}
}

//Ends timing of a failed block, calculates elapsed time
//and records the associated trace.
//- error             an error object associated with this trace.
func (c *TraceTiming) EndFailure(err error) {
	if c.trace != nil {
		elapsed := time.Now().UTC().UnixNano() - c.start
		c.trace.Failure(c.correlationId, c.component, c.operation, err, elapsed/int64(time.Millisecond))
	}
}
