package multiserver

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/http/server/gofiber/fiber.v2"
)

const (
	root    = fiber.PluginsRoot + ".multiServer"
	enabled = ".enabled"
	route   = ".route"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+enabled, true, "enable/disable multi server check route")
	config.Add(path+route, "/check", "define multi server check url")
}
