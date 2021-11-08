package datadog

import (
	"github.com/americanas-go/config"
	fibertrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gofiber/fiber.v2"
)

type Options struct {
	Enabled      bool
	TraceOptions []fibertrace.Option
}

func NewOptions(traceOptions ...fibertrace.Option) (*Options, error) {
	o := &Options{
		TraceOptions: traceOptions,
	}

	err := config.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}

func NewOptionsWithPath(path string, traceOptions ...fibertrace.Option) (opts *Options, err error) {
	opts, err = NewOptions(traceOptions...)
	if err != nil {
		return nil, err
	}

	err = config.UnmarshalWithPath(path, opts)
	if err != nil {
		return nil, err
	}

	return opts, nil
}
