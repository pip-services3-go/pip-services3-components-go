package log

type ILogger interface{
    Level() int
    SetLevel(value int)
	
	Log(level int, correlationId string, err error, message ...interface{})

    Fatal(correlationId string, err error, message ...interface{})
    Error(correlationId string, err error, message ...interface{})
    Warn(correlationId string, err error, message ...interface{})
    Info(correlationId string, err error, message ...interface{})
    Debug(correlationId string, err error, message ...interface{})
    Trace(correlationId string, err error, message ...interface{})
}
