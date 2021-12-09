package monitor

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/gofiber/fiber.v2"
)

const (
	root    = fiber.PluginsRoot + ".monitor"
	enabled = ".enabled"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+enabled, true, "enable/disable monitor middleware")
}
