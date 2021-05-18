package log

import (
	"time"

	"github.com/americanas-go/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type l func(format string, args ...interface{})

var lvl string

func m(logger log.Logger) (method l) {

	if lvl == "" {
		lvl = Level()
	}
	switch lvl {
	case "TRACE":
		method = logger.Tracef
	case "DEBUG":
		method = logger.Debugf
	default:
		method = logger.Infof
	}

	return method
}

func streamInterceptor() grpc.StreamClientInterceptor {
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

		x := m(logger)

		x("call stream")

		return clientStream, err
	}
}

func unaryInterceptor() grpc.UnaryClientInterceptor {
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

		x := m(logger)

		x("call unary")

		return err
	}
}
