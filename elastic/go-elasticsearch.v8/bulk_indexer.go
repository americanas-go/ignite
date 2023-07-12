package elasticsearch

import (
	"context"
	"github.com/americanas-go/log"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
)

// NewBulkIndexer returns elasticsearch bulk indexer with default options.
func NewBulkIndexer(ctx context.Context, client *elasticsearch.Client) (esutil.BulkIndexer, error) {

	logger := log.FromContext(ctx)

	o, err := NewBulkOptions()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	return NewBulkIndexerWithOptions(ctx, o, client)
}

// NewBulkIndexerWithConfigPath returns elasticsearch bulk indexer with options from config path.
func NewBulkIndexerWithConfigPath(ctx context.Context, path string, client *elasticsearch.Client) (esutil.BulkIndexer, error) {
	opts, err := NewBulkOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewBulkIndexerWithOptions(ctx, opts, client)
}

// NewBulkIndexerWithOptions returns elasticsearch bulk indexer with options.
func NewBulkIndexerWithOptions(ctx context.Context, options *BulkOptions, client *elasticsearch.Client) (bi esutil.BulkIndexer, err error) {

	logger := log.FromContext(ctx)

	bi, err = esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		NumWorkers:    options.NumWorkers, // The number of worker goroutines
		FlushBytes:    options.FlushBytes,
		FlushInterval: options.FlushInterval, // The periodic flush interval
		Client:        client,                // The Elasticsearch client
		DebugLogger:   &DebugLogger{},
		OnError: func(ctx context.Context, err error) {
			logger.Errorf("ERROR: %s", err.Error())
		},
		Index:      options.Index, // The default index name
		ErrorTrace: true,
		Pipeline:   options.Pipeline,
		Timeout:    options.Timeout,
		//Decoder:       nil,
		//OnFlushStart: nil,
		//OnFlushEnd:   nil,
		//FilterPath: nil,
		//Header:     nil,
		//Human:      false,
		//Pretty:     false,
		//Refresh:    "",
		//Routing:    "",
		//Source:              nil,
		//SourceExcludes:      nil,
		//SourceIncludes:      nil,
		//WaitForActiveShards: "",
	})

	logger.Infof("Created Elastic Bulk Indexer")

	return bi, err
}
