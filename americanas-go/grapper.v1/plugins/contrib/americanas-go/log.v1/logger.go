package log

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/grapper"
	"github.com/americanas-go/grapper/middleware/contrib/americanas-go/log.v1"
)

func New[R any](name string) grapper.Middleware[R] {
	ConfigAdd(name)
	config.Load()
	if o, _ := NewOptions(name); !o.Enabled {
		return nil
	}
	return log.New[R]()
}
