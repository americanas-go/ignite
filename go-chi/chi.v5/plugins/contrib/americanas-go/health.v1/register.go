package health

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/americanas-go/ignite/go-chi/chi.v5"
	"github.com/americanas-go/log"
	response "github.com/americanas-go/rest-response"
)

// Register registers health check handler for chi.
func Register(ctx context.Context) (*chi.Config, error) {
	if !IsEnabled() {
		return nil, nil
	}

	logger := log.FromContext(ctx)

	healthRoute := getRoute()

	logger.Tracef("configuring health router on %s in chi", healthRoute)

	healthHandler := NewHealthHandler()

	return &chi.Config{
		Routes: []chi.ConfigRouter{
			{
				Method:      http.MethodGet,
				HandlerFunc: healthHandler.Get(ctx),
				Pattern:     healthRoute,
			},
		},
	}, nil
}

// NewHealthHandler returns a new HealthHandler.
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// HealthHandler represents a health check handler.
type HealthHandler struct {
}

// Get returns a http handler for health check.
func (u *HealthHandler) Get(ctx context.Context) http.HandlerFunc {
	resp, httpCode := response.NewHealth(ctx)
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(httpCode)
		json.NewEncoder(w).Encode(resp)
	}
}
