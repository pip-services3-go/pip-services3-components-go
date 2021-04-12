package trace

import (
	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	cconv "github.com/pip-services3-go/pip-services3-commons-go/convert"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	log "github.com/pip-services3-go/pip-services3-components-go/log"
)

//Tracer that dumps recorded traces to logger.

//### Configuration parameters ###
//- options:
//    - log_level:         log level to record traces (default: debug)
//
//### References ###
//- \*:logger:\*:\*:1.0           [[ILogger]] components to dump the captured counters
//- \*:context-info:\*:\*:1.0     (optional) [[ContextInfo]] to detect the context id and specify counters source
//See [[Tracer]]
//See [[CachedCounters]]
//See [[CompositeLogger]]
//### Example ###
//     tracer = NewLogTracer();
//    tracer.SetReferences(NewReferencesFromTuples(
//        NewDescriptor("pip-services", "logger", "console", "default", "1.0"), NewConsoleLogger()
//    ));
//     timing := trcer.BeginTrace("123", "mycomponent", "mymethod");
//
//        ...
//        timing.EndTrace();
//    if err != nil {
//        timing.EndFailure(err);
//    }

type LogTracer struct {
	logger   *log.CompositeLogger
	logLevel int
}

//Creates a new instance of the tracer.
func NewLogTracer() *LogTracer {
	return &LogTracer{
		logger:   log.NewCompositeLogger(),
		logLevel: log.Debug,
	}
}

//Configures component by passing configuration parameters.
//- config    configuration parameters to be set.
func (c *LogTracer) Configure(config *cconf.ConfigParams) {
	logLvl := config.GetAsObject("options.log_level")
	if logLvl == nil {
		logLvl = c.logLevel
	}
	c.logLevel = log.LogLevelConverter.ToLogLevel(logLvl)
}

//Sets references to dependent components.
//- references 	references to locate the component dependencies.
func (c *LogTracer) SetReferences(references cref.IReferences) {
	c.logger.SetReferences(references)
}

func (c *LogTracer) logTrace(correlationId string, component string, operation string, err error, duration int64) {
	builder := ""

	if err != nil {
		builder += "Failed to execute "
	} else {
		builder += "Executed "
	}

	builder += component
	builder += "."
	builder += operation

	if duration > 0 {
		builder += " in " + cconv.StringConverter.ToString(duration) + " msec"
	}

	if err != nil {
		c.logger.Error(correlationId, err, builder)
	} else {
		c.logger.Log(c.logLevel, correlationId, nil, builder)
	}
}

//Records an operation trace with its name and duration
//- correlationId     (optional) transaction id to trace execution through call chain.
//- component         a name of called component
//- operation         a name of the executed operation.
//- duration          execution duration in milliseconds.
func (c *LogTracer) Trace(correlationId string, component string, operation string, duration int64) {
	c.logTrace(correlationId, component, operation, nil, duration)
}

//Records an operation failure with its name, duration and error
//- correlationId     (optional) transaction id to trace execution through call chain.
//- component         a name of called component
//- operation         a name of the executed operation.
//- error             an error object associated with this trace.
//- duration          execution duration in milliseconds.
func (c *LogTracer) Failure(correlationId string, component string, operation string, err error, duration int64) {
	c.logTrace(correlationId, component, operation, err, duration)
}

//Begings recording an operation trace
//- correlationId     (optional) transaction id to trace execution through call chain.
//- component         a name of called component
//- operation         a name of the executed operation.
//Returns                 a trace timing object.
func (c *LogTracer) BeginTrace(correlationId string, component string, operation string) *TraceTiming {
	return NewTraceTiming(correlationId, component, operation, c)
}
