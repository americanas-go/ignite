package bigcache

import (
	"context"

	"github.com/allegro/bigcache/v3"
	"github.com/americanas-go/log"
)

// NewCacheWithConfig returns a cache with options from config path .
func NewCacheWithConfigPath(ctx context.Context, path string) (*bigcache.BigCache, error) {
	options, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewCacheWithOptions(ctx, options)
}

// NewCacheWithOptions returns a cache with options.
func NewCacheWithOptions(ctx context.Context, opt *Options) (cache *bigcache.BigCache, err error) {

	logger := log.FromContext(ctx)

	cfg := bigcache.Config{
		Shards:             opt.Shards,
		LifeWindow:         opt.LifeWindow,
		CleanWindow:        opt.CleanWindow,
		MaxEntriesInWindow: opt.MaxEntriesInWindow,
		MaxEntrySize:       opt.MaxEntrySize,
		HardMaxCacheSize:   opt.HardMaxCacheSize,
		Verbose:            opt.Verbose,
		Logger:             log.GetLogger(),
		// Hasher:             nil,
		// OnRemove:           nil,
		// OnRemoveWithReason: nil,
	}

	cache, err = bigcache.NewBigCache(cfg)
	if err != nil {
		return nil, err
	}

	logger.Infof("Created bigcache")

	return cache, nil
}

// NewCache returns a cache.
func NewCache(ctx context.Context) (*bigcache.BigCache, error) {
	o, err := NewOptions()
	if err != nil {
		return nil, err
	}
	return NewCacheWithOptions(ctx, o)
}
