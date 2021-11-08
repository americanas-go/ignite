package log

import (
	"context"
	"time"

	"github.com/americanas-go/log"
	"google.golang.org/grpc"
)

func Register(ctx context.Context) ([]grpc.DialOption, []grpc.CallOption) {
	o, err := NewOptions()
	if err != nil {
		return nil, nil
	}
	h := NewLogWithOptions(o)
	return h.Register(ctx)
}

type Log struct {
	options *Options
}

func NewLogWithOptions(options *Options) *Log {
	return &Log{options: options}
}

func NewLogWithConfigPath(path string) (*Log, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewLogWithOptions(o), nil
}

func NewLog() *Log {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewLogWithOptions(o)
}

func (i *Log) Register(ctx context.Context) ([]grpc.DialOption, []grpc.CallOption) {

	if !i.options.Enabled {
		return nil, nil
	}

	logger := log.FromContext(ctx)
	logger.Debug("logger interceptor successfully enabled in grpc client")

	return []grpc.DialOption{
		grpc.WithChainStreamInterceptor(i.streamInterceptor()),
		grpc.WithChainUnaryInterceptor(i.unaryInterceptor()),
	}, nil
}

type l func(format string, args ...interface{})

func (i *Log) m(logger log.Logger) (method l) {

	switch i.options.Level {
	case "TRACE":
		method = logger.Tracef
	case "DEBUG":
		method = logger.Debugf
	default:
		method = logger.Infof
	}

	return method
}

func (i *Log) streamInterceptor() grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {

		logger := log.FromContext(ctx)

		start := time.Now()
		clientStream, err := streamer(ctx, desc, cc, method, opts...)

		logger = logger.WithFields(log.Fields{
			"method":   method,
			"duration": time.Since(start),
		})

		if err != nil {
			logger = logger.WithField("error", err.Error())
		}

		x := i.m(logger)

		x("call stream")

		return clientStream, err
	}
}

func (i *Log) unaryInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

		logger := log.FromContext(ctx)

		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)

		logger = logger.WithFields(log.Fields{
			"method":   method,
			"duration": time.Since(start),
			"request":  req,
			"response": reply,
		})

		if err != nil {
			logger = logger.WithField("error", err.Error())
		}

		x := i.m(logger)

		x("call unary")

		return err
	}
}
