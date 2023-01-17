package json

import (
	"github.com/americanas-go/ignite"
	"github.com/sirupsen/logrus"
)

// NewFormatter returns logrus formatter for json.
func NewFormatter() (logrus.Formatter, error) {
	return ignite.NewOptionsWithPath[logrus.JSONFormatter](root)
}
