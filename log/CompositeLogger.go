package log

import "github.com/pip-services-go/pip-services-commons-go/refer"

type CompositeLogger struct {
	Logger
	loggers []ILogger
}

func NewCompositeLogger() *CompositeLogger {
	c := &CompositeLogger{
		loggers: []ILogger{},
	}
	c.Logger = *InheritLogger(c)
	c.SetLevel(Trace)
	return c
}

func NewCompositeLoggerFromReferences(references refer.IReferences) *CompositeLogger {
	c := NewCompositeLogger()
	c.SetReferences(references)
	return c
}

func (c *CompositeLogger) SetReferences(references refer.IReferences) {
	c.Logger.SetReferences(references)

	if c.loggers == nil {
		c.loggers = []ILogger{}
	}

	loggers := references.GetOptional(refer.NewDescriptor("*", "logger", "*", "*", "*"))
	for _, l := range loggers {
		if l == c {
			continue
		}

		logger, ok := l.(ILogger)
		if ok {
			c.loggers = append(c.loggers, logger)
		}
	}
}

func (c *CompositeLogger) Write(level int, correlationId string, err error, message string) {
	if c.loggers == nil && len(c.loggers) == 0 {
		return
	}

	for _, logger := range c.loggers {
		logger.Log(level, correlationId, err, message)
	}
}
