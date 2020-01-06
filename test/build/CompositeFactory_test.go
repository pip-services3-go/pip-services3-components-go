package test_build

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-commons-go/v3/refer"
	"github.com/pip-services3-go/pip-services3-components-go/v3/build"
	"github.com/stretchr/testify/assert"
)

func TestCompositeFactory(t *testing.T) {
	factory := build.NewCompositeFactory()

	subFactory := build.NewFactory()
	descriptor := refer.NewDescriptor("test", "object", "default", "*", "1.0")
	subFactory.Register(descriptor, newObject)

	factory.Add(subFactory)

	locator := factory.CanCreate(descriptor)
	assert.NotNil(t, locator)
	locator = factory.CanCreate("123")
	assert.Nil(t, locator)

	obj, err := factory.Create(descriptor)
	assert.Nil(t, err)
	assert.Equal(t, "ABC", obj)
	obj, err = factory.Create("123")
	assert.NotNil(t, err)
	assert.Nil(t, obj)

	factory.Remove(subFactory)
	locator = factory.CanCreate(descriptor)
	assert.Nil(t, locator)
}
