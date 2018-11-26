package count

import "time"

type Timing struct {
	_start    time.Time
	_callback ITimingCallback
	_counter  string
}

func NewTiming(counter string, callback ITimingCallback) *Timing {
	return &Timing{
		_counter:  counter,
		_callback: callback,
		_start:    time.Now(),
	}
}

func (tm *Timing) EndTiming() {
	if tm._callback != nil {
		tm._callback.EndTiming(tm._counter, time.Since(tm._start))
	}
}
