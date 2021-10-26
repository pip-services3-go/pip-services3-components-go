package state

// Interface for state storages that are used to store and retrieve transaction states.
type IStateStore interface {

	// Loads state from the store using its key.
	// If value is missing in the store it returns nil.
	//
	// - correlationId     (optional) transaction id to trace execution through call chain.
	// - key               a unique state key.
	// Returns             the state value or nil if value wasn't found.
	Load(correlationId string, key string) interface{}

	// Loads an array of states from the store using their keys.
	//
	// - correlationId     (optional) transaction id to trace execution through call chain.
	// - keys              unique state keys.
	// Returns                 an array with state values and their corresponding keys.
	LoadBulk(correlationId string, keys []string) []*StateValue

	// Saves state into the store.
	//
	// - correlationId     (optional) transaction id to trace execution through call chain.
	// - key               a unique state key.
	// - value             a state value.
	// Returns             The state that was stored in the store.
	Save(correlationId string, key string, value interface{}) interface{}

	// Deletes a state from the store by its key.
	//
	// - correlationId     (optional) transaction id to trace execution through call chain.
	// - key               a unique value key.
	Delete(correlationId string, key string) interface{}
}
