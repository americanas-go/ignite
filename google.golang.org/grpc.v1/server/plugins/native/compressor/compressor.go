package compressor

import (
	"context"

	"github.com/americanas-go/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
)

func Register(ctx context.Context) []grpc.ServerOption {
	o, err := NewOptions()
	if err != nil {
		return nil
	}
	h := NewCompressorWithOptions(o)
	return h.Register(ctx)
}

type Compressor struct {
	options *Options
}

func NewCompressorWithOptions(options *Options) *Compressor {
	return &Compressor{options: options}
}

func NewCompressorWithConfigPath(path string) (*Compressor, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewCompressorWithOptions(o), nil
}

func NewCompressor() *Compressor {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewCompressorWithOptions(o)
}

func (i *Compressor) Register(ctx context.Context) []grpc.ServerOption {

	logger := log.FromContext(ctx)
	logger.Debug("compressor successfully enabled in grpc server")

	err := gzip.SetLevel(i.options.Level)
	if err != nil {
		logger.Fatalf("could not set level: %s", err.Error())
	}

	return nil
}
