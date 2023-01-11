package logrus

import (
	"github.com/americanas-go/ignite/sirupsen/logrus.v1/formatter/cloudwatch"
	"github.com/americanas-go/ignite/sirupsen/logrus.v1/formatter/json"
	"github.com/americanas-go/ignite/sirupsen/logrus.v1/formatter/text"
	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
	lg "github.com/sirupsen/logrus"
)

// NewLogger returns logger with default options.
func NewLogger(hooks ...lg.Hook) log.Logger {
	options := options()
	options.Hooks = hooks

	var formatter lg.Formatter

	switch FormatterType() {
	case "CLOUDWATCH":
		formatter, _ = cloudwatch.NewFormatter()
	case "JSON":
		formatter, _ = json.NewFormatter()
	default:
		formatter, _ = text.NewFormatter()
	}

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
