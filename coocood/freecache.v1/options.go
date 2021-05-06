package freecache

import (
	"github.com/americanas-go/config"
)

type Options struct {
	CacheSize int
}

type Option func(options *Options)

func WithCacheSize(cacheSize int) Option {
	return func(options *Options) {
		options.CacheSize = cacheSize
	}
}

func NewOptions() (*Options, error) {
	o := &Options{}

	err := config.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}

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
