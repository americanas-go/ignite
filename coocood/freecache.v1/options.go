package freecache

import (
	"github.com/americanas-go/ignite"
)

// Options represents cache options.
type Options struct {
	CacheSize int
}

// Option represents an option.
type Option func(options *Options)

// WithCacheSize returns option that defines cache size.
func WithCacheSize(cacheSize int) Option {
	return func(options *Options) {
		options.CacheSize = cacheSize
	}
}

// NewOptions returns options from config file or environment vars.
func NewOptions() (*Options, error) {
	return ignite.NewOptionsWithPath[Options](root)
}

// NewOptionsWithPath unmarshals a given key path into options and returns it.
// NewOptionsWithPath unmarshals a given key path into options and returns it.
func NewOptionsWithPath(path string) (opts *Options, err error) {
	return ignite.NewOptionsWithPath[Options](root, path)
}
