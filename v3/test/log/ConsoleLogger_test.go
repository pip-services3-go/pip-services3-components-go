package test_log

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-components-go/v3/log"
)

func newConsoleLoggerFixture() *LoggerFixture {
	logger := log.NewConsoleLogger()
	fixture := NewLoggerFixture(logger)
	return fixture
}

func TestConsoleLogLevel(t *testing.T) {
	fixture := newConsoleLoggerFixture()
	fixture.TestLogLevel(t)
}

func TestConsoleSimpleLogging(t *testing.T) {
	fixture := newConsoleLoggerFixture()
	fixture.TestSimpleLogging(t)
}

func TestConsoleErrorLogging(t *testing.T) {
	fixture := newConsoleLoggerFixture()
	fixture.TestErrorLogging(t)
}
