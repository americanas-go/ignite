package datadog

import (
	"github.com/americanas-go/ignite"
	chitrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/go-chi/chi.v5"
)

// Options struct which represents a datadog plugin for chi options.
type Options struct {
	Enabled      bool
	TraceOptions []chitrace.Option
}

// NewOptions returns options from config file or environment vars.
func NewOptions(traceOptions ...chitrace.Option) (*Options, error) {
	opts := &Options{TraceOptions: traceOptions}

	return ignite.MergeOptionsWithPath[Options](opts, root)
}

// NewOptionsWithPath returns options from config path.
func NewOptionsWithPath(path string, traceOptions ...chitrace.Option) (opts *Options, err error) {
	opts, err = NewOptions(traceOptions...)
	if err != nil {
		return nil, err
	}

	return ignite.MergeOptionsWithPath[Options](opts, path)
}
