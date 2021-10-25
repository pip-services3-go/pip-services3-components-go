package state

import "time"

//Data object to store state values with their keys used by [[MemoryStateEntry]]
type StateEntry struct {
	key            string
	value          interface{}
	lastUpdateTime int64 // timestamp in ms
}

// NewStateEntry method creates a new instance of the state entry and assigns its values.
// - key       a unique key to locate the value.
// - value     a value to be stored.
func NewStateEntry(key string, value interface{}) *StateEntry {
	return &StateEntry{
		key:            key,
		value:          value,
		lastUpdateTime: time.Now().UTC().UnixNano() / (int64)(1000),
	}
}

// GetKey method gets the key to locate the state value.
// Returns the value key.
func (c *StateEntry) GetKey() string {
	return c.key
}

// GetValue method gets the sstate value.
// Returns the value object.
func (c *StateEntry) GetValue() interface{} {
	return c.value
}

// GetLastUpdateTime method gets the last update time.
// Returns the timestamp when the value ware stored.
func (c *StateEntry) GetLastUpdateTime() int64 {
	return c.lastUpdateTime
}

// SetValue method sets a new state value.
// - value     a new cached value.
func (c *StateEntry) SetValue(value interface{}) {
	c.value = value
	c.lastUpdateTime = time.Now().UTC().UnixNano() / (int64)(1000)
}
