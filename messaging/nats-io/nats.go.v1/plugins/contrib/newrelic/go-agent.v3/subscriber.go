package newrelic

import (
	newrelic "github.com/americanas-go/ignite/newrelic/go-agent.v3"
	"github.com/americanas-go/log"
	"github.com/nats-io/nats.go"
	"github.com/newrelic/go-agent/v3/integrations/nrnats"
)

// Register registers a new newrelic plugin for nats subscriber.
func Register(msgHandler nats.MsgHandler) nats.MsgHandler {
	o, err := NewOptions()
	if err != nil {
		return nil
	}
	h := NewSubscriberWithOptions(o)
	return h.Register(msgHandler)
}

// NewSubscriber returns a new newrelic plugin with default options.
func NewSubscriber() *Subscriber {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewSubscriberWithOptions(o)
}

// NewSubscriberWithConfigPath returns a new newrelic plugin with options from config path.
func NewSubscriberWithConfigPath(path string) (*Subscriber, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewSubscriberWithOptions(o), nil
}

// NewSubscriberWithOptions returns a new newrelic plugin with options.
func NewSubscriberWithOptions(options *Options) *Subscriber {
	return &Subscriber{options: options}
}

// Subscriber represents newrelic plugin for nats subscriber.
type Subscriber struct {
	options *Options
}

// Register registers this newrelic plugin for nats subscriber.
func (p *Subscriber) Register(msgHandler nats.MsgHandler) nats.MsgHandler {
	if !p.options.Enabled || !newrelic.IsEnabled() {
		return msgHandler
	}

	return nrnats.SubWrapper(newrelic.Application(), msgHandler)
}
