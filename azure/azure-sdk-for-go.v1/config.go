package azure

import (
	"github.com/americanas-go/config"
)

const (
	Root         = "ignite.azure"
	clientId     = ".clientId"
	clientSecret = ".clientSecret"
	tenantId     = ".tenantId"
)

func init() {
	ConfigAdd(Root)
}

func ConfigAdd(path string) {
	config.Add(path+clientId, "", "defines the azure client id", config.WithHide())
	config.Add(path+clientSecret, "", "defines the azure client secret", config.WithHide())
	config.Add(path+tenantId, "", "defines the azure tenant id")
}
