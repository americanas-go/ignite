package server

import (
	"github.com/americanas-go/ignite"
	"golang.org/x/net/http2"
)

// NewServer returns a pointer with new Server
func NewServer() (*http2.Server, error) {
	return ignite.NewOptionsWithPath[http2.Server](root)
}

// NewServerWithPath returns a pointer with new Server
func NewServerWithPath(path string) (srv *http2.Server, err error) {
	return ignite.NewOptionsWithPath[http2.Server](root, path)
}
