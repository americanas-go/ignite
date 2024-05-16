package azblob

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

// NewClient returns .
func NewClient(ctx context.Context, credential azcore.TokenCredential) (*azblob.Client, error) {
	o, err := NewOptions()
	if err != nil {
		return nil, err
	}

	return NewClientWithOptions(ctx, credential, o)
}

// NewClientWithConfigPath NewConfigWithOptions returns azure client with options from config path.
func NewClientWithConfigPath(ctx context.Context, credential azcore.TokenCredential, path string) (*azblob.Client, error) {
	opts, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewClientWithOptions(ctx, credential, opts)
}

// NewClientWithOptions returns azure client with options.
func NewClientWithOptions(ctx context.Context, credential azcore.TokenCredential, options *Options) (*azblob.Client, error) {
	url := fmt.Sprintf("https://%s.blob.core.windows.net/", options.AccountName)
	return azblob.NewClient(url, credential, nil)
}
