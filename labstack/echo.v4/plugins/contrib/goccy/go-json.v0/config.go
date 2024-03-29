package json

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/labstack/echo.v4"
)

const (
	root    = echo.PluginsRoot + ".goccyGoJson"
	enabled = ".enabled"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+enabled, true, "enable/disable goccy/go-json encoder")
}
