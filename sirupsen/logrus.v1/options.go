package logrus

import (
	"github.com/americanas-go/ignite"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

// NewOptions returns options from config file or environment vars.
func NewOptions() (*logrus.Options, error) {
	return ignite.NewOptionsWithPath[logrus.Options](root)
}
