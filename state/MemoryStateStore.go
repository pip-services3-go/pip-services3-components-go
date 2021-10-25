package state

import (
	"time"

	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	"github.com/pip-services3-go/pip-services3-commons-go/errors"
)

// State store that keeps states in the process memory.
//
// Remember: This implementation is not suitable for synchronization of distributed processes.
//
// ### Configuration parameters ###
//
// __options:__
// - timeout:               default caching timeout in milliseconds (default: disabled)
//
// See ICache
//
// ### Example ###
//
//     let store = new MemoryStateStore();
//
//     let value = await store.load("123", "key1");
//     ...
//     await store.save("123", "key1", "ABC");
//
type MemoryStateStore struct { //implements IStateStore, IReconfigurable
	states  map[string]interface{}
	timeout int64
}

// Creates a new instance of the state store.

func NewEmptyMemoryStateStore() *MemoryStateStore {
	return &MemoryStateStore{
		states:  make(map[string]interface{}, 0),
		timeout: 0,
	}
}

// Configures component by passing configuration parameters.
//
// - config    configuration parameters to be set.
func (c *MemoryStateStore) Configure(config *cconf.ConfigParams) {
	c.timeout = config.GetAsLongWithDefault("options.timeout", c.timeout)
}

// Clears component state.
//
// - correlationId 	(optional) transaction id to trace execution through call chain.
func (c *MemoryStateStore) cleanup() {
	if c.timeout == 0 {
		return
	}

	cutOffTime := time.Now().UTC().UnixNano() - c.timeout
	// Cleanup obsolete entries
	for prop := range c.states {
		// Remove obsolete entry
		if entry, ok := c.states[prop].(*StateEntry); ok && entry.GetLastUpdateTime() < cutOffTime {
			delete(c.states, prop)
		}
	}
}

// Loads stored value from the store using its key.
// If value is missing in the store it returns nil.
//
// - correlationId     (optional) transaction id to trace execution through call chain.
// - key               a unique state key.
// Returns                 the state value or <code>nil</code> if value wasn't found.
func (c *MemoryStateStore) Load(correlationId string, key string) interface{} {
	if len(key) == 0 {
		panic(errors.NewError("Key cannot be empty"))
	}
	// Cleanup the stored states
	c.cleanup()
	// Get entry from the store
	entry, ok := c.states[key].(*StateEntry)

	// Store has nothing
	if !ok || entry == nil {
		return nil
	}
	return entry.GetValue()
}

// Loads an array of states from the store using their keys.
//
// - correlationId     (optional) transaction id to trace execution through call chain.
// - keys              unique state keys.
// Returns                 an array with state values.
func (c *MemoryStateStore) LoadBulk(correlationId string, keys []string) []*StateValue {
	// Cleanup the stored states
	c.cleanup()

	result := make([]*StateValue, 0)
	for _, key := range keys {
		value := c.Load(correlationId, key)
		result = append(result, &StateValue{Key: key, Value: value})
	}
	return result
}

// Saves state into the store
//
// - correlationId     (optional) transaction id to trace execution through call chain.
// - key               a unique state key.
// - value             a state value to store.
// Returns                 The value that was stored in the cache.
func (c *MemoryStateStore) Save(correlationId string, key string, value interface{}) interface{} {
	if len(key) == 0 {
		panic(errors.NewError("Key cannot be empty"))
	}

	// Cleanup the stored states
	c.cleanup()

	// Get the entry
	entry, ok := c.states[key].(*StateEntry)

	// Shortcut to remove entry from the cache
	if value == nil {
		delete(c.states, key)
		return value
	}

	// Update the entry
	if ok && entry != nil {
		entry.SetValue(value)
	} else { // Or create a new entry
		entry = NewStateEntry(key, value)
		c.states[key] = entry
	}

	return value
}

// Deletes a state from the store by its key.
//
// - correlationId     (optional) transaction id to trace execution through call chain.
// - key               a unique state key.

func (c *MemoryStateStore) Delete(correlationId string, key string) interface{} {
	if len(key) == 0 {
		panic(errors.NewError("Key cannot be empty"))
	}

	// Cleanup the stored states
	c.cleanup()

	// Get the entry
	entry, ok := c.states[key].(*StateEntry)

	// Remove entry from the cache
	if ok && entry != nil {
		delete(c.states, key)
		return entry.GetValue()
	}

	return nil
}
