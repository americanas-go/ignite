package datadog

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/grapper"
	"github.com/americanas-go/grapper/middleware/contrib/datadog/dd-trace-go.v1"
)

func NewAnyError[R any](name string) grapper.AnyErrorMiddleware[R] {
	ConfigAdd(name)
	config.Load()
	if o, _ := NewOptions(name); !o.Enabled {
		return nil
	}
	return datadog.NewAnyErrorMiddleware[R](name, "wrapper")
}

func NewAny[R any](name string) grapper.AnyMiddleware[R] {
	ConfigAdd(name)
	config.Load()
	if o, _ := NewOptions(name); !o.Enabled {
		return nil
	}
	return datadog.NewAnyMiddleware[R](name, "wrapper")
}

func NewError(name string) grapper.ErrorMiddleware {
	ConfigAdd(name)
	config.Load()
	if o, _ := NewOptions(name); !o.Enabled {
		return nil
	}
	return datadog.NewErrorMiddleware(name, "wrapper")
}
