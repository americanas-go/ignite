package logrus

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/sirupsen/logrus.v1/formatter/cloudwatch"
	"github.com/americanas-go/ignite/sirupsen/logrus.v1/formatter/json"
	"github.com/americanas-go/ignite/sirupsen/logrus.v1/formatter/text"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	lg "github.com/sirupsen/logrus"
)

func NewOptions() (*logrus.Options, error) {
	o := &logrus.Options{}
	err := config.UnmarshalWithPath(root, o)

	if err != nil && !isFormatterErr(err) {
		return nil, err
	}

	f, err := getFormatter()
	if err != nil {
		return nil, err
	}

	o.Formatter = f

	return o, nil
}

func isFormatterErr(err error) bool {
	const formatterErrMsg = "'Formatter' expected type 'logrus.Formatter', got 'string'"
	if e, ok := err.(*mapstructure.Error); ok {
		return len(e.Errors) == 1 && e.Errors[0] == formatterErrMsg
	}
	return false
}

func getFormatter() (lg.Formatter, error) {
	f := Formatter()
	switch f {
	case JSONFormatter:
		return json.NewFormatter()
	case CloudWatchFormatter:
		return cloudwatch.NewFormatter()
	case TextFormatter:
		return text.NewFormatter()
	default:
		return nil, errors.Errorf("unsupported formatter \"%s\"", f)

	}

}
