package datadog

import (
	"github.com/americanas-go/ignite"
	fibertrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gofiber/fiber.v2"
)

// Options represents datadog plugin for fiber options.
type Options struct {
	Enabled      bool
	TraceOptions []fibertrace.Option
}

// NewOptions returns options from config file or environment vars.
func NewOptions(traceOptions ...fibertrace.Option) (*Options, error) {
	opts := &Options{
		TraceOptions: traceOptions,
	}

	return ignite.MergeOptionsWithPath[Options](opts, root)
}

// NewOptionsWithPath unmarshals options based a given key path.
func NewOptionsWithPath(path string, traceOptions ...fibertrace.Option) (opts *Options, err error) {
	opts, err = NewOptions(traceOptions...)
	if err != nil {
		return nil, err
	}

	return ignite.MergeOptionsWithPath[Options](opts, path)
}
