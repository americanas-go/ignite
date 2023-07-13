package nats

import (
	"sync"

	contextfx "github.com/americanas-go/ignite/go.uber.org/fx.v1/module/context"
	"github.com/americanas-go/ignite/nats-io/nats.go.v1"
	"go.uber.org/fx"
)

var once sync.Once

// Module fx module for nats.
func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {
		options = fx.Options(
			contextfx.Module(),
			fx.Provide(
				nats.NewConn,
			),
		)
	})

	return options
}
