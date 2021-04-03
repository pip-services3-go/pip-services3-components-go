package test_cache

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pip-services3-go/pip-services3-components-go/cache"
)

func TestNullCache(t *testing.T) {
	cache := cache.NewNullCache()

	value, err := cache.Retrieve("", "key1")
	assert.Nil(t, value)
	assert.Nil(t, err)

	var str string
	ref, err := cache.RetrieveAs("", "key1", &str)
	assert.Nil(t, ref)
	assert.Equal(t, str, "")
	assert.Nil(t, err)

	value, err = cache.Store("", "key1", "value1", 0)
	assert.Equal(t, "value1", value)
	assert.Nil(t, err)
}
