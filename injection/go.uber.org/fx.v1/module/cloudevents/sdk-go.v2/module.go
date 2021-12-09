package cloudevents

import (
	"sync"

	cloudevents "github.com/americanas-go/ignite/cloudevents/sdk-go.v2"
	contextfx "github.com/americanas-go/ignite/go.uber.org/fx.v1/module/context"
	"go.uber.org/fx"
)

var once sync.Once

// Module fx module for cloudevents.
func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {

		options = fx.Options(
			contextfx.Module(),
			fx.Provide(
				cloudevents.NewHTTP,
			),
		)
	})

	return options
}
