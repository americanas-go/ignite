package context

import (
	"context"
	"sync"

	"go.uber.org/fx"
)

var once sync.Once

func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {
		options = fx.Options(
			fx.Provide(
				context.Background,
			),
		)
	})

	return options
}
