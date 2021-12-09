package server

import (
	"context"
	"sync"

	"github.com/americanas-go/ignite/go.uber.org/fx.v1/module/americanas-go/multiserver.v1"
	contextfx "github.com/americanas-go/ignite/go.uber.org/fx.v1/module/context"
	"github.com/americanas-go/ignite/google.golang.org/grpc.v1/server"
	s "github.com/americanas-go/multiserver"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type params struct {
	fx.In
	Plugins []server.Plugin `optional:"true"`
}

var once sync.Once

// Module fx module for gRPC server.
func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {

		options = fx.Options(
			contextfx.Module(),
			fx.Provide(
				func(ctx context.Context, p params) *server.Server {
					return server.NewServer(ctx, p.Plugins...)
				},
				func(srv *server.Server) *grpc.Server {
					return srv.Server()
				},
				func(srv *server.Server) grpc.ServiceRegistrar {
					return srv.ServiceRegistrar()
				},
				fx.Annotated{
					Group: multiserver.ServersGroupKey,
					Target: func(srv *server.Server) s.Server {
						return srv
					},
				},
			),
		)

	})

	return options
}
