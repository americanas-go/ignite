package fiber

import (
	"context"
	"sync"

	server "github.com/americanas-go/multiserver"
	f "github.com/gofiber/fiber/v2"
	contextfx "github.com/jvitoroc/ignite/go.uber.org/fx.v1/module/context"
	serverfx "github.com/jvitoroc/ignite/go.uber.org/fx.v1/module/core/server"
	"github.com/jvitoroc/ignite/gofiber/fiber.v2"
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
					Group: serverfx.ServersGroupKey,
					Target: func(srv *fiber.Server) server.Server {
						return srv
					},
				},
			),
		)
	})

	return options
}
