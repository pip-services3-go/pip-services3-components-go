package count

import (
	"sort"
	"strconv"

	"github.com/pip-services3-go/pip-services3-commons-go/v3/convert"
	"github.com/pip-services3-go/pip-services3-commons-go/v3/refer"
	"github.com/pip-services3-go/pip-services3-components-go/v3/log"
)

/*

Performance counters that periodically dumps counters measurements to logger.

Configuration parameters
options:
interval: interval in milliseconds to save current counters measurements (default: 5 mins)
reset_timeout: timeout in milliseconds to reset the counters. 0 disables the reset (default: 0)
References
*:logger:*:*:1.0 ILogger components to dump the captured counters
*:context-info:*:*:1.0 (optional) ContextInfo to detect the context id and specify counters source
see
Counter

see
CachedCounters

see
CompositeLogger

Example
counters := NewLogCounters();
counters.SetReferences(NewReferencesFromTuples(
    NewDescriptor("pip-services", "logger", "console", "default", "1.0"), NewConsoleLogger()
));

counters.Increment("mycomponent.mymethod.calls");
timing := counters.BeginTiming("mycomponent.mymethod.exec_time");
defer  timing.EndTiming();

// do something

counters.Dump();
*/
type LogCounters struct {
	CachedCounters
	logger *log.CompositeLogger
}

// Creates a new instance of the counters.
// Returns *LogCounters
func NewLogCounters() *LogCounters {
	c := &LogCounters{
		logger: log.NewCompositeLogger(),
	}
	c.CachedCounters = *InheritCacheCounters(c)
	return c
}

// Sets references to dependent components.
// Parameters:
// 			- references refer.IReferences
// 			references to locate the component dependencies.
func (c *LogCounters) SetReferences(references refer.IReferences) {
	c.logger.SetReferences(references)
}

func (c *LogCounters) counterToString(counter *Counter) string {
	result := "Counter " + counter.Name + " { "
	result = result + "\"type\": " + strconv.Itoa(counter.Type)
	switch counter.Type {
	case Interval:
	case Statistics:
		result = result + ", \"count\": " + convert.StringConverter.ToString(counter.Count)
		result = result + ", \"min\": " + convert.StringConverter.ToString(counter.Min)
		result = result + ", \"max\": " + convert.StringConverter.ToString(counter.Max)
		result = result + ", \"avg\": " + convert.StringConverter.ToString(counter.Average)
	case LastValue:
		result = result + ", \"last\": " + convert.StringConverter.ToString(counter.Last)
	case Timestamp:
		result = result + ", \"time\": " + convert.StringConverter.ToString(counter.Time)
	case Increment:
		result = result + ", \"count\": " + convert.StringConverter.ToString(counter.Count)
	}
	result = result + " }"
	return result
}

// Saves the current counters measurements.
// Parameters:
// 			- counters []*Counter
// 			current counters measurements to be saves.
func (c *LogCounters) Save(counters []*Counter) error {
	if counters == nil || len(counters) == 0 {
		return nil
	}

	sort.Slice(counters, func(i, j int) bool {
		return counters[i].Name < counters[2].Name
	})

	for _, counter := range counters {
		c.logger.Info("", c.counterToString(counter))
	}

	return nil
}
