package datadog

import (
	"github.com/americanas-go/config"
	echotrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/labstack/echo.v4"
)

type Options struct {
	Enabled      bool
	TraceOptions []echotrace.Option
}

func NewOptions(traceOptions ...echotrace.Option) (*Options, error) {
	o := &Options{TraceOptions: traceOptions}

	err := config.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}

func NewOptionsWithPath(path string, traceOptions ...echotrace.Option) (opts *Options, err error) {

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
