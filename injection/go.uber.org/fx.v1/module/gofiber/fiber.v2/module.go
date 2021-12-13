package fiber

import (
	"context"
	"sync"

	"github.com/americanas-go/ignite/http/server/gofiber/fiber.v2"
	"github.com/americanas-go/ignite/injection/go.uber.org/fx.v1/module/americanas-go/multiserver.v1"
	contextfx "github.com/americanas-go/ignite/injection/go.uber.org/fx.v1/module/context"
	server "github.com/americanas-go/multiserver"
	f "github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type params struct {
	fx.In
	Plugins []fiber.Plugin `optional:"true"`
}

var once sync.Once

// Module fx module for fiber server.
func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {

		options = fx.Options(
			contextfx.Module(),
			fx.Provide(
				func(ctx context.Context, p params) *fiber.Server {
					return fiber.NewServer(ctx, p.Plugins...)
				},
				func(srv *fiber.Server) *f.App {
					return srv.App()
				},
			),
			fx.Provide(
				fx.Annotated{
					Group: multiserver.ServersGroupKey,
					Target: func(srv *fiber.Server) server.Server {
						return srv
					},
				},
			),
		)
	})

	return options
}
