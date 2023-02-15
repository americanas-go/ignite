package semaphore

import "github.com/americanas-go/ignite"

// Options semaphore plugin for echo server options.
type Options struct {
	Enabled bool
	Limit   int64
}

// NewOptions returns options from config file or environment vars.
func NewOptions() (*Options, error) {
	return ignite.NewOptionsWithPath[Options](root)
}

// NewOptionsWithPath unmarshals a given key path into options and returns it.
func NewOptionsWithPath(path string) (opts *Options, err error) {
	return ignite.NewOptionsWithPath[Options](root, path)
}