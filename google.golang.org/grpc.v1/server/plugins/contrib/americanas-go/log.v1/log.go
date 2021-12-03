package log

import (
	"context"
	"time"

	"github.com/americanas-go/log"
	"google.golang.org/grpc"
)

// Register registers a new logger plugin for grpc server.
func Register(ctx context.Context) []grpc.ServerOption {
	o, err := NewOptions()
	if err != nil {
		return nil
	}
	h := NewLogWithOptions(o)
	return h.Register(ctx)
}

// Log represents logger plugin for grpc server.
type Log struct {
	options *Options
}

// NewLogWithOptions returns a new logger plugin with options.
func NewLogWithOptions(options *Options) *Log {
	return &Log{options: options}
}

// NewLogWithConfigPath returns a new logger plugin with options from config path.
func NewLogWithConfigPath(path string) (*Log, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewLogWithOptions(o), nil
}

// NewLog returns a new logger plugin with default options.
func NewLog() *Log {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewLogWithOptions(o)
}

// Register registers this logger plugin for grpc server.
func (i *Log) Register(ctx context.Context) []grpc.ServerOption {

	if !i.options.Enabled {
		return nil
	}

	logger := log.FromContext(ctx)
	logger.Debug("logger interceptor successfully enabled in grpc server")

	return []grpc.ServerOption{
		grpc.ChainStreamInterceptor(i.streamInterceptor()),
		grpc.ChainUnaryInterceptor(i.unaryInterceptor()),
	}
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

func (i *Log) streamInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		logger := log.FromContext(stream.Context())

		start := time.Now()
		wrapper := &recvWrapper{stream}
		err := handler(srv, wrapper)

		logger = logger.WithFields(log.Fields{
			"method": info.FullMethod,

			"duration": time.Since(start),
		})

		if err != nil {
			logger = logger.WithField("error", err.Error())
		}

		xx := i.m(logger)
		xx("stream request received")
		return err
	}
}

func (i *Log) unaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

		logger := log.FromContext(ctx)

		start := time.Now()
		resp, err = handler(ctx, req)

		logger = logger.WithFields(log.Fields{
			"method":   info.FullMethod,
			"duration": time.Since(start),
			"response": resp,
			"req":      req,
		})

		if err != nil {
			logger = logger.WithField("error", err.Error())
		}

		xx := i.m(logger)
		xx("unary request received")
		return resp, err
	}
}

type recvWrapper struct {
	grpc.ServerStream
}
