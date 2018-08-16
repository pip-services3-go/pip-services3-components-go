package log

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/pip-services-go/pip-services-commons-go/convert"
)

type ConsoleLogger struct {
	Logger
}

func NewConsoleLogger() *ConsoleLogger {
	c := &ConsoleLogger{}
	c.Logger = *InheritLogger(c)
	return c
}

func (c *ConsoleLogger) Write(level int, correlationId string, err error, message string) {
	if c.Level() < level {
		return
	}

	if correlationId == "" {
		correlationId = "---"
	}
	levelStr := LogLevelConverter.ToString(level)
	dateStr := convert.StringConverter.ToString(time.Now().UTC())

	build := strings.Builder{}
	build.WriteString("[")
	build.WriteString(correlationId)
	build.WriteString(":")
	build.WriteString(levelStr)
	build.WriteString(":")
	build.WriteString(dateStr)
	build.WriteString("] ")

	build.WriteString(message)

	if err != nil {
		if len(message) == 0 {
			build.WriteString("Error: ")
		} else {
			build.WriteString(": ")
		}

		build.WriteString(c.ComposeError(err))
	}

	build.WriteString("\n")
	output := build.String()

	if level == Fatal || level == Error || level == Warn {
		fmt.Fprintf(os.Stderr, output)
	} else {
		fmt.Printf(output)
	}
}
