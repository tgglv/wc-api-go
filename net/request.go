package net // import "github.com/tgglv/wc-api-go/net"

import "net/http"

// RequestEnricher adds Basic Authentication settings in Request in case of Basic Authentication
type RequestEnricher interface {
	EnrichRequest(r *http.Request, URL string)
}
