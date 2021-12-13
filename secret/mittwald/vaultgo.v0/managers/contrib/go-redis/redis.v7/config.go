package redis

import (
	vault "github.com/americanas-go/ignite/secret/mittwald/vaultgo.v0"
)

const (
	root = vault.ManagersRoot + ".redis"
)

func init() {
	vault.ManagerConfigAdd(root)
}
