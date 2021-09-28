package vault

import "github.com/americanas-go/config"

const (
	root         = "ignite.vault"
	addr         = ".addr"
	caPath       = ".caPath"
	tp           = ".type"
	k8sRoot      = ".k8s"
	k8sRole      = k8sRoot + ".role"
	jwtRoot      = k8sRoot + ".jwt"
	jwtFile      = jwtRoot + ".file"
	jwtContent   = jwtRoot + ".content"
	token        = ".token"
	ManagersRoot = root + ".managers"

	secretPath       = ".secretPath"
	watcherRoot      = ".watcher"
	watcherEnabled   = watcherRoot + ".enabled"
	watcherIncrement = watcherRoot + ".increment"
	keys             = ".keys"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+addr, "", "defines vault address")
	config.Add(path+caPath, "", "defines ca path")
	config.Add(path+tp, "TOKEN", "defines type TOKEN/K8S/JWT")
	config.Add(path+k8sRole, "default", "defines k8s role")
	config.Add(path+jwtFile, "", "defines jwt file")
	config.Add(path+jwtContent, "", "defines jwt content")
	config.Add(path+token, "XPTOTOKEN", "defines vault token")
}

func ManagerConfigAdd(path string) {
	config.Add(path+secretPath, "", "sets manager vault secret path")
	config.Add(path+watcherEnabled, true, "enable/disable manager vault watcher")
	config.Add(path+watcherIncrement, 120, "defines increment on manager vault watcher")
	config.Add(path+keys, map[string]string{"username": "username", "password": "password"}, "defines keys map")
}
