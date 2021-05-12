package newrelic

import (
	"sync"

	contextfx "github.com/jvitoroc/ignite/go.uber.org/fx.v1/module/context"
	newrelic "github.com/jvitoroc/ignite/newrelic/go-agent.v3"
	"go.uber.org/fx"
)

var once sync.Once

func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {
		options = fx.Options(
			contextfx.Module(),
			fx.Invoke(
				newrelic.NewApplication,
			),
		)
	})

	return options
}
