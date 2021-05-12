package main

import (
	"context"

	"github.com/americanas-go/config"
	"github.com/americanas-go/log"
	asns "github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/jvitoroc/ignite/aws/aws-sdk-go.v2"
	"github.com/jvitoroc/ignite/aws/aws-sdk-go.v2/client/sns"
	"github.com/jvitoroc/ignite/sirupsen/logrus.v1"
)

const Bucket = "aws.s3.bucket"

func init() {
	config.Add(Bucket, "example", "s3 example bucket")
}

func main() {

	config.Load()

	// create background context
	ctx := context.Background()

	// start logrus
	// zap.NewLogger()
	logrus.NewLogger()

	// get logrus instance from context
	logger := log.FromContext(ctx)

	// create default aws config
	awsConfig := aws.NewConfig(ctx)

	// create sns client
	snsClient := asns.NewFromConfig(awsConfig)
	client := sns.NewClient(snsClient)

	input := &asns.PublishInput{
		Message:                nil,
		MessageAttributes:      nil,
		MessageDeduplicationId: nil,
		MessageGroupId:         nil,
		MessageStructure:       nil,
		PhoneNumber:            nil,
		Subject:                nil,
		TargetArn:              nil,
		TopicArn:               nil,
	}

	// publish
	err := client.Publish(ctx, input)
	if err != nil {
		logger.Fatalf(err.Error())
	}

}
