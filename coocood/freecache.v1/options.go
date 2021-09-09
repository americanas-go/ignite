package freecache

import (
	"github.com/americanas-go/config"
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
	o := &Options{}

	err := config.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}

// NewOptionsWithPath unmarshals a given key path into options and returns it.
func NewOptionsWithPath(path string) (opts *Options, err error) {

	opts, err = NewOptions()
	if err != nil {
		return nil, err
	}

	err = config.UnmarshalWithPath(path, opts)
	if err != nil {
		return nil, err
	}

	return opts, nil
}
