package bodydump

import (
	"context"

	"github.com/americanas-go/ignite/labstack/echo.v4"
	"github.com/americanas-go/log"
	e "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Register(ctx context.Context, server *echo.Server) error {
	if !IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("enabling body dump middleware in echo")

	server.Use(middleware.BodyDump(bodyDump))

	logger.Debug("body dump middleware successfully enabled in echo")

	return nil
}

func bodyDump(c e.Context, reqBody []byte, resBody []byte) {
	logger := log.FromContext(c.Request().Context())
	logger.Info("request body --->")
	logger.Info(string(reqBody))
	logger.Info("response body -->")
	logger.Info(string(resBody))
}
