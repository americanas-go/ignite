package cloudwatch

import (
	"github.com/americanas-go/config"
	"github.com/ravernkoh/cwlogsfmt"
	"github.com/sirupsen/logrus"
)

func NewFormatter() (logrus.Formatter, error) {

	fmt := &cwlogsfmt.CloudWatchLogsFormatter{}

	err := config.UnmarshalWithPath(root, fmt)
	if err != nil {
		return nil, err
	}

	return fmt, nil
}
