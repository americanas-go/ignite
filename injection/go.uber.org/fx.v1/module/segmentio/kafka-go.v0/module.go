package kafka

import (
	"sync"

	contextfx "github.com/americanas-go/ignite/go.uber.org/fx.v1/module/context"
	"github.com/americanas-go/ignite/segmentio/kafka-go.v0"
	"go.uber.org/fx"
)

var once sync.Once

// Module fx module for kafka connection.
func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {
		options = fx.Options(
			contextfx.Module(),
			fx.Provide(
				kafka.NewConn,
			),
		)
	})

	return options
}
