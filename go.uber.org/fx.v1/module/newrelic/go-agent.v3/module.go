package newrelic

import (
	"sync"

	contextfx "github.com/americanas-go/ignite/go.uber.org/fx.v1/module/context"
	newrelic "github.com/americanas-go/ignite/newrelic/go-agent.v3"
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
