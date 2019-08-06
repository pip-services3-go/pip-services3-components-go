package test_test

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-components-go/test"
	"github.com/stretchr/testify/assert"
)

func TestShutdown(t *testing.T) {
	sd := test.NewShutdown()

	defer func() {
		err := recover()
		assert.NotNil(t, err)
	}()

	sd.Shutdown()
}
