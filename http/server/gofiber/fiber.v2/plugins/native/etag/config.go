package etag

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/http/server/gofiber/fiber.v2"
)

const (
	root    = fiber.PluginsRoot + ".etag"
	enabled = ".enabled"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+enabled, true, "enable/disable etag middleware")
}
