package sqs

import (
	"context"

	"github.com/americanas-go/errors"
	"github.com/americanas-go/log"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go/aws"
)

// Client knows how to publish on sqs
type Client interface {
	Publish(ctx context.Context, input *sqs.SendMessageInput) error
	ResolveQueueUrl(ctx context.Context, queueName string) (*string, error)
}

type sqsClient interface {
	SendMessage(ctx context.Context, params *sqs.SendMessageInput, optFns ...func(*sqs.Options)) (*sqs.SendMessageOutput, error)
	GetQueueUrl(ctx context.Context, params *sqs.GetQueueUrlInput, optFns ...func(*sqs.Options)) (*sqs.GetQueueUrlOutput, error)
}

// Client holds client and resource name
type client struct {
	client    sqsClient
	queueUrls map[string]*string
}

// NewClient returns a initialized client
func NewClient(c *sqs.Client) Client {
	return &client{c, map[string]*string{}}
}

// Publish publish message on sns
func (c *client) Publish(ctx context.Context, input *sqs.SendMessageInput) error {

	logger := log.FromContext(ctx).
		WithTypeOf(*c).
		WithField("subject", input.QueueUrl)

	logger.Tracef("sending message to sqs")

	response, err := c.client.SendMessage(ctx, input)
	if err != nil {
		return errors.Wrap(err, errors.New("error sending message to sqs"))
	}

	logger.
		WithField("message_id", *response.MessageId).
		Debug("message sent to sqs")

	return nil
}

func (c *client) ResolveQueueUrl(ctx context.Context, queueName string) (*string, error) {
	if queueUrl, ok := c.queueUrls[queueName]; ok {
		return queueUrl, nil
	}

	result, err := c.client.GetQueueUrl(ctx, &sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	})

	if err != nil {
		return nil, errors.Wrap(err, errors.New("error resolving sqs queue url"))
	}

	if result == nil || result.QueueUrl == nil {
		return nil, errors.Errorf("sqs queue %s not found", queueName)
	}

	c.queueUrls[queueName] = result.QueueUrl

	return result.QueueUrl, nil
}
