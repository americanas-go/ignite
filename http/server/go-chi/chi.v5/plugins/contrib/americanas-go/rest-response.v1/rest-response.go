package status

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/americanas-go/ignite/http/server/go-chi/chi.v5"
	"github.com/americanas-go/log"
	response "github.com/americanas-go/rest-response"
)

// Register registers a default status to a new chi config.
func Register(ctx context.Context) (*chi.Config, error) {
	l := NewStatus()
	return l.Register(ctx)
}

// Status struct that represents a status router plugin.
type Status struct {
	options *Options
}

// NewStatusWithOptions returns a status router with options.
func NewStatusWithOptions(options *Options) *Status {
	return &Status{options: options}
}

// NewStatusWithOptions returns a status router with options from config path.
func NewStatusWithConfigPath(path string) (*Status, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewStatusWithOptions(o), nil
}

// NewStatus returns a status router with default options.
func NewStatus() *Status {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewStatusWithOptions(o)
}

// Register registers the status router to a new chi config.
func (i *Status) Register(ctx context.Context) (*chi.Config, error) {
	if !i.options.Enabled {
		return nil, nil
	}

	logger := log.FromContext(ctx)

	statusRoute := i.options.Route

	logger.Tracef("configuring status router on %s in chi", statusRoute)

	return &chi.Config{
		Routes: []chi.ConfigRouter{
			{
				Method:      http.MethodGet,
				HandlerFunc: handler(),
				Pattern:     statusRoute,
			},
		},
	}, nil
}

func handler() http.HandlerFunc {
	resourceStatus := response.NewResourceStatus()
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(resourceStatus)

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(reqBodyBytes.Bytes())
		w.WriteHeader(http.StatusOK)
	}
}
