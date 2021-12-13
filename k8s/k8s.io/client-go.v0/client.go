package client

import (
	"context"

	"github.com/americanas-go/log"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// NewClientset returns a new kubernetes client set with default options.
func NewClientset(ctx context.Context) *kubernetes.Clientset {

	logger := log.FromContext(ctx)

	o, err := NewOptions()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	return NewClientsetWithOptions(ctx, o)
}

// NewClientsetWithConfigPath returns a new kubernetes client set with options from config path.
func NewClientsetWithConfigPath(ctx context.Context, path string) (*kubernetes.Clientset, error) {
	options, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewClientsetWithOptions(ctx, options), nil
}

// NewClientsetWithOptions returns a new kubernetes client set with options.
func NewClientsetWithOptions(ctx context.Context, options *Options) *kubernetes.Clientset {

	logger := log.FromContext(ctx).
		WithField("kubeConfigContext", options.KubeConfigContext).
		WithField("kubeConfigPath", options.KubeConfigPath)

	logger.Tracef("creating k8s client")

	var err error
	var config *rest.Config

	switch options.Type {
	case "INCLUSTER":
		config, err = rest.InClusterConfig()
	default:
		config, err = fromKubeConfig(options.KubeConfigContext, options.KubeConfigPath)
	}

	if err != nil {
		logger.Error(err.Error())
		return nil
	}

	var client *kubernetes.Clientset

	client, err = kubernetes.NewForConfig(config)
	if err != nil {
		logger.Error(err.Error())
		return nil
	}

	return client
}

func fromKubeConfig(context string, path string) (*rest.Config, error) {
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: path},
		&clientcmd.ConfigOverrides{
			CurrentContext: context,
		}).ClientConfig()
}
