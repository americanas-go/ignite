package hystrix

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/grapper"
	"github.com/americanas-go/grapper/middleware/contrib/afex/hystrix-go.v0"
	h "github.com/americanas-go/ignite/afex/hystrix-go.v0"
	"github.com/americanas-go/log"
)

func NewAnyError[R any](name string) grapper.AnyErrorMiddleware[R] {
	ConfigAdd(name)
	config.Load()
	if o, _ := NewOptions(name); !o.Enabled {
		return nil
	}
	if err := h.ConfigureCommand(name); err != nil {
		log.Error(err.Error())
	}
	return hystrix.NewAnyErrorMiddleware[R](name)
}

func NewAny[R any](name string) grapper.AnyMiddleware[R] {
	ConfigAdd(name)
	config.Load()
	if o, _ := NewOptions(name); !o.Enabled {
		return nil
	}
	if err := h.ConfigureCommand(name); err != nil {
		log.Error(err.Error())
	}
	return hystrix.NewAnyMiddleware[R](name)
}

func NewError(name string) grapper.ErrorMiddleware {
	ConfigAdd(name)
	config.Load()
	if o, _ := NewOptions(name); !o.Enabled {
		return nil
	}
	if err := h.ConfigureCommand(name); err != nil {
		log.Error(err.Error())
	}
	return hystrix.NewErrorMiddleware(name)
}
