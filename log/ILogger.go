package log

type ILogger interface {
	Level() int
	SetLevel(value int)

	Log(level int, correlationId string, err error, message string, args ...interface{})

	Fatal(correlationId string, err error, message string, args ...interface{})
	Error(correlationId string, err error, message string, args ...interface{})
	Warn(correlationId string, message string, args ...interface{})
	Info(correlationId string, message string, args ...interface{})
	Debug(correlationId string, message string, args ...interface{})
	Trace(correlationId string, message string, args ...interface{})
}
