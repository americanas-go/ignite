package echo_pprof_v1

import (
	"context"

	"github.com/americanas-go/ignite/labstack/echo.v4"
	"github.com/americanas-go/log"
	echopprof "github.com/hiko1129/echo-pprof"
)

func Register(ctx context.Context, server *echo.Server) error {

	if !IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("configuring pprof in echo")

	echopprof.Wrap(server.Instance())

	logger.Debug("pprof configured with echo with success")

	return nil
}
