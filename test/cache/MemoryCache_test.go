package test_cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/pip-services3-go/pip-services3-components-go/cache"
)

func TestMemoryCache(t *testing.T) {
	cache := cache.NewMemoryCache()

	value, err := cache.Retrieve("", "key1")
	assert.Nil(t, value)
	assert.Nil(t, err)

	value, err = cache.Store("", "key1", "value1", 250)
	assert.Equal(t, "value1", value)
	assert.Nil(t, err)

	value, err = cache.Retrieve("", "key1")
	assert.Equal(t, "value1", value)
	assert.Nil(t, err)

	time.Sleep(500 * time.Millisecond)

	value, err = cache.Retrieve("", "key1")
	assert.Nil(t, value)
	assert.Nil(t, err)
}
