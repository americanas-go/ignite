package prometheus

import (
	"context"
	"github.com/americanas-go/config"
	"github.com/americanas-go/grapper"
	prometheus "github.com/americanas-go/grapper/middleware/contrib/prometheus/client_golang.v1"
)

func NewAnyError[R any](ctx context.Context, name string) grapper.AnyErrorMiddleware[R] {
	ConfigAdd(name)
	config.Load()
	if o, _ := NewOptions(name); !o.Enabled {
		return nil
	}
	return prometheus.NewAnyErrorMiddleware[R](ctx)
}

func NewAny[R any](ctx context.Context, name string) grapper.AnyMiddleware[R] {
	ConfigAdd(name)
	config.Load()
	if o, _ := NewOptions(name); !o.Enabled {
		return nil
	}
	return prometheus.NewAnyMiddleware[R](ctx)
}

func NewError(ctx context.Context, name string) grapper.ErrorMiddleware {
	ConfigAdd(name)
	config.Load()
	if o, _ := NewOptions(name); !o.Enabled {
		return nil
	}
	return prometheus.NewErrorMiddleware(ctx)
}
