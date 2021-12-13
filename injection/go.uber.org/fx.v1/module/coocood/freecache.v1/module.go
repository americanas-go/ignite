package freecache

import (
	"sync"

	"github.com/americanas-go/ignite/cache/coocood/freecache.v1"
	contextfx "github.com/americanas-go/ignite/injection/go.uber.org/fx.v1/module/context"
	"go.uber.org/fx"
)

var once sync.Once

// Module fx module for freecache.
func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {
		options = fx.Options(
			contextfx.Module(),
			fx.Provide(
				freecache.NewCache,
			),
		)
	})

	return options
}
