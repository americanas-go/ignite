package newrelic

import (
	"sync"

	newrelic "github.com/americanas-go/ignite/apm/newrelic/go-agent.v3"
	contextfx "github.com/americanas-go/ignite/injection/go.uber.org/fx.v1/module/context"
	"go.uber.org/fx"
)

var once sync.Once

// Module fx module for newrelic agent.
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
