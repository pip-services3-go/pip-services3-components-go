package log

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pip-services3-go/pip-services3-commons-go/config"
	"github.com/pip-services3-go/pip-services3-commons-go/errors"
	"github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/pip-services3-go/pip-services3-components-go/info"
)

type ILogWriter interface {
	Write(level int, correlationId string, err error, message string)
}

type Logger struct {
	level  int
	source string
	writer ILogWriter
}

func InheritLogger(writer ILogWriter) *Logger {
	return &Logger{
		level:  Info,
		source: "",
		writer: writer,
	}
}

func (c *Logger) Level() int {
	return c.level
}

func (c *Logger) SetLevel(value int) {
	c.level = value
}

func (c *Logger) Source() string {
	return c.source
}

func (c *Logger) SetSource(value string) {
	c.source = value
}

func (c *Logger) Configure(cfg *config.ConfigParams) {
	c.level = LogLevelConverter.ToLogLevel(cfg.GetAsStringWithDefault("level", strconv.Itoa(c.level)))
	c.source = cfg.GetAsStringWithDefault("source", c.source)
}

func (c *Logger) SetReferences(references refer.IReferences) {
	contextInfo, ok := references.GetOneOptional(
		refer.NewDescriptor("pip-services", "context-info", "*", "*", "1.0")).(info.ContextInfo)
	if ok && c.source == "" {
		c.source = contextInfo.Name
	}
}

func (c *Logger) ComposeError(err error) string {
	builder := strings.Builder{}

	appErr, ok := err.(*errors.ApplicationError)
	if ok {
		builder.WriteString(appErr.Message)
		if appErr.Cause != "" {
			builder.WriteString(" Caused by: ")
			builder.WriteString(appErr.Cause)
		}
		if appErr.StackTrace != "" {
			builder.WriteString(" Stack trace: ")
			builder.WriteString(appErr.StackTrace)
		}
	} else {
		builder.WriteString(err.Error())
	}

	return builder.String()
}

func (c *Logger) FormatAndWrite(level int, correlationId string, err error, message string, args []interface{}) {
	if args != nil && len(args) > 0 {
		message = fmt.Sprintf(message, args...)
	}

	if c.writer != nil {
		c.writer.Write(level, correlationId, err, message)
	}
}

func (c *Logger) Log(level int, correlationId string, err error, message string, args ...interface{}) {
	c.FormatAndWrite(level, correlationId, err, message, args)
}

func (c *Logger) Fatal(correlationId string, err error, message string, args ...interface{}) {
	c.FormatAndWrite(Fatal, correlationId, err, message, args)
}

func (c *Logger) Error(correlationId string, err error, message string, args ...interface{}) {
	c.FormatAndWrite(Error, correlationId, err, message, args)
}

func (c *Logger) Warn(correlationId string, message string, args ...interface{}) {
	c.FormatAndWrite(Warn, correlationId, nil, message, args)
}

func (c *Logger) Info(correlationId string, message string, args ...interface{}) {
	c.FormatAndWrite(Info, correlationId, nil, message, args)
}

func (c *Logger) Debug(correlationId string, message string, args ...interface{}) {
	c.FormatAndWrite(Debug, correlationId, nil, message, args)
}

func (c *Logger) Trace(correlationId string, message string, args ...interface{}) {
	c.FormatAndWrite(Trace, correlationId, nil, message, args)
}
