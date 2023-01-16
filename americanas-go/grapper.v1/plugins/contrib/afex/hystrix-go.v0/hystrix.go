package hystrix

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/grapper"
	"github.com/americanas-go/grapper/middleware/contrib/afex/hystrix-go.v0"
	h "github.com/americanas-go/ignite/afex/hystrix-go.v0"
	"github.com/americanas-go/log"
)

func New[R any](name string) grapper.Middleware[R] {
	ConfigAdd(name)
	config.Load()
	if o, _ := NewOptions(name); !o.Enabled {
		return nil
	}
	if err := h.ConfigureCommand(name); err != nil {
		log.Error(err.Error())
	}
	return hystrix.New[R](name)
}
