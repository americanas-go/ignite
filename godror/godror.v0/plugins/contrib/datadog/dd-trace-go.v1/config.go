package datadog

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/godror/godror.v0"
)

const (
	root    = godror.PluginsRoot + ".datadog"
	enabled = ".enabled"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+enabled, true, "enable/disable datadog integration")
}
