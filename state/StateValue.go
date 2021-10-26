package state

// A data object that holds a retrieved state value with its key.
type StateValue struct {
	Key   string      `json:"key" bson:"key"`     // A unique state key
	Value interface{} `json:"value" bson:"value"` // A stored state value
}
