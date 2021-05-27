package fiber

import (
	"context"
	"sync"

	"github.com/americanas-go/ignite/go.uber.org/fx.v1/module/americanas-go/multiserver.v1"
	contextfx "github.com/americanas-go/ignite/go.uber.org/fx.v1/module/context"
	"github.com/americanas-go/ignite/gofiber/fiber.v2"
	server "github.com/americanas-go/multiserver"
	f "github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type params struct {
	fx.In
	Plugins []fiber.Plugin `optional:"true"`
}

var once sync.Once

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
