package count

import "time"

type ICounters interface {
    BeginTiming(name string) Timing
	Stats(name string, value float32)
	Last(name string, value float32)
	TimestampNow(name string)
	Timestamp(name string, value time.Time)
	IncrementOne(name string)
	Increment(name string, value float32)
}