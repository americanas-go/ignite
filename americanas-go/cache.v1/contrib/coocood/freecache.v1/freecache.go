package freecache

import (
	"context"

	"github.com/americanas-go/cache"
	"github.com/coocood/freecache"
)

type fcache struct {
	cache   *freecache.Cache
	options *Options
}

func (c *fcache) Del(ctx context.Context, key string) error {
	c.cache.Del([]byte(key))
	return nil
}

func (c *fcache) Get(ctx context.Context, key string) (data []byte, err error) {
	return c.cache.Get([]byte(key))
}

func (c *fcache) Set(ctx context.Context, key string, data []byte) (err error) {

	seconds := c.options.TTL.Seconds()

	if err = c.cache.Set([]byte(key), data, int(seconds)); err != nil {
		return err
	}

	return nil
}

func New(cache *freecache.Cache, options *Options) cache.Driver {
	return &fcache{cache: cache, options: options}
}
