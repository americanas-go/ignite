package bigquery

import (
	"context"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/option"
)

func NewClient(ctx context.Context) (*bigquery.Client, error) {
	opt, err := NewOptions()
	if err != nil {
		return nil, err
	}
	return NewClientWithOptions(ctx, opt)
}

func NewClientWithConfigPath(ctx context.Context, path string) (*bigquery.Client, error) {
	options, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewClientWithOptions(ctx, options)
}

func NewClientWithOptions(ctx context.Context, options *Options) (*bigquery.Client, error) {

	var opts []option.ClientOption

	if options.Credentials.JSON != nil {
		opts = append(opts, option.WithCredentialsJSON(options.Credentials.JSON))
	} else {
		opts = append(opts, option.WithCredentialsFile(options.Credentials.File))
	}

	return bigquery.NewClient(ctx, options.ProjectID)
}
