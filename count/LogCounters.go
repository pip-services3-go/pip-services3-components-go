package count

import (
	"sort"
	"strconv"

	"github.com/pip-services3-go/pip-services3-commons-go/convert"
	"github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/pip-services3-go/pip-services3-components-go/log"
)

type LogCounters struct {
	CachedCounters
	logger *log.CompositeLogger
}

func NewLogCounters() *LogCounters {
	c := &LogCounters{
		logger: log.NewCompositeLogger(),
	}
	c.CachedCounters = *InheritCacheCounters(c)
	return c
}

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
