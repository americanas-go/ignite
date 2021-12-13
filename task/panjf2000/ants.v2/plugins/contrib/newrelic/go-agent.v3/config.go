package newrelic

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/panjf2000/ants.v2"
)

const (
	root    = ants.PluginsRoot + ".newrelic"
	enabled = ".enabled"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+enabled, true, "enable/disable newrelic integration")
}
