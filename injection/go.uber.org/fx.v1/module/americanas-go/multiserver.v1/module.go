package multiserver

import (
	"context"
	"sync"

	"github.com/americanas-go/ignite/cmd/spf13/cobra.v1"
	contextfx "github.com/americanas-go/ignite/injection/go.uber.org/fx.v1/module/context"
	server "github.com/americanas-go/multiserver"
	c "github.com/spf13/cobra"
	"go.uber.org/fx"
)

const (
	ServersGroupKey = "_gi_server_servers_"
)

type srvParams struct {
	fx.In
	Servers []server.Server `group:"_gi_server_servers_"`
}

var once sync.Once

// Module fx module for multiserver.
func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {

		options = fx.Options(
			contextfx.Module(),
			fx.Invoke(
				func(ctx context.Context, p srvParams) error {

					return cobra.Run(
						&c.Command{
							Run: func(cmd *c.Command, args []string) {
								server.Serve(ctx, p.Servers...)
							},
						},
					)

				},
			),
		)
	})

	return options
}
