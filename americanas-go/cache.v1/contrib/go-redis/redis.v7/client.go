package freecache

import (
	"context"

	"github.com/americanas-go/cache"
	"github.com/go-redis/redis/v7"
)

type client struct {
	cache   *redis.Client
	options *Options
}

func (c *client) Del(ctx context.Context, key string) error {
	return c.cache.WithContext(ctx).Del(key).Err()
}

func (c *client) Get(ctx context.Context, key string) (data []byte, err error) {
	return c.cache.WithContext(ctx).Get(key).Bytes()
}

func (c *client) Set(ctx context.Context, key string, data []byte) (err error) {
	c.cache.WithContext(ctx).Del(key)

	if err = c.cache.WithContext(ctx).Set(key, data, c.options.TTL).Err(); err != nil {
		return err
	}

	return nil
}

func NewClient(cache *redis.Client, options *Options) cache.Driver {
	return &client{cache: cache, options: options}
}
