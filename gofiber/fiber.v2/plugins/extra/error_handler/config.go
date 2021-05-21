package error_handler

import (
	"github.com/americanas-go/config"
	gifiber "github.com/americanas-go/ignite/gofiber/fiber.v2"
)

const (
	root    = gifiber.PluginsRoot + ".errorHandler"
	enabled = root + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable error handler middleware")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
