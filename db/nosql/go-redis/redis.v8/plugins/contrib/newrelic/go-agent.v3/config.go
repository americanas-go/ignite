package newrelic

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/db/nosql/go-redis/redis.v8"
)

const (
	root    = redis.PluginsRoot + ".newrelic"
	enabled = ".enabled"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+enabled, true, "enable/disable newrelic integration")
}
