package echo

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"strconv"
	"time"

	"github.com/americanas-go/ignite/golang.org/x/net.v0/http2/server"
	"github.com/americanas-go/log"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/http2"
)

// Plugin defines an echo server Plugin function to execute.
type Plugin func(context.Context, *Server) error

// Server represents a echo server.
type Server struct {
	instance *echo.Echo
	options  *Options
}

// NewServer returns a new echo server with default options.
func NewServer(ctx context.Context, plugins ...Plugin) *Server {
	opt, err := NewOptions()
	if err != nil {
		panic(err)
	}
	return NewServerWithOptions(ctx, opt, plugins...)
}

// NewServerWithConfigPath returns a new echo server with options from config path.
func NewServerWithConfigPath(ctx context.Context, path string) (*Server, error) {
	options, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewServerWithOptions(ctx, options), nil
}

// NewServerWithOptions returns a new echo server with options.
func NewServerWithOptions(ctx context.Context, opt *Options, plugins ...Plugin) *Server {

	instance := echo.New()

	instance.HideBanner = opt.HideBanner
	instance.DisableHTTP2 = opt.DisableHTTP2
	instance.Logger = WrapLogger(log.GetLogger())

	srv := &Server{instance: instance, options: opt}

	for _, plugin := range plugins {
		if err := plugin(ctx, srv); err != nil {
			panic(err)
		}
	}

	return srv
}

// Instance returns the wrapped echo server instance.
func (s *Server) Instance() *echo.Echo {
	return s.instance
}

// Options returns echo server options.
func (s *Server) Options() *Options {
	return s.options
}

// GET registers the handler and middlewares for GET route at path.
func (s *Server) GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.instance.GET(path, h, m...)
}

// POST registers the handler and middlewares for POST route at path.
func (s *Server) POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.instance.POST(path, h, m...)
}

// PUT registers the handler and middlewares for PUT route at path.
func (s *Server) PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.instance.PUT(path, h, m...)
}

// DELETE registers the handler and middlewares for DELETE route at path.
func (s *Server) DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.instance.DELETE(path, h, m...)
}

// OPTIONS registers the handler and middlewares for OPTIONS route at path.
func (s *Server) OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.instance.OPTIONS(path, h, m...)
}

// PATCH registers the handler and middlewares for PATCH route at path.
func (s *Server) PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.instance.PATCH(path, h, m...)
}

// HEAD registers the handler and middlewares for HEAD route at path.
func (s *Server) HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.instance.HEAD(path, h, m...)
}

// CONNECT registers the handler and middlewares for CONNECT route at path.
func (s *Server) CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.instance.CONNECT(path, h, m...)
}

// Group creates a router group with prefix and registers middlewares.
func (s *Server) Group(prefix string, m ...echo.MiddlewareFunc) *echo.Group {
	return s.instance.Group(prefix, m...)
}

// Add registers the handler and middlewares for method route at path.
func (s *Server) Add(method string, path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.instance.Add(method, path, h, m...)
}

// Any registers the handler and middlewares for all method routes at path.
func (s *Server) Any(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) []*echo.Route {
	return s.instance.Any(path, h, m...)
}

// Use registers middlewares that will run after router.
func (s *Server) Use(middleware ...echo.MiddlewareFunc) {
	s.instance.Use(middleware...)
}

// Static registers a static path to server static files at root dir.
func (s *Server) Static(prefix, root string) *echo.Route {
	return s.instance.Static(prefix, root)
}

// Match registers the handler and middlewares for method routes at path.
func (s *Server) Match(methods []string, path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) []*echo.Route {
	return s.instance.Match(methods, path, handler, middleware...)
}

// Pre registers middlewares that will run before router.
func (s *Server) Pre(middleware ...echo.MiddlewareFunc) {
	s.instance.Pre(middleware...)
}

// File registers a path to a server file.
func (s *Server) File(path, file string, m ...echo.MiddlewareFunc) *echo.Route {
	return s.instance.File(path, file, m...)
}

// Serve starts echo server.
func (s *Server) Serve(ctx context.Context) {
	logger := log.FromContext(ctx)

	address := ":" + strconv.Itoa(s.options.Port)
	var err error

	if s.options.Protocol == "H2C" {

		logger.Infof("starting Echo H2C Server. https://echo.labstack.com/")

		var srv *http2.Server
		srv, err = server.NewServerWithPath(hc2Root)

		err = s.instance.StartH2CServer(address, srv)

	} else if s.options.TLS.Enabled {

		logger.Infof("starting Echo TLS Server. https://echo.labstack.com/")

		if s.options.TLS.Type == "AUTO" {
			// err = s.instance.StartAutoTLS(address)
			var cert, key []byte
			cert, key, err = generateCertificate(s.options.TLS.Auto.Host)
			if err == nil {
				err = s.instance.StartTLS(address, cert, key)
			}
		} else {
			err = s.instance.StartTLS(address, s.options.TLS.File.Cert, s.options.TLS.File.Key)
		}

	} else {

		logger.Infof("starting Echo Server. https://echo.labstack.com/")

		err = s.instance.Start(address)
	}

	if err != nil {
		logger.Error(err)
	}
}

// Shutdown stops echo server.
func (s *Server) Shutdown(ctx context.Context) {
	logger := log.FromContext(ctx)
	err := s.instance.Shutdown(ctx)
	if err != nil {
		logger.Warn(err)
	}
}

// generateCertificate generates a test certificate and private key based on the given host.
func generateCertificate(host string) ([]byte, []byte, error) {
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
			Organization: []string{"echo http server"},
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
