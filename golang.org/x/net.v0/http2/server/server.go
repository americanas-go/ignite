package server

import (
	"github.com/americanas-go/config"
	"golang.org/x/net/http2"
)

// NewServer returns a pointer with new Server
func NewServer() (*http2.Server, error) {
	srv := &http2.Server{}

	err := config.UnmarshalWithPath(root, srv)
	if err != nil {
		return nil, err
	}

	return srv, nil
}

// NewServerWithPath returns a pointer with new Server
func NewServerWithPath(path string) (srv *http2.Server, err error) {
	srv, err = NewServer()
	if err != nil {
		return nil, err
	}

	err = config.UnmarshalWithPath(path, srv)
	if err != nil {
		return nil, err
	}

	return srv, nil
}
