package datadog

import (
	"context"

	"github.com/americanas-go/log"
	"github.com/gofiber/fiber/v2"
	datadog "github.com/jvitoroc/ignite/datadog/dd-trace-go.v1"
	fibertrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gofiber/fiber.v2"
)

func Register(ctx context.Context, instance *fiber.App) error {

	if !IsEnabled() || !datadog.IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling datadog middleware in fiber")

	instance.Use(fibertrace.Middleware(fibertrace.WithServiceName(datadog.Service())))

	logger.Debug("datadog middleware successfully enabled in fiber")

	return nil
}
