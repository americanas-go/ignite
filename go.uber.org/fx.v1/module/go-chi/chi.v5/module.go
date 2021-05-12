package chi

import (
	"context"
	"sync"

	server "github.com/americanas-go/multiserver"
	c "github.com/go-chi/chi/v5"
	"github.com/jvitoroc/ignite/go-chi/chi.v5"
	contextfx "github.com/jvitoroc/ignite/go.uber.org/fx.v1/module/context"
	serverfx "github.com/jvitoroc/ignite/go.uber.org/fx.v1/module/core/server"
	"go.uber.org/fx"
)

var once sync.Once

type params struct {
	fx.In
	Plugins []chi.Plugin `optional:"true"`
}

func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {

		options = fx.Options(
			contextfx.Module(),
			fx.Provide(
				func(ctx context.Context, p params) *chi.Server {
					return chi.NewServer(ctx, p.Plugins...)
				},
				func(srv *chi.Server) *c.Mux {
					return srv.Mux()
				},
				fx.Annotated{
					Group: serverfx.ServersGroupKey,
					Target: func(srv *chi.Server) server.Server {
						return srv
					},
				},
			),
		)
	})

	return options
}
