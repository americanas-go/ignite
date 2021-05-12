package recover

import (
	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/gofiber/fiber.v2"
)

const (
	enabled = fiber.PluginsRoot + ".recover.enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable recover middleware")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
