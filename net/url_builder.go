package net // import "github.com/tgglv/wc-api-go/net"

import (
	"github.com/tgglv/wc-api-go/request"
)

// URLBuilder interface
type URLBuilder interface {
	GetURL(req request.Request) string
}
