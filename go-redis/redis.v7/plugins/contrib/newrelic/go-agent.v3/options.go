package newrelic

import (
	"github.com/americanas-go/ignite"
)

// Options represents the newrelic plugin for redis options.
type Options struct {
	Enabled bool
}

// NewOptions returns options from config file or environment vars.
func NewOptions() (*Options, error) {
	return ignite.NewOptionsWithPath[Options](root)
}

// NewOptionsWithPath unmarshals a given key path into options and returns it.
func NewOptionsWithPath(path string) (opts *Options, err error) {
	return ignite.NewOptionsWithPath[Options](root, path)
}

return ignite.NewOptionsWithPath[Options](root, path)
}
