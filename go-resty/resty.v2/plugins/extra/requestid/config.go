package requestid

import (
	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/go-resty/resty.v2"
)

const (
	root    = resty.PluginsRoot + ".requestid"
	enabled = root + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable requestId")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
