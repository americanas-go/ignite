package datadog

import (
	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/go.mongodb.org/mongo-driver.v1"
)

const (
	root    = mongo.PluginsRoot + ".datadog"
	enabled = root + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable datadog integration")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
