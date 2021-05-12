package semaphore

import (
	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/labstack/echo.v4"
)

const (
	semaphoreRoot = echo.PluginsRoot + ".semaphore"
	enabled       = semaphoreRoot + ".enabled"
	limit         = semaphoreRoot + ".limit"
)

func init() {
	config.Add(enabled, true, "enable/disable semaphore middleware")
	config.Add(limit, 10000, "defines numbers for concurrent connections")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}

func GetLimit() int {
	return config.Int(limit)
}
