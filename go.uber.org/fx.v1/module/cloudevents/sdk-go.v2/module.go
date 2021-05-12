package cloudevents

import (
	"sync"

	cloudevents "github.com/jvitoroc/ignite/cloudevents/sdk-go.v2"
	contextfx "github.com/jvitoroc/ignite/go.uber.org/fx.v1/module/context"
	"go.uber.org/fx"
)

var once sync.Once

func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {

		options = fx.Options(
			contextfx.Module(),
			fx.Provide(
				cloudevents.NewDefaultClient,
			),
		)
	})

	return options
}
