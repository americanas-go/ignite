package bigcache

import (
	"context"

	"github.com/allegro/bigcache"
	"github.com/americanas-go/cache"
	cbigcache "github.com/americanas-go/cache/contrib/allegro/bigcache.v1"
	ibigcache "github.com/americanas-go/ignite/allegro/bigcache.v1"
)

// NewCacheWithConfigPath returns a cache with options from config path .
func NewCacheWithConfigPath(ctx context.Context, path string) (cache.Driver, error) {
	options, err := ibigcache.NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewCacheWithOptions(ctx, options)
}

// NewCacheWithOptions returns a cache with options.
func NewCacheWithOptions(ctx context.Context, opt *ibigcache.Options) (c cache.Driver, err error) {
	var bc *bigcache.BigCache
	bc, err = ibigcache.NewCacheWithOptions(ctx, opt)
	if err != nil {
		return nil, err
	}
	return cbigcache.New(bc), nil
}

// NewCache returns a cache.
func NewCache(ctx context.Context) (c cache.Driver, err error) {
	var options *ibigcache.Options
	options, err = ibigcache.NewOptionsWithPath(root)
	if err != nil {
		return nil, err
	}

	var bc *bigcache.BigCache
	bc, err = ibigcache.NewCacheWithOptions(ctx, options)
	if err != nil {
		return nil, err
	}
	return cbigcache.New(bc), nil
}
