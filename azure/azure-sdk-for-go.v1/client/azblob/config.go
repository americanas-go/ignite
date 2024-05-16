package azblob

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/azure/azure-sdk-for-go.v1"
)

const (
	root        = azure.Root + ".azblob"
	accountName = ".accountName"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+accountName, "unknown", "defines the azure azblob account name", config.WithHide())
}
