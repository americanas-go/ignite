package pprof

import (
	"context"

	"github.com/americanas-go/ignite/labstack/echo.v4"
	"github.com/americanas-go/log"
	echopprof "github.com/hiko1129/echo-pprof"
)

// Register registers a new pprof plugin for echo server.
func Register(ctx context.Context, server *echo.Server) error {
	o, err := NewOptions()
	if err != nil {
		return nil
	}
	h := NewPProfWithOptions(o)
	return h.Register(ctx, server)
}

// PProf represents pprof plugin for echo server.
type PProf struct {
	options *Options
}

// NewPProfWithOptions returns a new pprof plugin with options.
func NewPProfWithOptions(options *Options) *PProf {
	return &PProf{options: options}
}

// NewPProfWithConfigPath returns a new pprof plugin with options from config path.
func NewPProfWithConfigPath(path string) (*PProf, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewPProfWithOptions(o), nil
}

// NewPProf returns a new pprof plugin with default options.
func NewPProf() *PProf {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewPProfWithOptions(o)
}

// Register registers this pprof plugin for echo server.
func (i *PProf) Register(ctx context.Context, server *echo.Server) error {

	if !i.options.Enabled {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("configuring pprof in echo")

	echopprof.Wrap(server.Instance())

	logger.Debug("pprof configured with echo with success")

	return nil
}
