package log

import (
	"github.com/americanas-go/ignite/go.uber.org/zap.v1"
	"github.com/americanas-go/ignite/rs/zerolog.v1"
	"github.com/americanas-go/ignite/sirupsen/logrus.v1"
	"github.com/americanas-go/log"
)

// New initializes the log according to the configured type and formatter.
func New() {
	switch Type() {
	case "NOOP":
		log.NewNoop()
	case "ZEROLOG":
		zerolog.NewLogger()
	case "ZAP":
		zap.NewLogger()
	default:
		logrus.NewLogger()
	}
}
