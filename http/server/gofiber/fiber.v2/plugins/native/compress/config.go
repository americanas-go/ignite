package compress

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/http/server/gofiber/fiber.v2"
)

const (
	root    = fiber.PluginsRoot + ".compress"
	enabled = ".enabled"
	level   = ".level"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+enabled, true, "enable/disable compress middleware")
	config.Add(path+level, 0, "compress level (disabled: -1, default: 0, best speed: 1, best compression: 2)")
}
