package count

import (
	"time"
)

/*
Callback object returned by ICounters.beginTiming to end timing of execution block and update the associated counter.

Example
  timing := counters.BeginTiming("mymethod.exec_time");
  defer  timing.EndTiming();

*/
type CounterTiming struct {
	start    time.Time
	callback ICounterTimingCallback
	counter  string
}

// Creates a new instance of the timing callback object.
// Retruns *CounterTiming
func NewEmptyCounterTiming() *CounterTiming {
	return &CounterTiming{
		start: time.Now(),
	}
}

// Creates a new instance of the timing callback object.
// Parameters:
//   - counter string
//   an associated counter name
//   - callback ICounterTimingCallback
//   a callback that shall be called when EndTiming is called.
// Retruns *CounterTiming
func NewCounterTiming(counter string, callback ICounterTimingCallback) *CounterTiming {
	return &CounterTiming{
		start:    time.Now(),
		callback: callback,
		counter:  counter,
	}
}

// Ends timing of an execution block, calculates elapsed time and updates the associated counter.
func (c *CounterTiming) EndTiming() {
	if c.callback == nil {
		return
	}

	elapsed := time.Since(c.start).Seconds() * 1000
	c.callback.EndTiming(c.counter, float32(elapsed))
}
