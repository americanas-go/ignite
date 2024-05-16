package azure

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

// NewCredential returns azure credential.
func NewCredential(ctx context.Context) (azcore.TokenCredential, error) {
	o, err := NewOptions()
	if err != nil {
		return nil, err
	}

	return NewCredentialWithOptions(ctx, o)
}

// NewCredentialWithConfigPath NewConfigWithOptions returns azure client with options from config path.
func NewCredentialWithConfigPath(ctx context.Context, path string) (azcore.TokenCredential, error) {
	opts, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewCredentialWithOptions(ctx, opts)
}

// NewCredentialWithOptions returns azure client with options.
func NewCredentialWithOptions(ctx context.Context, options *Options) (azcore.TokenCredential, error) {
	if options.ClientId != "" && options.ClientSecret != "" && options.TenantId != "" {
		return azidentity.NewClientSecretCredential(options.ClientId, options.ClientSecret, options.TenantId, nil)
	}
	return azidentity.NewDefaultAzureCredential(nil)
}
