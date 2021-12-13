package redis

import (
	"context"
	"sync"

	"github.com/americanas-go/ignite/db/nosql/go-redis/redis.v7"
	contextfx "github.com/americanas-go/ignite/injection/go.uber.org/fx.v1/module/context"
	r "github.com/go-redis/redis/v7"
	"go.uber.org/fx"
)

var once sync.Once

type clusterParams struct {
	fx.In
	Plugins []redis.ClusterPlugin `optional:"true"`
}

// ClusterModule fx module for redis cluster client.
func ClusterModule() fx.Option {
	options := fx.Options()

	once.Do(func() {

		options = fx.Options(
			contextfx.Module(),
			fx.Provide(
				func(ctx context.Context, p clusterParams) (*r.ClusterClient, error) {
					return redis.NewClusterClient(ctx, p.Plugins...)
				},
			),
		)
	})

	return options
}

type clientParams struct {
	fx.In
	Plugins []redis.Plugin `optional:"true"`
}

// ClientModule fx module for redis client.
func ClientModule() fx.Option {
	options := fx.Options()

	once.Do(func() {

		options = fx.Options(
			contextfx.Module(),
			fx.Provide(
				func(ctx context.Context, p clientParams) (*r.Client, error) {
					return redis.NewClient(ctx, p.Plugins...)
				},
			),
		)
	})

	return options
}
