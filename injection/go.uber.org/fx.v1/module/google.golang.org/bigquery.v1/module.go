package bigquery

import (
	"sync"

	"github.com/americanas-go/ignite/gcp/google.golang.org/bigquery.v1"
	contextfx "github.com/americanas-go/ignite/injection/go.uber.org/fx.v1/module/context"
	"go.uber.org/fx"
)

var once sync.Once

// Module fx module for bigQuery client.
func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {

		options = fx.Options(
			contextfx.Module(),
			fx.Provide(
				bigquery.NewClient,
			),
		)

	})

	return options
}
