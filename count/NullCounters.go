package count

import "time"

type NullCounters struct {
}

func NewNullCounters() *NullCounters {
	return &NullCounters{}
}

func (nc *NullCounters) BeginTiming(name string) *Timing {
	return NewTiming("", nil)
}

func (nc *NullCounters) Stats(name string, value time.Duration) {
	return
}

func (nc *NullCounters) Last(name string, value time.Duration) {
	return
}

func (nc *NullCounters) TimestampNow(name string) {
	return
}

func (nc *NullCounters) Timestamp(name string, value time.Time) {
	return
}

func (nc *NullCounters) IncrementOne(name string) {
	return
}

func (nc *NullCounters) Increment(name string, value time.Duration) {
	return
}
