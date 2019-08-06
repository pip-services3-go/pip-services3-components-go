package count

import (
	"time"

	"github.com/pip-services3-go/pip-services3-commons-go/refer"
)

type CompositeCounters struct {
	counters []ICounters
}

func NewCompositeCounters() *CompositeCounters {
	c := &CompositeCounters{
		counters: []ICounters{},
	}
	return c
}

func NewCompositeCountersFromReferences(references refer.IReferences) *CompositeCounters {
	c := NewCompositeCounters()
	c.SetReferences(references)
	return c
}

func (c *CompositeCounters) SetReferences(references refer.IReferences) {
	if c.counters == nil {
		c.counters = []ICounters{}
	}

	counters := references.GetOptional(refer.NewDescriptor("*", "counters", "*", "*", "*"))
	for _, l := range counters {
		if l == c {
			continue
		}

		counter, ok := l.(ICounters)
		if ok {
			c.counters = append(c.counters, counter)
		}
	}
}

func (c *CompositeCounters) BeginTiming(name string) *Timing {
	return NewTiming(name, c)
}

func (c *CompositeCounters) EndTiming(name string, elapsed float32) {
	for _, counter := range c.counters {
		callback, ok := counter.(ITimingCallback)
		if ok {
			callback.EndTiming(name, elapsed)
		}
	}
}

func (c *CompositeCounters) Stats(name string, value float32) {
	for _, counter := range c.counters {
		counter.Stats(name, value)
	}
}

func (c *CompositeCounters) Last(name string, value float32) {
	for _, counter := range c.counters {
		counter.Last(name, value)
	}
}

func (c *CompositeCounters) TimestampNow(name string) {
	c.Timestamp(name, time.Now())
}

func (c *CompositeCounters) Timestamp(name string, value time.Time) {
	for _, counter := range c.counters {
		counter.Timestamp(name, value)
	}
}

func (c *CompositeCounters) IncrementOne(name string) {
	c.Increment(name, 1)
}

func (c *CompositeCounters) Increment(name string, value int) {
	for _, counter := range c.counters {
		counter.Increment(name, value)
	}
}
