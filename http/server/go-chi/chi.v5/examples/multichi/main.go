package main

import (
	"context"
	"net/http"

	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/http/server/go-chi/chi.v5"
	multiserverplugin "github.com/americanas-go/ignite/http/server/go-chi/chi.v5/plugins/contrib/americanas-go/multi-server.v1"
	"github.com/americanas-go/ignite/http/server/go-chi/chi.v5/plugins/native/recoverer"
	"github.com/americanas-go/ignite/http/server/net/server"
	"github.com/americanas-go/ignite/log/americanas-go/log.v1"
	"github.com/americanas-go/multiserver"
)

const httpServerRoot = "ignite.http2.server"

func init() {
	server.ConfigAdd(httpServerRoot)
}

func Get(ctx context.Context) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

func main() {

	config.Load()
	log.New()

	chi1Ctx := context.Background()
	chi1Srv := chi.NewServer(chi1Ctx,
		multiserverplugin.Register,
		recoverer.Register,
	)

	chi1Srv.Get("/hello", Get(chi1Ctx))

	srv2Options, err := server.NewOptionsWithPath(httpServerRoot)
	if err != nil {
		panic(err)
	}

	chi2Ctx := context.Background()
	chi2Srtv := chi.NewServerWithOptions(chi2Ctx, srv2Options)

	msCtx := context.Background()
	multiserver.Serve(msCtx, chi1Srv, chi2Srtv)
}
