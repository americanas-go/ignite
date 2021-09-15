package swagger

import (
	"context"

	"github.com/americanas-go/ignite/labstack/echo.v4"
	"github.com/americanas-go/log"
	eswagger "github.com/swaggo/echo-swagger"
)

func Register(ctx context.Context, server *echo.Server) error {
	o, err := NewOptions()
	if err != nil {
		return nil
	}
	h := NewSwaggerWithOptions(o)
	return h.Register(ctx, server)
}

type Swagger struct {
	options *Options
}

func NewSwaggerWithOptions(options *Options) *Swagger {
	return &Swagger{options: options}
}

func NewSwaggerWithConfigPath(path string) (*Swagger, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewSwaggerWithOptions(o), nil
}

func NewSwagger() *Swagger {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewSwaggerWithOptions(o)
}

func (i *Swagger) Register(ctx context.Context, server *echo.Server) error {

	if !IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	swaggerRoute := GetRoute()

	logger.Tracef("configuring swagger router on %s in echo", swaggerRoute)

	server.GET(swaggerRoute, eswagger.WrapHandler)

	logger.Debugf("swagger router configured on %s in echo", swaggerRoute)

	return nil
}
