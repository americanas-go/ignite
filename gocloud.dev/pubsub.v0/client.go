package pubsub

import (
	"context"
	"strings"

	"github.com/americanas-go/log"
	"gocloud.dev/pubsub"
	_ "gocloud.dev/pubsub/awssnssqs"
	_ "gocloud.dev/pubsub/gcppubsub"
	_ "gocloud.dev/pubsub/kafkapubsub"
	_ "gocloud.dev/pubsub/mempubsub"
)

// Queue types
const (
	SQS    = "sqs"
	SNS    = "sns"
	KAFKA  = "kafka"
	NATS   = "nats"
	PUBSUB = "pubsub"
)

// NewTopic ..
func NewTopic(ctx context.Context) (*pubsub.Topic, error) {

	logger := log.FromContext(ctx)

	o, err := NewOptions()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	return NewTopicWithOptions(ctx, o)
}

// NewTopicWithOptions start a new topic for send message
func NewTopicWithOptions(ctx context.Context, o *Options) (*pubsub.Topic, error) {

	logger := log.FromContext(ctx)

	addResource(o)

	if o.Region != "" {
		o.Resource = appendRegion(o.Resource, o.Region)
	}

	topic, err := pubsub.OpenTopic(ctx, o.Resource)
	if err != nil {
		return nil, err
	}

	logger.Infof("open topic for send message")
	return topic, nil

}

// NewSubscription ..
func NewSubscription(ctx context.Context) (*pubsub.Subscription, error) {

	logger := log.FromContext(ctx)

	o, err := NewOptions()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	return NewSubscriptionWithOptions(ctx, o)
}

// NewSubscriptionWithOptions ..
func NewSubscriptionWithOptions(ctx context.Context, o *Options) (*pubsub.Subscription, error) {

	logger := log.FromContext(ctx)

	addResource(o)

	if o.Region != "" {
		o.Resource = appendRegion(o.Resource, o.Region)
	}

	subscription, err := pubsub.OpenSubscription(ctx, o.Resource)
	if err != nil {
		return nil, err
	}

	logger.Infof("open subscription for listen")
	return subscription, nil

}

func addResource(o *Options) {
	switch strings.ToLower(o.Type) {
	case SQS:
		o.Resource = addSQSResource(o.Resource)
	case SNS:
		o.Resource = addSNSResource(o.Resource)
	case KAFKA:
		// TODO: https://gocloud.dev/howto/pubsub/publish/#kafka
	case NATS:
		// TODO: https://gocloud.dev/howto/pubsub/publish/#nats
	case PUBSUB:
		o.Resource = addPUBSUBResource(o.Resource)
	default:
		o.Resource = addMEMResource(o.Resource)
	}
}

func addMEMResource(topicName string) string {
	return "mem://" + topicName
}

func addSQSResource(url string) string {
	newURL := strings.Replace(url, "https://", "awssqs://", -1)
	return newURL
}

func addSNSResource(arn string) string {
	return "awssns:///" + arn
}

func addPUBSUBResource(url string) string {
	return "gcppubsub://" + url
}

func appendRegion(add, region string) string {
	return add + "?region=" + region
}
