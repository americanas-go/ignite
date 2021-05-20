package logrus

import (
	"github.com/americanas-go/log"
	logrus "github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
	lg "github.com/sirupsen/logrus"
)

const (
	TextFormatter       = "TEXT"
	JSONFormatter       = "JSON"
	CloudWatchFormatter = "AWS_CLOUD_WATCH"
)

func NewLoggerWithFormatter(formatter lg.Formatter, hooks ...lg.Hook) log.Logger {
	options := options()
	options.Hooks = hooks
	options.Formatter = formatter
	return logrus.NewLoggerWithOptions(options)
}

func NewLogger(hooks ...lg.Hook) log.Logger {
	options := options()
	options.Hooks = hooks
	return logrus.NewLoggerWithOptions(options)
}

func options() *logrus.Options {
	options, err := NewOptions()
	if err != nil {
		panic(err)
	}
	return options
}
