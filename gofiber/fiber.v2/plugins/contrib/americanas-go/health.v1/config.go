package health

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/gofiber/fiber.v2"
)

const (
	root    = fiber.PluginsRoot + ".health"
	enabled = ".enabled"
	route   = ".route"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+enabled, true, "enable/disable health route")
	config.Add(path+route, "/health", "define status url")
}
