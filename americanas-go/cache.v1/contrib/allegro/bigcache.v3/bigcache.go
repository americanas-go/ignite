package bigcache

import (
	"context"

	"github.com/allegro/bigcache/v3"
	"github.com/americanas-go/cache"
	cbigcache "github.com/americanas-go/cache/driver/contrib/allegro/bigcache.v3"
)

// NewDriver returns a cache.
func NewDriver(ctx context.Context, cache *bigcache.BigCache) (c cache.Driver, err error) {
	return cbigcache.New(cache), nil
}
