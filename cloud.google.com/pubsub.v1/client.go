package pubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

// NewClient returns a new bigquery client with default options.
func NewClient(ctx context.Context) (*pubsub.Client, error) {
	opt, err := NewOptions()
	if err != nil {
		return nil, err
	}
	return NewClientWithOptions(ctx, opt)
}

// NewClientWithConfigPath returns a new bigquery client with options from config path.
func NewClientWithConfigPath(ctx context.Context, path string) (*pubsub.Client, error) {
	options, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewClientWithOptions(ctx, options)
}

// NewClientWithOptions returns a new bigquery client with options.
func NewClientWithOptions(ctx context.Context, options *Options) (*pubsub.Client, error) {

	var opts []option.ClientOption

	if options.Credentials.JSON != "" {
		opts = append(opts, option.WithCredentialsJSON([]byte(options.Credentials.JSON)))
	} else {
		opts = append(opts, option.WithCredentialsFile(options.Credentials.File))
	}

	return pubsub.NewClient(ctx, options.ProjectID, opts...)
}
