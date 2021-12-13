package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/http/server/go-chi/chi.v5"
	multiserverplugin "github.com/americanas-go/ignite/http/server/go-chi/chi.v5/plugins/contrib/americanas-go/multi-server.v1"
	"github.com/americanas-go/ignite/http/server/go-chi/chi.v5/plugins/native/recoverer"
	"github.com/americanas-go/ignite/log/americanas-go/log.v1"
	"github.com/americanas-go/multiserver"
)

func Get(ctx context.Context) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

func main() {

	config.Load()
	log.New()

	chiCtx := context.Background()
	chiSrv := chi.NewServer(chiCtx,
		multiserverplugin.Register,
		recoverer.Register,
	)

	chiSrv.Get("/hello", Get(chiCtx))

	msCtx := context.Background()
	multiserver.Serve(msCtx, chiSrv, &LocalServer{})
}

type LocalServer struct {
}

func (s *LocalServer) Serve(ctx context.Context) {
	time.Sleep(30 * time.Second)
	fmt.Printf("finished")
}

func (s *LocalServer) Shutdown(ctx context.Context) {
}
