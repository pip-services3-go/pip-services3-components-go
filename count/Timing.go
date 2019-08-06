package count

import (
	"time"
)

type Timing struct {
	start    time.Time
	callback ITimingCallback
	counter  string
}

func NewEmptyTiming() *Timing {
	return &Timing{
		start: time.Now(),
	}
}

func NewTiming(counter string, callback ITimingCallback) *Timing {
	return &Timing{
		start:    time.Now(),
		callback: callback,
		counter:  counter,
	}
}

func (c *Timing) EndTiming() {
	if c.callback == nil {
		return
	}

	elapsed := time.Since(c.start).Seconds() * 1000
	c.callback.EndTiming(c.counter, float32(elapsed))
}
