package logrus

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

// NewOptions returns options from config file or environment vars.
func NewOptions() (*logrus.Options, error) {
	o := &logrus.Options{}

	err := config.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
