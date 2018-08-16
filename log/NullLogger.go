package log

type NullLogger struct{}

func NewNullLogger() *NullLogger {
	c := &NullLogger{}
	return c
}

func (c *NullLogger) Level() int {
	return None
}

func (c *NullLogger) SetLevel(value int) {
}

func (c *NullLogger) Log(level int, correlationId string, err error, message string, args ...interface{}) {
}

func (c *NullLogger) Fatal(correlationId string, err error, message string, args ...interface{}) {
}

func (c *NullLogger) Error(correlationId string, err error, message string, args ...interface{}) {
}

func (c *NullLogger) Warn(correlationId string, message string, args ...interface{}) {
}

func (c *NullLogger) Info(correlationId string, message string, args ...interface{}) {
}

func (c *NullLogger) Debug(correlationId string, message string, args ...interface{}) {
}

func (c *NullLogger) Trace(correlationId string, message string, args ...interface{}) {
}
