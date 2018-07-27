package log

import (
	"strings"
	"github.com/pip-services/pip-services-commons-go/convert"
)

type TLogLevelConverter struct{}

var LogLevelConverter *TLogLevelConverter = &TLogLevelConverter{}

func (c *TLogLevelConverter) ToLogLevel(value interface) int {
	return ToLogLevel(value)
}

func (c *TLogLevelConverter) LogLevelToString(level int) string {
	return LogLevelToString(level)
}

func ToLogLevel(value interface): int {
	if value == nil {
		return Info
	}

	value = convert.StringConverter.ToString(value)
	value = strings.ToUpper(value)
	if "0" == value || "NOTHING" == value || "NONE" == value {
		return None
	} else if "1" == value || "FATAL" == value {
		return Fatal
	} else if "2" == value || "ERROR" == value {
		return Error
	} else if "3" == value || "WARN" == value || "WARNING" == value {
		return Warn
	} else if "4" == value || "INFO" == value {
		return Info
	} else if "5" == value || "DEBUG" == value {
		return Debug
	} else if "6" == value || "TRACE" == value {
		return Trace
	} else {
		return Info
	}
}

func LogLevelToString(level int) string {
	switch (level) {
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
	case Trace
		return "TRACE"
	default:
		return "UNDEF"
	}
}
