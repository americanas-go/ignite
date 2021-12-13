package bodydump

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/http/server/labstack/echo.v4"
)

const (
	root    = echo.PluginsRoot + ".bodydump"
	enabled = ".enabled"
)

func init() {

}

func ConfigAdd(path string) {
	config.Add(path+enabled, true, "enable/disable body dump middleware")
}
