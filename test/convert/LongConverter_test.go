package convert

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/pip-services/pip-services-commons-go/convert"
)

func TestToLong(t *testing.T) {
    assert.Nil(t, convert.ToNullableLong(nil))
    
    assert.Equal(t, int64(123), convert.ToLong(123))
    assert.Equal(t, int64(123), convert.ToLong(123.456))
    assert.Equal(t, int64(123), convert.ToLong("123"))
    assert.Equal(t, int64(123), convert.ToLong("123.456"))

    assert.Equal(t, int64(123), convert.ToLongWithDefault(nil, 123))
    assert.Equal(t, int64(0), convert.ToLongWithDefault(false, 123))
    assert.Equal(t, int64(123), convert.ToLongWithDefault("ABC", 123))
}
