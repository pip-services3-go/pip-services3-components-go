package test_tracer

import (
	"errors"
	"testing"

	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	clog "github.com/pip-services3-go/pip-services3-components-go/log"
	"github.com/pip-services3-go/pip-services3-components-go/trace"
	ctrace "github.com/pip-services3-go/pip-services3-components-go/trace"
)

func newLogTracer() *ctrace.LogTracer {
	tracer := trace.NewLogTracer()
	tracer.SetReferences(cref.NewReferencesFromTuples(cref.NewDescriptor("pip-services", "logger", "null", "default", "1.0"), clog.NewNullLogger()))
	return tracer
}

func TestSimpleTracing(t *testing.T) {
	tracer := newLogTracer()
	tracer.Trace("123", "mycomponent", "mymethod", 123456)
	tracer.Failure("123", "mycomponent", "mymethod", errors.New("Test error"), 123456)
}

func TestTraceTiming(t *testing.T) {
	tracer := newLogTracer()
	var timing = tracer.BeginTrace("123", "mycomponent", "mymethod")
	timing.EndTrace()

	timing = tracer.BeginTrace("123", "mycomponent", "mymethod")
	timing.EndFailure(errors.New("Test error"))
}
