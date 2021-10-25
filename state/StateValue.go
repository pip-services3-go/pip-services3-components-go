package state

// A data object that holds a retrieved state value with its key.
type StateValue struct {
	// A unique state key
	Key string `json:"key" bson:"key"`
	// A stored state value;
	Value interface{} `json:"value" bson:"value"`
}
