package requestid

import (
	"context"
	"net/http"

	"github.com/americanas-go/ignite/go-chi/chi.v5"
	"github.com/americanas-go/log"
	"github.com/go-chi/chi/v5/middleware"
)

// Register registers requestID middleware for chi.
func Register(ctx context.Context) (*chi.Config, error) {
	if !IsEnabled() {
		return nil, nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling requestID middleware in chi")

	return &chi.Config{
		Middlewares: []func(http.Handler) http.Handler{
			middleware.RequestID,
		},
	}, nil
}
