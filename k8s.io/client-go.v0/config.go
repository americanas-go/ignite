package client

import "github.com/americanas-go/config"

const (
	root           = "ignite.k8s"
	kubeConfigPath = ".kubeConfigPath"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+kubeConfigPath, "~/.kube/config", "defines kubeconfig request")
}
