package bigcache

import (
	"context"

	"github.com/allegro/bigcache"
	"github.com/americanas-go/cache"
	cbigcache "github.com/americanas-go/cache/contrib/allegro/bigcache.v1"
)

// NewDriver returns a cache.
func NewDriver(ctx context.Context, cache *bigcache.BigCache) (c cache.Driver, err error) {
	return cbigcache.New(cache), nil
}
