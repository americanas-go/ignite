package text

import (
	"github.com/americanas-go/ignite"
	"github.com/sirupsen/logrus"
)

// NewFormatter returns logrus formatter for text.
func NewFormatter() (logrus.Formatter, error) {
	return ignite.NewOptionsWithPath[logrus.TextFormatter](root)
}
