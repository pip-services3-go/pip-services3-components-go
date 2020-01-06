package test_count

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-components-go/v3/count"
)

func TestLogCountersSimpleCounters(t *testing.T) {
	counters := count.NewLogCounters()
	fixture := NewCountersFixture(&counters.CachedCounters)
	fixture.TestSimpleCounters(t)
}

func TestLogCountersMeasureElapsedTime(t *testing.T) {
	counters := count.NewLogCounters()
	fixture := NewCountersFixture(&counters.CachedCounters)
	fixture.TestMeasureElapsedTime(t)
}
