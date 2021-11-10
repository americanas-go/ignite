package ftp

import (
	"context"
	"time"

	"github.com/americanas-go/log"
	"github.com/jlaffaye/ftp"
	"gopkg.in/matryer/try.v1"
)

func NewServerConnWithOptions(ctx context.Context, options *Options) (*ftp.ServerConn, error) {

	var conn *ftp.ServerConn

	err := try.Do(func(attempt int) (bool, error) {
		var err error
		conn, err = ftp.Dial(options.Addr, ftp.DialWithTimeout(time.Duration(options.Timeout)*time.Second))
		return attempt < options.Retry, err
	})
	if err != nil {
		return nil, err
	}

	err = conn.Login(options.User, options.Password)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func NewServerConnWithConfigPath(ctx context.Context, path string) (*ftp.ServerConn, error) {
	options, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewServerConnWithOptions(ctx, options)
}

func NewServerConn(ctx context.Context) (*ftp.ServerConn, error) {

	logger := log.FromContext(ctx)

	o, err := NewOptions()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	return NewServerConnWithOptions(ctx, o)
}
