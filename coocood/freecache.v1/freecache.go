package freecache

import (
	"context"

	"github.com/americanas-go/log"
	"github.com/coocood/freecache"
)

func NewCacheWithOptions(ctx context.Context, o *Options, opts ...Option) (cache *freecache.Cache, err error) {

	logger := log.FromContext(ctx)

	for _, opt := range opts {
		opt(o)
	}

	cache = freecache.NewCache(o.CacheSize)

	logger.Infof("Created cache with size %v", o.CacheSize)

	return cache, err

}

func NewCache(ctx context.Context, opts ...Option) (*freecache.Cache, error) {

	logger := log.FromContext(ctx)

	o, err := NewOptions()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	return NewCacheWithOptions(ctx, o, opts...)
}
