package text

import (
	"github.com/americanas-go/config"
	"github.com/sirupsen/logrus"
)

func NewFormatter() (logrus.Formatter, error) {

	fmt := &logrus.TextFormatter{}

	err := config.UnmarshalWithPath(root, fmt)
	if err != nil {
		return nil, err
	}

	return fmt, nil
}
