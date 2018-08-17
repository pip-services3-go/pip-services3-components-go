package test_log

import (
	"testing"

	"github.com/pip-services-go/pip-services-commons-go/refer"
	"github.com/pip-services-go/pip-services-components-go/log"
)

func newCompositeLoggerFixture() *LoggerFixture {
	logger := log.NewCompositeLogger()

	refs := refer.NewReferencesFromTuples(
		log.ConsoleLoggerDescriptor, log.NewConsoleLogger(),
		log.CompositeLoggerDescriptor, logger,
	)
	logger.SetReferences(refs)

	fixture := NewLoggerFixture(logger)
	return fixture
}

func TestCompositeLogLevel(t *testing.T) {
	fixture := newCompositeLoggerFixture()
	fixture.TestLogLevel(t)
}

func TestCompositeSimpleLogging(t *testing.T) {
	fixture := newCompositeLoggerFixture()
	fixture.TestSimpleLogging(t)
}

func TestCompositeErrorLogging(t *testing.T) {
	fixture := newCompositeLoggerFixture()
	fixture.TestErrorLogging(t)
}