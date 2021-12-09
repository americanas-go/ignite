package log

import (
	"github.com/americanas-go/ignite/log/go.uber.org/zap.v1"
	"github.com/americanas-go/ignite/log/rs/zerolog.v1"
	"github.com/americanas-go/ignite/log/sirupsen/logrus.v1"
	"github.com/americanas-go/ignite/log/sirupsen/logrus.v1/formatter/cloudwatch"
	"github.com/americanas-go/ignite/log/sirupsen/logrus.v1/formatter/json"
	"github.com/americanas-go/ignite/log/sirupsen/logrus.v1/formatter/text"
	lr "github.com/sirupsen/logrus"
)

// New initializes the log according to the configured type and formatter.
func New() {
	switch Type() {
	case "ZEROLOG":
		zerolog.NewLogger()
	case "ZAP":
		zap.NewLogger()
	default:

		var formatter lr.Formatter
		var err error

		switch LogrusFormatter() {
		case "CLOUDWATCH":
			formatter, err = cloudwatch.NewFormatter()
		case "JSON":
			formatter, err = json.NewFormatter()
		default:
			formatter, err = text.NewFormatter()
		}

		if err != nil {
			panic(err)
		}

		logrus.NewLogger(formatter)
	}
}
