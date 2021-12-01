package swagger

import (
	"context"

	"github.com/americanas-go/ignite/labstack/echo.v4"
	"github.com/americanas-go/log"
	eswagger "github.com/swaggo/echo-swagger"
)

// Register registers a new swagger plugin for echo server.
func Register(ctx context.Context, server *echo.Server) error {
	o, err := NewOptions()
	if err != nil {
		return nil
	}
	h := NewSwaggerWithOptions(o)
	return h.Register(ctx, server)
}

// Swagger represents swagger plugin for echo server.
type Swagger struct {
	options *Options
}

// NewSwaggerWithOptions returns a new swagger plugin with options.
func NewSwaggerWithOptions(options *Options) *Swagger {
	return &Swagger{options: options}
}

// NewSwaggerWithConfigPath returns a new swagger plugin with options from config path.
func NewSwaggerWithConfigPath(path string) (*Swagger, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewSwaggerWithOptions(o), nil
}

// NewSwagger returns a new swagger plugin with default options.
func NewSwagger() *Swagger {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewSwaggerWithOptions(o)
}

// Register registers this swagger plugin for echo server.
func (i *Swagger) Register(ctx context.Context, server *echo.Server) error {

	if !i.options.Enabled {
		return nil
	}

	logger := log.FromContext(ctx)

	swaggerRoute := i.options.Route

	logger.Tracef("configuring swagger router on %s in echo", swaggerRoute)

	server.GET(swaggerRoute, eswagger.WrapHandler)

	logger.Debugf("swagger router configured on %s in echo", swaggerRoute)

	return nil
}
