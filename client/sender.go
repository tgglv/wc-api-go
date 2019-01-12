package client // import "github.com/tgglv/wc-api-go/client"

import (
	"github.com/tgglv/wc-api-go/request"
	"net/http"
)

// Sender interface
type Sender interface {
	Send(req request.Request) (resp *http.Response, err error)
}
