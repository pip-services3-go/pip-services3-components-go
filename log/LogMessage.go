package log

import (
	"time"

	"github.com/pip-services-go/pip-services-commons-go/errors"
)

type LogMessage struct {
	Time          time.Time               `json:"time"`
	Source        string                  `json:"source"`
	Level         int                     `json:"level"`
	CorrelationId string                  `json:"correlation_id"`
	Error         errors.ErrorDescription `json:"error"`
	Message       string                  `json:"message"`
}

func NewLogMessage(level int, source string, correlationId string,
	err errors.ErrorDescription, message string) LogMessage {
	return LogMessage{
		Time:          time.Now().UTC(),
		Source:        source,
		Level:         level,
		CorrelationId: correlationId,
		Error:         err,
		Message:       message,
	}
}
