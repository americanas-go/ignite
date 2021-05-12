package stripslashes

import (
	"context"
	"net/http"

	"github.com/americanas-go/log"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jvitoroc/ignite/go-chi/chi.v5"
)

func Register(ctx context.Context) (*chi.Config, error) {

	if !IsEnabled() {
		return nil, nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling stripslashes middleware in chi")

	return &chi.Config{
		Middlewares: []func(http.Handler) http.Handler{
			middleware.StripSlashes,
		},
	}, nil
}
