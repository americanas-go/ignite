package redis

import (
	"context"
	e "github.com/americanas-go/ignite/elastic/go-elasticsearch.v8"
	"github.com/elastic/go-elasticsearch/v8"
	"sync"

	contextfx "github.com/americanas-go/ignite/go.uber.org/fx.v1/module/context"
	"go.uber.org/fx"
)

var once sync.Once

type clientParams struct {
	fx.In
	Plugins []e.Plugin `optional:"true"`
}

// Module fx module for elasticsearch client.
func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {

		options = fx.Options(
			contextfx.Module(),
			fx.Provide(
				func(ctx context.Context, p clientParams) (*elasticsearch.Client, error) {
					return e.NewClient(ctx, p.Plugins...)
				},
			),
		)
	})

	return options
}

var bonce sync.Once

// BulkModule fx module for elasticsearch bulk indexer.
func BulkModule() fx.Option {
	options := fx.Options()

	bonce.Do(func() {
		options = fx.Options(
			Module(),
			fx.Provide(
				e.NewBulkIndexer,
			),
		)
	})

	return options
}
