package datadog

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/db/nosql/go.mongodb.org/mongo-driver.v1"
)

const (
	root    = mongo.PluginsRoot + ".datadog"
	enabled = ".enabled"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+enabled, true, "enable/disable datadog integration")
}
