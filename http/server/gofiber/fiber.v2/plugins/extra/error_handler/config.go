package error_handler

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/gofiber/fiber.v2"
)

const (
	root    = fiber.PluginsRoot + ".errorHandler"
	enabled = ".enabled"
)

func init() {

}

func ConfigAdd(path string) {
	config.Add(path+enabled, true, "enable/disable error handler middleware")
}
