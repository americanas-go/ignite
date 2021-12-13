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

// NewTopic creates a topic publisher
func NewTopic(ctx context.Context) (*pubsub.Topic, error) {
	o, err := NewOptions()
	if err != nil {
		return nil, err
	}

	return NewTopicWithOptions(ctx, o)
}

// NewTopicWithConfigPath start a new topic for sending messages with options from config path.
func NewTopicWithConfigPath(ctx context.Context, path string) (*pubsub.Topic, error) {
	options, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewTopicWithOptions(ctx, options)
}

// NewTopicWithOptions start a new topic for sending messages.
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

// NewSubscription creates a subscription.
func NewSubscription(ctx context.Context) (*pubsub.Subscription, error) {
	o, err := NewOptions()
	if err != nil {
		return nil, err
	}

	return NewSubscriptionWithOptions(ctx, o)
}

// NewSubscriptionWithConfigPath return a subscription with options from config path.
func NewSubscriptionWithConfigPath(ctx context.Context, path string) (*pubsub.Subscription, error) {
	options, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewSubscriptionWithOptions(ctx, options)
}

// NewSubscriptionWithOptions return a subscription with options from config path.
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
