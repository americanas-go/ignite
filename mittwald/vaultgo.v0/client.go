package vault

import (
	"context"

	"github.com/americanas-go/log"
	vault "github.com/mittwald/vaultgo"
)

// NewClientWithOptions returns a new vault client with options.
func NewClientWithOptions(ctx context.Context, options *Options) (*vault.Client, error) {

	var clientOpts vault.ClientOpts

	switch options.Type {
	case "TOKEN":
		clientOpts = vault.WithAuthToken(options.Token)
	case "K8S":
		var jwt vault.KubernetesAuthOpt
		if options.K8s.Jwt.File != "" {
			jwt = vault.WithJwtFromFile(options.K8s.Jwt.File)
		} else if options.K8s.Jwt.Content != "" {
			jwt = vault.WithJwtFromFile(options.K8s.Jwt.Content)
		}
		if jwt == nil {
			clientOpts = vault.WithKubernetesAuth(
				options.K8s.Role,
			)
		} else {
			clientOpts = vault.WithKubernetesAuth(
				options.K8s.Role,
				jwt,
			)
		}
	}

	client, err := vault.NewClient(
		options.Addr,
		vault.WithCaPath(options.CaPath),
		clientOpts,
	)

	if err != nil {
		return nil, err
	}

	return client, err
}

// NewClientWithConfigPath returns a new vault client with options from config path.
func NewClientWithConfigPath(ctx context.Context, path string) (*vault.Client, error) {
	options, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewClientWithOptions(ctx, options)
}

// NewClient returns a new vault client with default options.
func NewClient(ctx context.Context) (*vault.Client, error) {

	logger := log.FromContext(ctx)

	o, err := NewOptions()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	return NewClientWithOptions(ctx, o)
}
