package datadog

import (
	"github.com/americanas-go/ignite"
	echotrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/labstack/echo.v4"
)

type Options struct {
	Enabled      bool
	TraceOptions []echotrace.Option
}

func NewOptions(traceOptions ...echotrace.Option) (*Options, error) {
	opts := &Options{TraceOptions: traceOptions}

	return ignite.MergeOptionsWithPath[Options](opts, root)
}

func NewOptionsWithPath(path string, traceOptions ...echotrace.Option) (opts *Options, err error) {

	opts, err = NewOptions(traceOptions...)
	if err != nil {
		return nil, err
	}

	return ignite.MergeOptionsWithPath[Options](opts, path)
}
