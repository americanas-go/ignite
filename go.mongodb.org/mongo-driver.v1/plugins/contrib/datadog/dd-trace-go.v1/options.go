package datadog

import (
	"github.com/americanas-go/ignite"
)

// Options represents datadog plugin for mongo options.
type Options struct {
	Enabled       bool
	ServiceName   string
	Analytics     bool
	AnalyticsRate float64
}

// NewOptions returns options from config file or environment vars.
func NewOptions() (*Options, error) {
	return ignite.NewOptionsWithPath[Options](root)
}

// NewOptionsWithPath unmarshals options based a given key path.
func NewOptionsWithPath(path string) (opts *Options, err error) {

	opts, err = NewOptions()
	if err != nil {
		return nil, err
	}

	return ignite.MergeOptionsWithPath[Options](opts, path)
}
