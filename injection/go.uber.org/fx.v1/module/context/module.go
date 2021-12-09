package context

import (
	"context"
	"sync"

	"go.uber.org/fx"
)

var once sync.Once

// Module fx module for context.
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
