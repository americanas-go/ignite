package main

import (
	"context"

	"github.com/americanas-go/config"
	"github.com/americanas-go/log"
	a "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/jvitoroc/ignite/aws/aws-sdk-go.v2"
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

	// create s3 client

	s3Client := s3.NewFromConfig(awsConfig)

	// set vars
	filename := "examplefile"
	bucket := config.String(Bucket)

	// prepare s3 request head
	input := &s3.HeadObjectInput{
		Bucket: a.String(bucket),
		Key:    a.String(filename),
	}

	// make a call
	head, err := s3Client.HeadObject(ctx, input)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	if err != nil {

		logger.Fatalf("unable check file %s in s3 bucket %s", filename, bucket)
	}

	logger = logger.WithFields(
		log.Fields{"lastModified": head.LastModified,
			"versionId": head.VersionId,
		})

	logger.Debugf("file %s exists on bucket %s", filename, bucket)

}
