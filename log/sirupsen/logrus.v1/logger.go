package logrus

import (
	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
	lg "github.com/sirupsen/logrus"
)

// NewLogger returns logger with default options.
func NewLogger(formatter lg.Formatter, hooks ...lg.Hook) log.Logger {
	options := options()
	options.Hooks = hooks
	options.Formatter = formatter
	return logrus.NewLoggerWithOptions(options)
}

func options() *logrus.Options {
	options, err := NewOptions()
	if err != nil {
		panic(err)
	}
	return options
}
