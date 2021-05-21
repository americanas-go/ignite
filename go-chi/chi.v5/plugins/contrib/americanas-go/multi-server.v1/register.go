package multiserver

import (
	"context"
	"net/http"

	"github.com/americanas-go/ignite/go-chi/chi.v5"
	"github.com/americanas-go/log"
	"github.com/americanas-go/multiserver"
)

func Register(ctx context.Context) (*chi.Config, error) {
	if !IsEnabled() {
		return nil, nil
	}

	logger := log.FromContext(ctx)

	checkRoute := getRoute()

	logger.Tracef("configuring multi server check router on %s in chi", checkRoute)

	return &chi.Config{
		Routes: []chi.ConfigRouter{
			{
				Method:      http.MethodGet,
				HandlerFunc: Get(ctx),
				Pattern:     checkRoute,
			},
		},
	}, nil
}

func Get(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		status := http.StatusOK

		if err := multiserver.Check(ctx); err != nil {
			status = http.StatusServiceUnavailable
			w.WriteHeader(status)
			w.Write([]byte("service unavailable"))
			return
		}

		w.WriteHeader(status)

	}
}
