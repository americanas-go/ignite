package error_handler

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/gofiber/fiber.v2"
)

const (
	root            = fiber.PluginsRoot + ".errorHandler"
	enabled         = ".enabled"
	logger          = ".logger"
	loggerEnabled   = logger + ".enabled"
	print4xx        = logger + ".print4xx"
	print5xx        = logger + ".print5xx"
	printStackTrace = logger + ".printStackTrace"
)

func init() {

}

func ConfigAdd(path string) {
	config.Add(path+enabled, true, "enable/disable error handler middleware")
	config.Add(path+loggerEnabled, true, "enable/disable print error log")
	config.Add(path+print4xx, true, "enable/disable print 4xx errors")
	config.Add(path+print5xx, false, "enable/disable error 5xx errors")
	config.Add(path+printStackTrace, false, "enable/disable error print stacktrace")
}
