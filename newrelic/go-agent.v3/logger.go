package newrelic

import (
	"github.com/americanas-go/log"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type Logger struct {
}

func NewLogger() newrelic.Logger {
	return &Logger{}
}

func (l *Logger) Error(msg string, context map[string]interface{}) {
	log.WithFields(context).Error(msg)
}

func (l *Logger) Warn(msg string, context map[string]interface{}) {
	log.WithFields(context).Warn(msg)
}

func (l *Logger) Info(msg string, context map[string]interface{}) {
	log.WithFields(context).Info(msg)
}

func (l *Logger) Debug(msg string, context map[string]interface{}) {
	log.WithFields(context).Debug(msg)
}

func (l *Logger) DebugEnabled() bool {
	return Debug()
}
