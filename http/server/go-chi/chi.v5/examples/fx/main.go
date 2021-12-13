package main

import (
	"context"
	"net/http"

	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/http/server/go-chi/chi.v5"
	"github.com/americanas-go/ignite/http/server/go-chi/chi.v5/plugins/contrib/americanas-go/health.v1"
	logplugin "github.com/americanas-go/ignite/http/server/go-chi/chi.v5/plugins/contrib/americanas-go/log.v1"
	multiserverplugin "github.com/americanas-go/ignite/http/server/go-chi/chi.v5/plugins/contrib/americanas-go/multi-server.v1"
	status "github.com/americanas-go/ignite/http/server/go-chi/chi.v5/plugins/contrib/americanas-go/rest-response.v1"
	datadog "github.com/americanas-go/ignite/http/server/go-chi/chi.v5/plugins/contrib/datadog/dd-trace-go.v1"
	newrelic "github.com/americanas-go/ignite/http/server/go-chi/chi.v5/plugins/contrib/newrelic/go-agent.v3"
	"github.com/americanas-go/ignite/http/server/go-chi/chi.v5/plugins/native/realip"
	"github.com/americanas-go/ignite/http/server/go-chi/chi.v5/plugins/native/recoverer"
	ifx "github.com/americanas-go/ignite/injection/go.uber.org/fx.v1"
	fxctx "github.com/americanas-go/ignite/injection/go.uber.org/fx.v1/module/context"
	fxchi "github.com/americanas-go/ignite/injection/go.uber.org/fx.v1/module/go-chi/chi.v5"
	ilog "github.com/americanas-go/ignite/log/americanas-go/log.v1"
	"github.com/americanas-go/log"
	"go.uber.org/fx"
)

func Get(ctx context.Context) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

func main() {

	config.Load()
	ilog.New()

	ifx.NewApp(
		fxctx.Module(),
		fxchi.Module(),
		fx.Provide(
			func() []chi.Plugin {
				return []chi.Plugin{
					multiserverplugin.Register,
					recoverer.Register,
					realip.Register,
					logplugin.Register,
					status.Register,
					health.Register,
					newrelic.Register,
					datadog.Register,
				}
			},
		),
		fx.Invoke(
			func(server *chi.Server, ctx context.Context) {
				server.Get("/hello", Get(ctx))
			},
		),
		fx.Invoke(
			func(lifecycle fx.Lifecycle, server *chi.Server) {
				lifecycle.Append(
					fx.Hook{
						OnStart: func(ctx context.Context) error {
							log.Info("starting server")
							go server.Serve(ctx)
							return nil
						},
						OnStop: func(ctx context.Context) error {
							log.Info("stopping server")
							server.Shutdown(ctx)
							return nil
						},
					},
				)
			},
		),
	).Run()

}
