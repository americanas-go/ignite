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

type Plugin func(context.Context, *Server) error

type Server struct {
	instance *echo.Echo
	options  *Options
}

func NewServer(ctx context.Context, plugins ...Plugin) *Server {
	opt, err := NewOptions()
	if err != nil {
		panic(err)
	}
	return NewServerWithOptions(ctx, opt, plugins...)
}

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

func (s *Server) Instance() *echo.Echo {
	return s.instance
}

func (s *Server) Options() *Options {
	return s.options
}
func (s *Server) GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.instance.GET(path, h, m...)
}

func (s *Server) POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.instance.POST(path, h, m...)
}

func (s *Server) PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.instance.PUT(path, h, m...)
}

func (s *Server) DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.instance.DELETE(path, h, m...)
}

func (s *Server) OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.instance.OPTIONS(path, h, m...)
}

func (s *Server) PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.instance.PATCH(path, h, m...)
}

func (s *Server) HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.instance.HEAD(path, h, m...)
}

func (s *Server) CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.instance.CONNECT(path, h, m...)
}

func (s *Server) Group(prefix string, m ...echo.MiddlewareFunc) *echo.Group {
	return s.instance.Group(prefix, m...)
}

func (s *Server) Add(method string, path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.instance.Add(method, path, h, m...)
}

func (s *Server) Any(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) []*echo.Route {
	return s.instance.Any(path, h, m...)
}

func (s *Server) Use(middleware ...echo.MiddlewareFunc) {
	s.instance.Use(middleware...)
}

func (s *Server) Static(prefix, root string) *echo.Route {
	return s.instance.Static(prefix, root)
}

func (s *Server) Match(methods []string, path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) []*echo.Route {
	return s.instance.Match(methods, path, handler, middleware...)
}

func (s *Server) Pre(middleware ...echo.MiddlewareFunc) {
	s.instance.Pre(middleware...)
}

func (s *Server) File(path, file string, m ...echo.MiddlewareFunc) *echo.Route {
	return s.instance.File(path, file, m...)
}

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
