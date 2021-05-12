package newrelic

import (
	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/go.mongodb.org/mongo-driver.v1"
)

const (
	root    = mongo.PluginsRoot + ".newrelic"
	enabled = root + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable mongo integration")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
