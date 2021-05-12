package freecache

import (
	"sync"

	"github.com/jvitoroc/ignite/coocood/freecache.v1"
	contextfx "github.com/jvitoroc/ignite/go.uber.org/fx.v1/module/context"
	"go.uber.org/fx"
)

var once sync.Once

func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {
		options = fx.Options(
			contextfx.Module(),
			fx.Invoke(
				freecache.NewCache,
			),
		)
	})

	return options
}
