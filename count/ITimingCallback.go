package count

import "time"

type ITimingCallback interface {
	EndTiming(name string, elapsed time.Duration)
}
