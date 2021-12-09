package godror

import (
	vault "github.com/americanas-go/ignite/mittwald/vaultgo.v0"
)

const (
	root = vault.ManagersRoot + ".godror"
)

func init() {
	vault.ManagerConfigAdd(root)
}
