package trace

import (
	"time"

	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	cerr "github.com/pip-services3-go/pip-services3-commons-go/errors"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	ctx "github.com/pip-services3-go/pip-services3-components-go/info"
)

// Abstract tracer that caches recorded traces in memory and periodically dumps them.
// Child classes implement saving cached traces to their specified destinations.
//
// ### Configuration parameters ###
//
// - source:            source (context) name
// - options:
//     - interval:        interval in milliseconds to save log messages (default: 10 seconds)
//     - maxcache_size:  maximum number of messages stored in this cache (default: 100)
//
// ### References ###
//
// - *:context-info:*:*:1.0     (optional) [[ContextInfo]] to detect the context id and specify counters source
//
// See [[ITracer]]
// See [[OperationTrace]]

type ICachedTraceSaver interface {
	Save(operations []*OperationTrace) error
}

type CachedTracer struct {
	source       string
	Cache        []*OperationTrace
	updated      bool
	lastDumpTime time.Time
	maxCacheSize int
	interval     int64
	saver        ICachedTraceSaver
}

// Creates a new instance of the logger.
func InheritCachedTracer(saver ICachedTraceSaver) *CachedTracer {
	return &CachedTracer{
		Cache:        make([]*OperationTrace, 0),
		updated:      false,
		lastDumpTime: time.Now().UTC(),
		maxCacheSize: 100,
		interval:     10000,
		saver:        saver,
	}
}

// Configures component by passing configuration parameters.
//
// - config    configuration parameters to be set.
func (c *CachedTracer) Configure(config *cconf.ConfigParams) {
	c.interval = config.GetAsLongWithDefault("options.interval", c.interval)
	c.maxCacheSize = config.GetAsIntegerWithDefault("options.maxcache_size", c.maxCacheSize)
	c.source = config.GetAsStringWithDefault("source", c.source)
}

// Sets references to dependent components.
//
// - references 	references to locate the component dependencies.
func (c *CachedTracer) SetReferences(references cref.IReferences) {
	ref := references.GetOneOptional(
		cref.NewDescriptor("pip-services", "context-info", "*", "*", "1.0"))
	if ref != nil {
		contextInfo, _ := ref.(*ctx.ContextInfo)
		if contextInfo != nil && c.source == "" {
			c.source = contextInfo.Name
		}
	}
}

// Writes a log message to the logger destination.
//
// - correlationId     (optional) transaction id to trace execution through call chain.
// - component         a name of called component
// - operation         a name of the executed operation.
// - error             an error object associated with this trace.
// - duration          execution duration in milliseconds.
func (c *CachedTracer) Write(correlationId string, component string, operation string, err error, duration int64) {

	var errorDesc *cerr.ErrorDescription

	if err != nil {
		errorDesc = cerr.NewErrorDescription(err)
	}

	trace := &OperationTrace{
		Time:          time.Now().UTC(),
		Source:        c.source,
		Component:     component,
		Operation:     operation,
		CorrelationId: correlationId,
		Duration:      duration,
		Error:         *errorDesc,
	}
	c.Cache = append(c.Cache, trace)
	c.Update()
}

// Records an operation trace with its name and duration
//
// - correlationId     (optional) transaction id to trace execution through call chain.
// - component         a name of called component
// - operation         a name of the executed operation.
// - duration          execution duration in milliseconds.
func (c *CachedTracer) Trace(correlationId string, component string, operation string, duration int64) {
	c.Write(correlationId, component, operation, nil, duration)
}

// Records an operation failure with its name, duration and error
//
// - correlationId     (optional) transaction id to trace execution through call chain.
// - component         a name of called component
// - operation         a name of the executed operation.
// - error             an error object associated with this trace.
// - duration          execution duration in milliseconds.
func (c *CachedTracer) Failure(correlationId string, component string, operation string, err error, duration int64) {
	c.Write(correlationId, component, operation, err, duration)
}

// Begings recording an operation trace
//
// - correlationId     (optional) transaction id to trace execution through call chain.
// - component         a name of called component
// - operation         a name of the executed operation.
// Returns                 a trace timing object.
func (c *CachedTracer) BeginTrace(correlationId string, component string, operation string) *TraceTiming {
	return NewTraceTiming(correlationId, component, operation, c)
}

// Clears (removes) all cached log messages.
func (c *CachedTracer) Clear() {
	c.Cache = make([]*OperationTrace, 0)
	c.updated = false
}

// Dumps (writes) the currently cached log messages.
// See [[Write]]
func (c *CachedTracer) Dump() {
	if c.updated {
		if !c.updated {
			return
		}

		traces := c.Cache
		c.Cache = make([]*OperationTrace, 0)

		err := c.saver.Save(traces)

		if err != nil {
			// Adds traces back to the cache
			traces = append(traces, c.Cache...)
			c.Cache = traces

			// Truncate cache to max size
			if len(c.Cache) > c.maxCacheSize {
				c.Cache = c.Cache[len(c.Cache)-c.maxCacheSize:]
			}
		}

		c.updated = false
		c.lastDumpTime = time.Now().UTC()
	}
}

// Makes trace cache as updated
// and dumps it when timeout expires.
// See [[Dump]]
func (c *CachedTracer) Update() {
	c.updated = true
	elapsed := int64(time.Since(c.lastDumpTime).Seconds() * 1000)
	if elapsed > c.interval {
		// Todo: Decide what to do with the error
		c.Dump()
	}
}
