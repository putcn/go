package raw

// #include <stddef.h>
// #cgo CFLAGS: -F /Library/Frameworks -framework Phidget22 -I /Library/Frameworks/Phidget22.framework/Headers
// #include "logging.h"
import "C"

import (
	"fmt"
)

type LogLevel int

const (
	Critical = C.PHIDGET_LOG_CRITICAL
	Error    = C.PHIDGET_LOG_ERROR
	Warning  = C.PHIDGET_LOG_WARNING
	Debug    = C.PHIDGET_LOG_DEBUG
	Info     = C.PHIDGET_LOG_INFO
	Verbose  = C.PHIDGET_LOG_VERBOSE
)

func DisableLogging() error {
	return result(C.PhidgetLog_disable())
}

func EnableLogging(level LogLevel, path string) error {
	return result(C.PhidgetLog_enable(C.Phidget_LogLevel(level), convertString(path)))
}

func Log(level LogLevel, format string, args ...interface{}) error {
	return result(C._log(C.PhidgetLog_enable(level), convertString(fmt.Sprintf(format, args...))))
}
