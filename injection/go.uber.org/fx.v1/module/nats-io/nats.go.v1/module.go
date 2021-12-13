package nats

import (
	"sync"

	contextfx "github.com/americanas-go/ignite/injection/go.uber.org/fx.v1/module/context"
	"github.com/americanas-go/ignite/messaging/nats-io/nats.go.v1"
	"go.uber.org/fx"
)

var subsOnce sync.Once

// Module fx module for nats subscriber.
func SubscriberModule() fx.Option {
	options := fx.Options()

	subsOnce.Do(func() {
		options = fx.Options(
			contextfx.Module(),
			fx.Provide(
				nats.NewSubscriber,
			),
		)
	})

	return options
}

var pubOnce sync.Once

// Module fx module for nats publisher.
func PublisherModule() fx.Option {
	options := fx.Options()

	pubOnce.Do(func() {
		options = fx.Options(
			contextfx.Module(),
			fx.Provide(
				nats.NewPublisher,
			),
		)
	})

	return options
}
