package elasticsearch

import (
	vault "github.com/americanas-go/ignite/secret/mittwald/vaultgo.v0"
)

const (
	root = vault.ManagersRoot + ".elasticsearch"
)

func init() {
	vault.ManagerConfigAdd(root)
}
