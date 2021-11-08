package nats

import (
	"context"

	"github.com/americanas-go/log"
	"github.com/nats-io/nats.go"
)

type Plugin func(context.Context, *nats.Conn) error

func NewConnWithOptions(ctx context.Context, options *Options, plugins ...Plugin) (*nats.Conn, error) {

	logger := log.FromContext(ctx)

	conn, err := nats.Connect(
		options.Url,
		nats.MaxReconnects(options.MaxReconnects),
		nats.ReconnectWait(options.ReconnectWait),
		nats.DisconnectErrHandler(disconnectedErrHandler),
		nats.ReconnectHandler(reconnectedHandler),
		nats.ClosedHandler(closedHandler),
	)

	if err != nil {
		return nil, err
	}

	for _, plugin := range plugins {
		if err := plugin(ctx, conn); err != nil {
			panic(err)
		}
	}

	logger.Infof("Connected to NATS server: %s", options.Url)

	return conn, nil
}

func NewConnWithConfigPath(ctx context.Context, path string) (*nats.Conn, error) {
	options, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewConnWithOptions(ctx, options)
}

func NewConn(ctx context.Context, plugins ...Plugin) (*nats.Conn, error) {

	logger := log.FromContext(ctx)

	o, err := NewOptions()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	return NewConnWithOptions(ctx, o, plugins...)
}

func disconnectedErrHandler(nc *nats.Conn, err error) {
	log.Error("Disconnected from nats server! will attempt reconnects")
}

func reconnectedHandler(nc *nats.Conn) {
	log.Warnf("Reconnected [%s]", nc.ConnectedUrl())
}

func closedHandler(nc *nats.Conn) {
	log.Errorf("Exiting: %v", nc.LastError())
}
