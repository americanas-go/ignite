package log

import (
	"github.com/jvitoroc/ignite/go.uber.org/zap.v1"
	"github.com/jvitoroc/ignite/rs/zerolog.v1"
	"github.com/jvitoroc/ignite/sirupsen/logrus.v1"
)

func New() {
	switch Impl() {
	case "ZEROLOG":
		zerolog.NewLogger()
	case "ZAP":
		zap.NewLogger()
	default:
		logrus.NewLogger()
	}
}
