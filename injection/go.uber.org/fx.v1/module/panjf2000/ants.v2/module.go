package ants

import (
	"sync"

	ants "github.com/americanas-go/ignite/task/panjf2000/ants.v2"
	a "github.com/panjf2000/ants/v2"
	"go.uber.org/fx"
)

var once sync.Once

// Module fx module for ants wrapper.
func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {

		options = fx.Options(
			fx.Provide(
				func(pool *a.Pool, m []ants.Middleware) *ants.Wrapper {
					return ants.NewWrapper(pool, m...)
				},
			),
		)
	})

	return options
}
