package state

// Dummy state store implementation that doesn't do anything.
//
// It can be used in testing or in situations when state management is not required
// but shall be disabled.
type NullStateStore struct {
}

func NewEmptyNullStateStore() *NullStateStore {
	return &NullStateStore{}
}

// Loads state from the store using its key.
// If value is missing in the stored it returns nil.
//
// - correlationId     (optional) transaction id to trace execution through call chain.
// - key               a unique state key.
// Returns                 the state value or nil if value wasn't found.
func (c *NullStateStore) Load(correlationId string, key string) interface{} {
	return nil
}

// Loads an array of states from the store using their keys.
//
// - correlationId     (optional) transaction id to trace execution through call chain.
// - keys              unique state keys.
// Returns                 an array with state values and their corresponding keys.
func (c *NullStateStore) LoadBulk(correlationId string, keys []string) []*StateValue {
	return []*StateValue{}
}

// Saves state into the store.
//
// - correlationId     (optional) transaction id to trace execution through call chain.
// - key               a unique state key.
// - value             a state value.
// Returns                 The state that was stored in the store.
func (c *NullStateStore) Save(correlationId string, key string, value interface{}) interface{} {
	return value
}

// Deletes a state from the store by its key.
//
// - correlationId     (optional) transaction id to trace execution through call chain.
// - key               a unique value key.
func (c *NullStateStore) Delete(correlationId string, key string) interface{} {
	return nil
}
