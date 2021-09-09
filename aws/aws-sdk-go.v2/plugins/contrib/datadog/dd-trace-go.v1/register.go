package datadog

import (
	"context"

	datadog "github.com/americanas-go/ignite/datadog/dd-trace-go.v1"
	"github.com/americanas-go/log"
	"github.com/aws/aws-sdk-go-v2/aws"
	awstrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/aws/aws-sdk-go-v2/aws"
)

// Register registers AWS tracing for Datadog.
func Register(ctx context.Context, awsCfg *aws.Config) error {

	if !IsEnabled() || !datadog.IsTracerEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling datadog middleware in aws")

	awstrace.AppendMiddleware(awsCfg)

	logger.Debug("datadog middleware successfully enabled in aws")

	return nil

}
