package fx

import (
	"github.com/americanas-go/log"
	"go.uber.org/fx"
)

type Logger struct {
	level string
}

func (p *Logger) Printf(format string, args ...interface{}) {
	switch p.level {
	case "INFO":
		log.Infof(format, args...)
	case "TRACE":
		log.Tracef(format, args...)
	default:
		log.Debugf(format, args...)
	}
}

func NewLogger() fx.Printer {
	return &Logger{level: LogLevel()}
}
