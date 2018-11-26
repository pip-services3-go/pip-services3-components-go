package count

import "time"

type ICounters interface {
	BeginTiming(name string) *Timing
	Stats(name string, value time.Duration)
	Last(name string, value time.Duration)
	TimestampNow(name string)
	Timestamp(name string, value time.Time)
	IncrementOne(name string)
	Increment(name string, value time.Duration)
}
