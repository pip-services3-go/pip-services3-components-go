package test_tracer

import (
	"errors"
	"testing"

	ctrace "github.com/pip-services3-go/pip-services3-components-go/trace"
)

func newNullTracer() *ctrace.NullTracer {
	return ctrace.NewNullTracer()
}

func TestSimpleNullTracing(t *testing.T) {
	tracer := newNullTracer()
	tracer.Trace("123", "mycomponent", "mymethod", 123456)
	tracer.Failure("123", "mycomponent", "mymethod", errors.New("Test error"), 123456)
}

func TestTraceNullTiming(t *testing.T) {
	tracer := newNullTracer()
	timing := tracer.BeginTrace("123", "mycomponent", "mymethod")
	timing.EndTrace()

	timing = tracer.BeginTrace("123", "mycomponent", "mymethod")
	timing.EndFailure(errors.New("Test error"))
}
