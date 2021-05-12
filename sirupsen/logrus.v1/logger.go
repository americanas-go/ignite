package logrus

import (
	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
	lg "github.com/sirupsen/logrus"
)

func NewLogger() log.Logger {
	return logrus.NewLogger(options())
}

func NewLoggerWithFormatter(formatter lg.Formatter) log.Logger {
	return logrus.NewLoggerWithFormatter(formatter, options())
}

func options() *logrus.Options {
	options, err := NewOptions()
	if err != nil {
		panic(err)
	}
	return options
}
