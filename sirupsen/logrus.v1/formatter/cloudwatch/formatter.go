package cloudwatch

import (
	"github.com/americanas-go/ignite"
	"github.com/ravernkoh/cwlogsfmt"
	"github.com/sirupsen/logrus"
)

// NewFormatter returns logrus formatter for cloudwatch.
func NewFormatter() (logrus.Formatter, error) {
	return ignite.NewOptionsWithPath[cwlogsfmt.CloudWatchLogsFormatter](root)
}
