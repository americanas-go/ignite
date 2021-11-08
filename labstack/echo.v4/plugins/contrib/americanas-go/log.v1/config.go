package log

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/labstack/echo.v4"
)

const (
	root    = echo.PluginsRoot + ".logger"
	enabled = ".enabled"
	level   = ".level"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+enabled, true, "enable/disable logging request middleware")
	config.Add(path+level, "DEBUG", "sets log level INFO/DEBUG/TRACE")
}
