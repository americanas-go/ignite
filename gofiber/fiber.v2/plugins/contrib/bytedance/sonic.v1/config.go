package json

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/gofiber/fiber.v2"
)

const (
	root    = fiber.PluginsRoot + ".bytedanceSonic"
	enabled = ".enabled"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+enabled, true, "enable/disable bytedance/sonic encoder")
}
