package logrus

import (
	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
	lg "github.com/sirupsen/logrus"
)

func NewLogger(hooks ...lg.Hook) log.Logger {
	return logrus.NewLogger(options(), hooks...)
}

func NewLoggerWithFormatter(formatter lg.Formatter, hooks ...lg.Hook) log.Logger {
	return logrus.NewLoggerWithFormatter(formatter, options(), hooks...)
}

func options() *logrus.Options {
	options, err := NewOptions()
	if err != nil {
		panic(err)
	}
	return options
}
