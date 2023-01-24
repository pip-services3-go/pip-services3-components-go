package test_connect

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-commons-go/config"
	"github.com/pip-services3-go/pip-services3-components-go/connect"
	"github.com/stretchr/testify/assert"
)

func TestConnectionUtilsConcatOptions(t *testing.T) {
	var options1 = config.NewConfigParamsFromTuples(
		"host", "server1",
		"port", "8080",
		"param1", "ABC",
	)

	var options2 = config.NewConfigParamsFromTuples(
		"host", "server2",
		"port", "8080",
		"param2", "XYZ",
	)

	var options = connect.ConnectionUtils.Concat(options1, options2)

	assert.Equal(t, 4, options.Len())
	assert.Equal(t, "server1,server2", *options.GetAsNullableString("host"))
	assert.Equal(t, "8080,8080", *options.GetAsNullableString("port"))
	assert.Equal(t, "ABC", *options.GetAsNullableString("param1"))
	assert.Equal(t, "XYZ", *options.GetAsNullableString("param2"))
}

func TestConnectionUtilsIncludeKeys(t *testing.T) {
	var options1 = config.NewConfigParamsFromTuples(
		"host", "server1",
		"port", "8080",
		"param1", "ABC",
	)

	var options = connect.ConnectionUtils.Include(options1, "host", "port")

	assert.Equal(t, 2, options.Len())
	assert.Equal(t, "server1", *options.GetAsNullableString("host"))
	assert.Equal(t, "8080", *options.GetAsNullableString("port"))
	assert.Nil(t, options.GetAsNullableString("param1"))
}

func TestConnectionUtilsExcludeKeys(t *testing.T) {
	var options1 = config.NewConfigParamsFromTuples(
		"host", "server1",
		"port", "8080",
		"param1", "ABC",
	)

	var options = connect.ConnectionUtils.Exclude(options1, "host", "port")

	assert.Equal(t, 1, options.Len())
	assert.Nil(t, options.GetAsNullableString("host"))
	assert.Nil(t, options.GetAsNullableString("port"))
	assert.Equal(t, "ABC", *options.GetAsNullableString("param1"))
}

func TestConnectionUtilsParseURI(t *testing.T) {
	var options = connect.ConnectionUtils.ParseUri("broker1", "kafka", 9092)
	assert.Equal(t, 4, options.Len())
	assert.Equal(t, "broker1:9092", *options.GetAsNullableString("servers"))
	assert.Equal(t, "kafka", *options.GetAsNullableString("protocol"))
	assert.Equal(t, "broker1", *options.GetAsNullableString("host"))
	assert.Equal(t, "9092", *options.GetAsNullableString("port"))

	options = connect.ConnectionUtils.ParseUri("tcp://broker1:8082", "kafka", 9092)
	assert.Equal(t, 4, options.Len())
	assert.Equal(t, "broker1:8082", *options.GetAsNullableString("servers"))
	assert.Equal(t, "tcp", *options.GetAsNullableString("protocol"))
	assert.Equal(t, "broker1", *options.GetAsNullableString("host"))
	assert.Equal(t, "8082", *options.GetAsNullableString("port"))

	options = connect.ConnectionUtils.ParseUri("tcp://user:pass123@broker1:8082", "kafka", 9092)
	assert.Equal(t, 6, options.Len())
	assert.Equal(t, "broker1:8082", *options.GetAsNullableString("servers"))
	assert.Equal(t, "tcp", *options.GetAsNullableString("protocol"))
	assert.Equal(t, "broker1", *options.GetAsNullableString("host"))
	assert.Equal(t, "8082", *options.GetAsNullableString("port"))
	assert.Equal(t, "user", *options.GetAsNullableString("username"))
	assert.Equal(t, "pass123", *options.GetAsNullableString("password"))

	options = connect.ConnectionUtils.ParseUri("tcp://user:pass123@broker1,broker2:8082", "kafka", 9092)
	assert.Equal(t, 6, options.Len())
	assert.Equal(t, "broker1:9092,broker2:8082", *options.GetAsNullableString("servers"))
	assert.Equal(t, "tcp", *options.GetAsNullableString("protocol"))
	assert.Equal(t, "broker1,broker2", *options.GetAsNullableString("host"))
	assert.Equal(t, "9092,8082", *options.GetAsNullableString("port"))
	assert.Equal(t, "user", *options.GetAsNullableString("username"))
	assert.Equal(t, "pass123", *options.GetAsNullableString("password"))

	options = connect.ConnectionUtils.ParseUri("tcp://user:pass123@broker1:8082,broker2:8082?param1=ABC&param2=XYZ", "kafka", 9092)
	assert.Equal(t, 8, options.Len())
	assert.Equal(t, "broker1:8082,broker2:8082", *options.GetAsNullableString("servers"))
	assert.Equal(t, "tcp", *options.GetAsNullableString("protocol"))
	assert.Equal(t, "broker1,broker2", *options.GetAsNullableString("host"))
	assert.Equal(t, "8082,8082", *options.GetAsNullableString("port"))
	assert.Equal(t, "user", *options.GetAsNullableString("username"))
	assert.Equal(t, "pass123", *options.GetAsNullableString("password"))
	assert.Equal(t, "ABC", *options.GetAsNullableString("param1"))
	assert.Equal(t, "XYZ", *options.GetAsNullableString("param2"))
}

func TestConnectionUtilsParseURI2(t *testing.T) {
	var options = config.NewConfigParamsFromTuples(
		"host", "broker1,broker2",
		"port", ",8082",
		"username", "user",
		"password", "pass123",
		"param1", "ABC",
		"param2", "XYZ",
		"param3", nil,
	)

	var uri = connect.ConnectionUtils.ComposeUri(options, "tcp", 9092)
	assert.Equal(t, len("tcp://user:pass123@broker1:9092,broker2:8082?param1=ABC&param2=XYZ&param3"), len(uri))

	uri = connect.ConnectionUtils.ComposeUri(options, "", 0)
	assert.Equal(t, len("user:pass123@broker1,broker2:8082?param1=ABC&param2=XYZ&param3"), len(uri))
}
