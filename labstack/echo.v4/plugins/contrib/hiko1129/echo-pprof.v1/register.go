package echo_pprof_v1

import (
	"context"

	"github.com/americanas-go/log"
	echopprof "github.com/hiko1129/echo-pprof"
	"github.com/labstack/echo/v4"
)

func Register(ctx context.Context, instance *echo.Echo) error {

	if !IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("configuring pprof in echo")

	echopprof.Wrap(instance)

	logger.Debug("pprof configured with echo with success")

	return nil
}
