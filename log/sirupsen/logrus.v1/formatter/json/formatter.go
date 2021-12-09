package json

import (
	"github.com/americanas-go/config"
	"github.com/sirupsen/logrus"
)

// NewFormatter returns logrus formatter for json.
func NewFormatter() (logrus.Formatter, error) {

	fmt := &logrus.JSONFormatter{}

	err := config.UnmarshalWithPath(root, fmt)
	if err != nil {
		return nil, err
	}

	return fmt, nil
}
