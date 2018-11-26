package count

import "time"

type Counter struct {
	Name    string      `json:"name"`
	Type    CounterType `json:"type"`
	Last    float32     `json:"last"`
	Count   int         `json:"count"`
	Min     float32     `json:"min"`
	Max     float32     `json:"max"`
	Average float32     `json:"average"`
	Time    time.Time   `json:"time"`
}

func NewCounter(name string, typ CounterType) *Counter {
	return &Counter{Name: name, Type: typ}
}
