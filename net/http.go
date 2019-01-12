package net // import "github.com/tgglv/wc-api-go/net"

import (
	"io"
	"net/http"
)

// RequestCreator creates an abstraction for standard method for creating requests
type RequestCreator interface {
	NewRequest(method, url string, body io.Reader) (*http.Request, error)
}

// Client creates an abstraction for standard method for sending requests
type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

// HTTP creates new requests like the standard library
type HTTP struct {
}

// NewRequest method creates new http.Request
func (h *HTTP) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequest(method, url, body)
}
