package datadog

import (
	"github.com/americanas-go/config"
	mongotrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/go.mongodb.org/mongo-driver/mongo"
)

type Options struct {
	Enabled bool
	Options []mongotrace.Option
}

func NewOptions(options ...mongotrace.Option) (*Options, error) {
	o := &Options{
		Options: options,
	}

	err := config.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}

func NewOptionsWithPath(path string, options ...mongotrace.Option) (opts *Options, err error) {
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
