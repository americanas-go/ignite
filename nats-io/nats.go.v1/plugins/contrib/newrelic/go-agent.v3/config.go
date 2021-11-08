package newrelic

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/nats-io/nats.go.v1"
)

const (
	root    = nats.PluginsRoot + ".newrelic"
	enabled = ".enabled"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+enabled, true, "enable/disable newrelic")
}
