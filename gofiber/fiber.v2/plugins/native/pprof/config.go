package pprof

import (
	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/gofiber/fiber.v2"
)

const (
	enabled = fiber.PluginsRoot + ".pprof.enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable pprof middleware")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
