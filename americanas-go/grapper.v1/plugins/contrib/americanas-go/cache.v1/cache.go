package cache

import (
	"context"
	"github.com/americanas-go/cache"
	"github.com/americanas-go/config"
	"github.com/americanas-go/grapper"
	gcache "github.com/americanas-go/grapper/middleware/contrib/americanas-go/cache.v1"
)

type Cache[T any] struct {
	manager *cache.Manager[T]
}

func New[T any](ctx context.Context, manager *cache.Manager[T]) *Cache[T] {
	return &Cache[T]{manager: manager}
}

func (c *Cache[T]) NewAnyError(ctx context.Context, name string) grapper.AnyErrorMiddleware[T] {
	ConfigAdd(name)
	config.Load()
	if o, _ := NewOptions(name); !o.Enabled {
		return nil
	}
	return gcache.NewAnyErrorMiddleware[T](ctx, c.manager)
}

func (c *Cache[T]) NewAny(ctx context.Context, name string) grapper.AnyMiddleware[T] {
	ConfigAdd(name)
	config.Load()
	if o, _ := NewOptions(name); !o.Enabled {
		return nil
	}
	return gcache.NewAnyMiddleware[T](ctx, c.manager)
}
