package test_log

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-components-go/v3/log"
	"github.com/stretchr/testify/assert"
)

func TestLogLevelConverter(t *testing.T) {
	level := log.LogLevelConverter.ToLogLevel("info")
	assert.Equal(t, log.Info, level)

	level = log.LogLevelConverter.ToLogLevel("4")
	assert.Equal(t, log.Info, level)

	str := log.LogLevelConverter.ToString(level)
	assert.Equal(t, "INFO", str)
}
