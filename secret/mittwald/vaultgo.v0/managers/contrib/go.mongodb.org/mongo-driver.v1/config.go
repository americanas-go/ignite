package mongo

import (
	vault "github.com/americanas-go/ignite/secret/mittwald/vaultgo.v0"
)

const (
	root = vault.ManagersRoot + ".mongo"
)

func init() {
	vault.ManagerConfigAdd(root)
}
