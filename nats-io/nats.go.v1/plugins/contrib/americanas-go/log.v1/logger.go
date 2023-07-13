package log

import (
	"context"
	"github.com/americanas-go/log"
	"github.com/nats-io/nats.go"
)

// Register registers a new logger plugin for nats connection.
func Register(ctx context.Context) (func(context.Context, *nats.Conn) error, func(context.Context, []nats.Option) ([]nats.Option, error)) {
	return nil, func(ctx context.Context, options []nats.Option) ([]nats.Option, error) {
		options = append(options,
			nats.DisconnectErrHandler(disconnectedErrHandler),
			nats.ReconnectHandler(reconnectedHandler),
			nats.ClosedHandler(closedHandler))
		return options, nil
	}
}

func disconnectedErrHandler(nc *nats.Conn, err error) {
	log.Warnf("Disconnected from nats server! will attempt reconnects")
}

func reconnectedHandler(nc *nats.Conn) {
	log.Warnf("Reconnected [%s]", nc.ConnectedUrl())
}

func closedHandler(nc *nats.Conn) {
	log.Errorf("Exiting: %v", nc.LastError())
}
