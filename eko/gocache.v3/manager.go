package gocache

import (
	"context"

	"github.com/americanas-go/log"
	"github.com/coocood/freecache"
	"github.com/eko/gocache/v3/cache"
)

// NewCacheWithOptions returns a cache with options from config path .
func NewCacheWithConfigPath(ctx context.Context, path string, opts ...Option) (*freecache.Manager, error) {
	options, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewManagerWithOptions(ctx, options, opts...)
}

// NewManagerWithOptions returns a cache with options.
func NewManagerWithOptions(ctx context.Context, o *Options, opts ...Option) (cache *freecache.Manager, err error) {

	cache :=

	logger := log.FromContext(ctx)

	for _, opt := range opts {
		opt(o)
	}

	cache = freecache.NewManager(o.ManagerSize)

	logger.Infof("Created cache with size %v", o.ManagerSize)

	return cache, err
}

// NewManager returns a cache.
func NewManager(ctx context.Context, opts ...Option) (*freecache.Manager, error) {

	o, err := NewOptions()
	if err != nil {
		return nil, err
	}

	return NewManagerWithOptions(ctx, o, opts...)
}
