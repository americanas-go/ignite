package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/http/server/go-chi/chi.v5"
	"github.com/americanas-go/ignite/http/server/go-chi/chi.v5/plugins/contrib/americanas-go/health.v1"
	"github.com/americanas-go/ignite/http/server/go-chi/chi.v5/plugins/contrib/americanas-go/log.v1"
	status "github.com/americanas-go/ignite/http/server/go-chi/chi.v5/plugins/contrib/americanas-go/rest-response.v1"
	"github.com/americanas-go/ignite/http/server/go-chi/chi.v5/plugins/native/realip"
	"github.com/americanas-go/ignite/http/server/go-chi/chi.v5/plugins/native/recoverer"
	ilog "github.com/americanas-go/ignite/log/americanas-go/log.v1"
)

const HelloWorldEndpoint = "app.endpoint.helloworld"

func init() {
	config.Add(HelloWorldEndpoint, "/hello-world", "helloworld endpoint")
}

type Config struct {
	App struct {
		Endpoint struct {
			Helloworld string
		}
	}
}

type Response struct {
	Message string
}

func Get(ctx context.Context) http.HandlerFunc {

	resp := Response{
		Message: "Hello World!!",
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}

func main() {

	config.Load()

	c := Config{}

	err := config.Unmarshal(&c)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	ilog.New()

	srv := chi.NewServer(ctx,
		recoverer.Register,
		realip.Register,
		log.Register,
		status.Register,
		health.Register)

	srv.Get(c.App.Endpoint.Helloworld, Get(ctx))

	srv.Serve(ctx)
}
