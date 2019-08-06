package count

import "time"

type NullCounters struct{}

func NewNullCounters() *NullCounters {
	return &NullCounters{}
}

func (c *NullCounters) BeginTiming(name string) *Timing {
	return NewEmptyTiming()
}

func (c *NullCounters) Stats(name string, value float32) {}

func (c *NullCounters) Last(name string, value float32) {}

func (c *NullCounters) TimestampNow(name string) {}

func (c *NullCounters) Timestamp(name string, value time.Time) {}

func (c *NullCounters) IncrementOne(name string) {}

func (c *NullCounters) Increment(name string, value float32) {}
