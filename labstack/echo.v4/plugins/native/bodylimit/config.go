package bodylimit

import (
	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/labstack/echo.v4"
)

const (
	enabled = echo.PluginsRoot + ".bodylimit.enabled"
	size    = echo.PluginsRoot + ".bodylimit.size"
)

func init() {
	config.Add(enabled, true, "enable/disable body limit middleware")
	config.Add(size, "8M", "body limit size")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}

func GetSize() string {
	return config.String(size)
}
