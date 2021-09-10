package datadog

import (
	"github.com/americanas-go/config"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
)

type Options struct {
	Enabled bool
	Options []sqltrace.Option
}

func NewOptions(options ...sqltrace.Option) (*Options, error) {
	o := &Options{
		Options: options,
	}

	err := config.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}

func NewOptionsWithPath(path string, options ...sqltrace.Option) (opts *Options, err error) {
	opts, err = NewOptions(options...)
	if err != nil {
		return nil, err
	}

	err = config.UnmarshalWithPath(path, opts)
	if err != nil {
		return nil, err
	}

	return opts, nil
}
