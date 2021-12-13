package newrelic

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/http/client/go-resty/resty.v2"
)

const (
	root    = resty.PluginsRoot + ".newrelic"
	enabled = ".enabled"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+enabled, true, "enable/disable newrelic integration")
}
