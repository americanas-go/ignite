package datadog

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/go.mongodb.org/mongo-driver.v1"
)

const (
	root          = mongo.PluginsRoot + ".datadog"
	enabled       = ".enabled"
	serviceName   = ".serviceName"
	analytics     = ".analytics"
	analyticsRate = ".analyticsRate"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+enabled, true, "enable/disable datadog integration")
	config.Add(path+serviceName, "mongo", "sets serviceName datadog integration")
	config.Add(path+analytics, false, "enable/disable analytics datadog integration")
	config.Add(path+analyticsRate, 1.0, "sets analytics rate datadog integration")
}
