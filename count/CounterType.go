package count

type CounterType int

const (
	Interval CounterType = iota
	LastValue
	Statistics
	Timestamp
	Increment
)
