package datadog

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/log"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
)

type Logger struct {
}

func NewLogger() ddtrace.Logger {
	return &Logger{}
}

func (l *Logger) Log(msg string) {

	var fn func(args ...interface{})

	switch config.String(logLevel) {
	case "INFO":
		fn = log.Info
	case "DEBUG":
		fn = log.Debug
	default:
		fn = log.Debug
	}

	fn(msg)
}
