package client

import "github.com/americanas-go/config"

const (
	root           = "ignite.k8s"
	kubeConfigPath = root + ".kubeConfigPath"
)

func init() {
	config.Add(kubeConfigPath, "~/.kube/config", "defines kubeconfig request")
}
