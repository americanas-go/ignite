package kafka

import (
	"context"

	"github.com/americanas-go/log"
	"github.com/segmentio/kafka-go"
)

func NewConnWithConfigPath(ctx context.Context, path string) (*kafka.Conn, error) {
	options, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewConnWithOptions(ctx, options)
}

func NewConnWithOptions(ctx context.Context, o *Options) (conn *kafka.Conn, err error) {

	logger := log.FromContext(ctx)

	switch o.ConnType {
	case "SERVER":
		conn, err = kafka.DialContext(context.Background(), o.Network, o.Address)
	case "PARTITION":
		conn, err = kafka.DialPartition(context.Background(), o.Network, o.Address, kafka.Partition{
			Topic: o.Topic,
			ID:    o.Partition,
		})
	default:
		conn, err = kafka.DialLeader(context.Background(), o.Network, o.Address, o.Topic, o.Partition)
	}
	if err != nil {
		logger.Fatal("failed to dial %s. %s", o.ConnType, err.Error())
	}

	logger.Infof("Created kafka connection to %v", o.Address)

	return conn, err
}

func NewConn(ctx context.Context) (*kafka.Conn, error) {

	logger := log.FromContext(ctx)

	o, err := NewOptions()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	return NewConnWithOptions(ctx, o)
}
