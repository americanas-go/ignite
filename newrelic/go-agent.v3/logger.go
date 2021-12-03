package newrelic

import (
	"github.com/americanas-go/log"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// Logger represents a newrelic logger.
type Logger struct {
}

// NewLogger returns a new logger instance.
func NewLogger() newrelic.Logger {
	return &Logger{}
}

// Error logs error message.
func (l *Logger) Error(msg string, context map[string]interface{}) {
	log.WithFields(context).Error(msg)
}

// Warn logs warn message.
func (l *Logger) Warn(msg string, context map[string]interface{}) {
	log.WithFields(context).Warn(msg)
}

// Info logs info message.
func (l *Logger) Info(msg string, context map[string]interface{}) {
	log.WithFields(context).Info(msg)
}

// Debug logs debug message.
func (l *Logger) Debug(msg string, context map[string]interface{}) {
	log.WithFields(context).Debug(msg)
}

// DebugEnabled returns debug config value.
func (l *Logger) DebugEnabled() bool {
	return Debug()
}
