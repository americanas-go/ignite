package datadog

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/data/nosql/go-redis/redis.v7"
)

const (
	root    = redis.PluginsRoot + ".datadog"
	enabled = ".enabled"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+enabled, true, "enable/disable datadog integration")
}
