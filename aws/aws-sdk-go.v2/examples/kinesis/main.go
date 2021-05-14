package main

import (
	"context"

	"github.com/americanas-go/config"
	ilog "github.com/americanas-go/ignite/americanas-go/log.v1"
	"github.com/americanas-go/ignite/aws/aws-sdk-go.v2"
	"github.com/americanas-go/ignite/aws/aws-sdk-go.v2/client/kinesis"
	"github.com/americanas-go/log"
	akinesis "github.com/aws/aws-sdk-go-v2/service/kinesis"
)

const Bucket = "aws.s3.bucket"

func init() {
	config.Add(Bucket, "example", "s3 example bucket")
}

func main() {

	config.Load()

	// create background context
	ctx := context.Background()

	ilog.New()

	// get logrus instance from context
	logger := log.FromContext(ctx)

	// create default aws config
	awsConfig := aws.NewConfig(ctx)

	// create sns client
	sqsClient := akinesis.NewFromConfig(awsConfig)
	client := kinesis.NewClient(sqsClient)

	input := &akinesis.PutRecordInput{
		Data:                      nil,
		PartitionKey:              nil,
		StreamName:                nil,
		ExplicitHashKey:           nil,
		SequenceNumberForOrdering: nil,
	}

	// publish
	err := client.Publish(ctx, input)
	if err != nil {
		logger.Fatalf(err.Error())
	}

}
