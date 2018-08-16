package log

import (
	"strings"

	"github.com/pip-services-go/pip-services-commons-go/convert"
)

type TLogLevelConverter struct{}

var LogLevelConverter *TLogLevelConverter = &TLogLevelConverter{}

func (c *TLogLevelConverter) ToLogLevel(value interface{}) int {
	return LogLevelFromString(value)
}

func (c *TLogLevelConverter) ToString(level int) string {
	return LogLevelToString(level)
}

func LogLevelFromString(value interface{}) int {
	if value == nil {
		return Info
	}

	str := convert.StringConverter.ToString(value)
	str = strings.ToUpper(str)
	if "0" == str || "NOTHING" == str || "NONE" == str {
		return None
	} else if "1" == str || "FATAL" == str {
		return Fatal
	} else if "2" == str || "ERROR" == str {
		return Error
	} else if "3" == str || "WARN" == str || "WARNING" == str {
		return Warn
	} else if "4" == str || "INFO" == str {
		return Info
	} else if "5" == str || "DEBUG" == str {
		return Debug
	} else if "6" == str || "TRACE" == str {
		return Trace
	} else {
		return Info
	}
}

func LogLevelToString(level int) string {
	switch level {
	case Fatal:
		return "FATAL"
	case Error:
		return "ERROR"
	case Warn:
		return "WARN"
	case Info:
		return "INFO"
	case Debug:
		return "DEBUG"
	case Trace:
		return "TRACE"
	default:
		return "UNDEF"
	}
}
