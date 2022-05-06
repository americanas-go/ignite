package gocql

import (
	vault "github.com/americanas-go/ignite/mittwald/vaultgo.v0"
)

const (
	root = vault.ManagersRoot + ".gocql"
)

func init() {
	vault.ManagerConfigAdd(root)
}
