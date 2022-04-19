package client

import (
	"context"
	"net"
	"strconv"

	"github.com/americanas-go/log"
	"storj.io/drpc/drpcconn"
)

// NewClientConn returns a new grpc client connection.
func NewClientConn(ctx context.Context) (*drpcconn.Conn, error) {
	opt, err := NewOptions()
	if err != nil {
		return nil, err
	}
	return NewClientConnWithOptions(ctx, opt)
}

// NewClientConnWithConfigPath returns a new grpc client connection with options from config path.
func NewClientConnWithConfigPath(ctx context.Context, path string) (*drpcconn.Conn, error) {
	options, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewClientConnWithOptions(ctx, options)
}

// NewClientConnWithOptions returns a new grpc client connection with options.
func NewClientConnWithOptions(ctx context.Context, options *Options) (*drpcconn.Conn, error) {

	var err error
	var rawconn net.Conn

	logger := log.FromContext(ctx)

	serverAddr := options.Host + ":" + strconv.Itoa(options.Port)

	rawconn, err = net.Dial("tcp", serverAddr)
	if err != nil {
		logger.Errorf("fail to dial: %v", err)
		return nil, err
	}

	conn := drpcconn.New(rawconn)
	// defer conn.Close()

	logger.Debugf("drpc client created for host %s", serverAddr)

	return conn, nil
}
