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

func streamInterceptor() grpc.StreamServerInterceptor {
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

		xx := m(logger)
		xx("stream request received")
		return err
	}
}

func unaryInterceptor() grpc.UnaryServerInterceptor {
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

		xx := m(logger)
		xx("unary request received")
		return resp, err
	}
}

type recvWrapper struct {
	grpc.ServerStream
}
