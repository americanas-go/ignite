package nats

import (
	"context"

	"github.com/americanas-go/log"
	"github.com/nats-io/nats.go"
)

// Plugin defines a function to process plugin.
type Plugin func(context.Context) (func(context.Context, *nats.Conn) error, func(context.Context, []nats.Option) ([]nats.Option, error))

// NewConnWithOptions registers a nats connection.
func NewConnWithOptions(ctx context.Context, options *Options, plugins ...Plugin) (*nats.Conn, error) {

	logger := log.FromContext(ctx)

	opts := []nats.Option{
		nats.MaxReconnects(options.MaxReconnects),
		nats.ReconnectWait(options.ReconnectWait),
		/*
			nats.(options.RequestChanLen),
			nats.(options.DrainTimeout),
			nats.(options.MaxChanLen),
			nats.(options.MaxPingOut),
			nats.(options.PingInterval),
			nats.(options.ReconnectBufSize),
			nats.(options.ReconnectJitter),
			nats.(options.ReconnectJitterTLS),
			nats.(options.Timeout),
		*/
	}

	for _, plugin := range plugins {
		if _, cfgPlugin := plugin(ctx); cfgPlugin != nil {
			var err error
			opts, err = cfgPlugin(ctx, opts)
			if err != nil {
				panic(err)
			}
		}
	}

	conn, err := nats.Connect(options.Url, opts...)

	if err != nil {
		return nil, err
	}

	for _, plugin := range plugins {
		if connPlugin, _ := plugin(ctx); connPlugin != nil {
			if err := connPlugin(ctx, conn); err != nil {
				panic(err)
			}
		}
	}

	logger.Infof("Connected to NATS server: %s", options.Url)

	return conn, nil
}

// NewConnWithConfigPath returns a new nats connection with options from config path.
func NewConnWithConfigPath(ctx context.Context, path string) (*nats.Conn, error) {
	options, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewConnWithOptions(ctx, options)
}

// NewConn returns a new connection with default options.
func NewConn(ctx context.Context, plugins ...Plugin) (*nats.Conn, error) {

	logger := log.FromContext(ctx)

	o, err := NewOptions()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	return NewConnWithOptions(ctx, o, plugins...)
}
