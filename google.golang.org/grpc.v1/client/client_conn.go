package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"strconv"

	"github.com/americanas-go/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
)

type Plugin func(ctx context.Context) ([]grpc.DialOption, []grpc.CallOption)

func NewClientConnWithOptions(ctx context.Context, options *Options, plugins ...Plugin) *grpc.ClientConn {

	var err error
	var conn *grpc.ClientConn
	var opts []grpc.DialOption

	logger := log.FromContext(ctx)

	serverAddr := options.Host + ":" + strconv.Itoa(options.Port)

	if options.TLS.Enabled {
		logger.Tracef("creating TLS grpc client for host %s", serverAddr)
		opts = addTlsOptions(ctx, options, opts)
	} else {
		logger.Tracef("creating insecure grpc client for host %s", serverAddr)
		opts = append(opts, grpc.WithInsecure())
	}

	callOpts := make([]grpc.CallOption, 0)

	opts = append(opts, grpc.WithInitialWindowSize(options.InitialWindowSize))
	opts = append(opts, grpc.WithInitialConnWindowSize(options.InitialConnWindowSize))

	if options.Block {
		opts = append(opts, grpc.WithBlock())
	}

	if options.HostOverwrite != "" {
		opts = append(opts, grpc.WithAuthority(options.HostOverwrite))
	}

	opts = append(opts, grpc.WithConnectParams(grpc.ConnectParams{
		Backoff: backoff.Config{
			BaseDelay:  options.ConnectParams.Backoff.BaseDelay,
			Multiplier: options.ConnectParams.Backoff.Multiplier,
			Jitter:     options.ConnectParams.Backoff.Jitter,
			MaxDelay:   options.ConnectParams.Backoff.MaxDelay,
		},
		MinConnectTimeout: options.ConnectParams.MinConnectTimeout,
	}))

	opts = append(opts, grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:                options.Keepalive.Time,
		Timeout:             options.Keepalive.Timeout,
		PermitWithoutStream: options.Keepalive.PermitWithoutStream,
	}))

	for _, plugin := range plugins {
		dopts, copts := plugin(ctx)
		if dopts != nil {
			opts = append(opts, dopts...)
		}
		if copts != nil {
			callOpts = append(callOpts, copts...)
		}
	}

	if len(callOpts) > 0 {
		opts = append(opts, grpc.WithDefaultCallOptions(callOpts...))
	}

	conn, err = grpc.Dial(serverAddr, opts...)

	if err != nil {
		logger.Fatalf("fail to dial: %v", err)
		return nil
	}

	logger.Debugf("grpc client created for host %s", serverAddr)

	return conn
}

func addTlsOptions(ctx context.Context, opt *Options, opts []grpc.DialOption) []grpc.DialOption {

	logger := log.FromContext(ctx)

	var creds credentials.TransportCredentials

	options := opt.TLS

	if options.CertFile != "" && options.KeyFile != "" {

		// Load the client certificates from disk
		cert, err := tls.LoadX509KeyPair(options.CertFile, options.KeyFile)
		if err != nil {
			logger.Fatalf("could not load client key pair: %s", err)
		}

		if options.CAFile != "" {

			// Create a certificate pool from the certificate authority
			certPool := x509.NewCertPool()
			ca, err := ioutil.ReadFile(options.CAFile)
			if err != nil {
				logger.Fatalf("could not read ca certificate: %s", err)
			}

			// Append the certificates from the CA
			if ok := certPool.AppendCertsFromPEM(ca); !ok {
				logger.Fatalf("failed to append ca certs")
			}

			creds = credentials.NewTLS(&tls.Config{
				ServerName:         opt.Host, // NOTE: this is required!
				Certificates:       []tls.Certificate{cert},
				RootCAs:            certPool,
				InsecureSkipVerify: options.InsecureSkipVerify,
			})

		} else {

			creds = credentials.NewTLS(&tls.Config{
				ServerName:         opt.Host, // NOTE: this is required!
				Certificates:       []tls.Certificate{cert},
				InsecureSkipVerify: options.InsecureSkipVerify,
			})

		}

	} else {

		creds = credentials.NewTLS(&tls.Config{
			ServerName:         opt.Host, // NOTE: this is required!
			Certificates:       []tls.Certificate{},
			InsecureSkipVerify: options.InsecureSkipVerify,
		})

	}

	return append(opts, grpc.WithTransportCredentials(creds))
}
