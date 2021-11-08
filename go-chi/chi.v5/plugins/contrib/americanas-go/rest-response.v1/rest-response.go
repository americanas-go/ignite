package status

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/americanas-go/ignite/go-chi/chi.v5"
	"github.com/americanas-go/log"
	response "github.com/americanas-go/rest-response"
)

func Register(ctx context.Context) (*chi.Config, error) {
	l := NewStatus()
	return l.Register(ctx)
}

type Status struct {
	options *Options
}

func NewStatusWithOptions(options *Options) *Status {
	return &Status{options: options}
}

func NewStatusWithConfigPath(path string) (*Status, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewStatusWithOptions(o), nil
}

func NewStatus() *Status {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewStatusWithOptions(o)
}

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
