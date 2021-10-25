package test_state

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-components-go/state"
	"github.com/stretchr/testify/assert"
)

var KEY1 string = "key1"
var KEY2 string = "key2"

var VALUE1 string = "value1"
var VALUE2 string = "value2"

type StateStoreFixture struct {
	state state.IStateStore
}

func NewStateStoreFixture(state state.IStateStore) *StateStoreFixture {
	return &StateStoreFixture{
		state: state,
	}
}

func (c *StateStoreFixture) TestSaveAndLoad(t *testing.T) {
	c.state.Save("", KEY1, VALUE1)
	c.state.Save("", KEY2, VALUE2)

	val := c.state.Load("", KEY1)
	assert.NotNil(t, val)
	assert.Equal(t, VALUE1, val)

	values := c.state.LoadBulk("", []string{KEY2})
	assert.True(t, len(values) == 1)
	assert.Equal(t, KEY2, values[0].Key)
	assert.Equal(t, VALUE2, values[0].Value)
}

func (c *StateStoreFixture) TestDelete(t *testing.T) {
	c.state.Save("", KEY1, VALUE1)

	c.state.Delete("", KEY1)

	val := c.state.Load("", KEY1)
	assert.Nil(t, val)
}
