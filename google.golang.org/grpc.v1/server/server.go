package server

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
	"net"
	"time"

	"github.com/americanas-go/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/channelz/service"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

type Plugin func(ctx context.Context) []grpc.ServerOption

type Server struct {
	server           *grpc.Server
	serviceRegistrar grpc.ServiceRegistrar
	options          *Options
}

func NewServer(ctx context.Context, plugins ...Plugin) *Server {
	opt, err := NewOptions()
	if err != nil {
		panic(err)
	}
	return NewServerWithOptions(ctx, opt, plugins...)
}

func NewServerWithConfigPath(ctx context.Context, path string) (*Server, error) {
	options, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewServerWithOptions(ctx, options), nil
}

func NewServerWithOptions(ctx context.Context, opt *Options, plugins ...Plugin) *Server {

	logger := log.FromContext(ctx)

	var s *grpc.Server

	var serverOptions []grpc.ServerOption

	if opt.TLS.Enabled {

		logger.Debug("configuring tls on grpc server")

		var creds credentials.TransportCredentials
		certPool := x509.NewCertPool()

		if opt.TLS.Type == "FILE" && opt.TLS.File.Cert != "" && opt.TLS.File.CA != "" {

			creds = tlsFromFile(ctx, opt, certPool)

		} else if opt.TLS.Type == "AUTO" {

			creds = autoTLS(ctx, opt, certPool)

		} else {

			creds = credentials.NewTLS(&tls.Config{
				ClientAuth:   tls.NoClientCert,
				Certificates: []tls.Certificate{},
				ClientCAs:    certPool,
			})

		}

		serverOptions = append(serverOptions, grpc.Creds(creds))
	}

	for _, plugin := range plugins {
		sopts := plugin(ctx)
		if sopts != nil {
			serverOptions = append(serverOptions, sopts...)
		}
	}

	serverOptions = append(serverOptions, grpc.MaxConcurrentStreams(uint32(opt.MaxConcurrentStreams)))
	serverOptions = append(serverOptions, grpc.InitialConnWindowSize(opt.InitialConnWindowSize))
	serverOptions = append(serverOptions, grpc.InitialWindowSize(opt.InitialWindowSize))

	s = grpc.NewServer(serverOptions...)

	return &Server{
		server:  s,
		options: opt,
	}
}

func (s *Server) Server() *grpc.Server {
	return s.server
}

func (s *Server) ServiceRegistrar() grpc.ServiceRegistrar {
	return s.server
}

func (s *Server) Serve(ctx context.Context) {

	logger := log.FromContext(ctx)

	service.RegisterChannelzServiceToServer(s.server)

	// Register reflection service on gRPC server.
	reflection.Register(s.server)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.options.Port))
	if err != nil {
		logger.Fatalf("failed to listen: %v", err.Error())
	}

	logger.Infof("grpc server started on port %v", s.options.Port)

	logger.Error(s.server.Serve(lis))
}

func (s *Server) Shutdown(ctx context.Context) {
	s.server.GracefulStop()
}

func autoTLS(ctx context.Context, options *Options, certPool *x509.CertPool) credentials.TransportCredentials {
	logger := log.FromContext(ctx)
	logger.Trace("configuring generated cert and key certificates on grpc server")

	var cert, key []byte
	var err error

	cert, key, err = generateCertificate(options.TLS.Auto.Host)
	if err != nil {
		logger.Fatal(err.Error())
	}

	// Load the certificates from disk
	var certificate tls.Certificate
	certificate, err = tls.X509KeyPair(cert, key)
	if err != nil {
		logger.Fatalf("could not load server key pair: %s", err.Error())
	}

	logger.Trace("cert and key certificates loaded")

	// Create the TLS credentials
	return credentials.NewTLS(&tls.Config{
		ClientAuth:   tls.NoClientCert,
		Certificates: []tls.Certificate{certificate},
		ClientCAs:    certPool,
	})
}

func tlsFromFile(ctx context.Context, options *Options, certPool *x509.CertPool) credentials.TransportCredentials {
	logger := log.FromContext(ctx)
	logger.Trace("configuring cert and key certificates from files on grpc server")

	// Load the certificates from disk
	certificate, err := tls.LoadX509KeyPair(options.TLS.File.Cert, options.TLS.File.Key)
	if err != nil {
		logger.Fatalf("could not load server key pair: %s", err.Error())
	}

	logger.Trace("cert and key certificates loaded")

	if options.TLS.File.CA != "" {

		logger.Trace("configuring ca certificate on grpc server")

		ca, err := ioutil.ReadFile(options.TLS.File.CA)
		if err != nil {
			logger.Fatalf("could not read ca certificate: %s", err.Error())
		}

		// Append the client certificates from the CA
		if ok := certPool.AppendCertsFromPEM(ca); !ok {
			logger.Fatalf("failed to append client certs")
		}

		logger.Trace("ca certificate loaded")

	}

	// Create the TLS credentials
	return credentials.NewTLS(&tls.Config{
		ClientAuth:   tls.NoClientCert,
		Certificates: []tls.Certificate{certificate},
		ClientCAs:    certPool,
	})
}

// generateCertificate generates a test certificate and private key based on the given host.
func generateCertificate(host string) ([]byte, []byte, error) {

	log.Tracef("generating a certificate and private key based on the given host %s", host)

	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, nil, err
	}

	cert := &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"grpc http server"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		SignatureAlgorithm:    x509.SHA256WithRSA,
		DNSNames:              []string{host},
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	certBytes, err := x509.CreateCertificate(
		rand.Reader, cert, cert, &priv.PublicKey, priv,
	)

	p := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(priv),
		},
	)

	b := pem.EncodeToMemory(
		&pem.Block{
			Type:  "CERTIFICATE",
			Bytes: certBytes,
		},
	)

	return b, p, err
}
