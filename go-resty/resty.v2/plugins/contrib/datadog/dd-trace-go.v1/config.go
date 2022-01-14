package datadog

import (
	"github.com/americanas-go/config"
	girest "github.com/americanas-go/ignite/go-resty/resty.v2"
)

const (
	root          = girest.PluginsRoot + ".datadog"
	operationName = ".operationName"
	enabled       = ".enabled"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+operationName, "http.request", "defines span operation name")
	config.Add(path+enabled, true, "enable/disable datadog integration")
}
