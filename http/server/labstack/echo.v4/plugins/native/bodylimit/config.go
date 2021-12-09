package bodylimit

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/labstack/echo.v4"
)

const (
	root    = echo.PluginsRoot + ".bodyLimit"
	enabled = ".enabled"
	size    = ".size"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(enabled, true, "enable/disable body limit middleware")
	config.Add(size, "8M", "body limit size")
}
