package net // import "github.com/tgglv/wc-api-go/net"

import (
	"github.com/tgglv/wc-api-go/request"
	"net/http"
	"strings"
)

// Sender provides HTTP Requests
type Sender struct {
	requestEnricher RequestEnricher
	urlBuilder      URLBuilder
	httpClient      Client
	requestCreator  RequestCreator
}

// Send method sends requests to WooCommerce API
func (s *Sender) Send(req request.Request) (resp *http.Response, err error) {
	request := s.prepareRequest(req)
	return s.httpClient.Do(request)
}

func (s *Sender) prepareRequest(req request.Request) *http.Request {
	URL := s.urlBuilder.GetURL(req)

	var reader *strings.Reader
	hasBody := ("POST" == req.Method || "PUT" == req.Method) && len(req.JSONBody) > 0
	var JSONBody string
	if hasBody {
		JSONBody = req.JSONBody
	} else {
		JSONBody = ""
	}
	reader = strings.NewReader(JSONBody)
	request, _ := s.requestCreator.NewRequest(req.Method, URL, reader)
	if hasBody {
		request.Header.Add("Content-type", "application/json")
	}
	s.requestEnricher.EnrichRequest(request, URL)
	return request
}

// SetRequestEnricher ...
func (s *Sender) SetRequestEnricher(a RequestEnricher) {
	s.requestEnricher = a
}

// SetURLBuilder ...
func (s *Sender) SetURLBuilder(urlBuilder URLBuilder) {
	s.urlBuilder = urlBuilder
}

// SetHTTPClient ...
func (s *Sender) SetHTTPClient(c Client) {
	s.httpClient = c
}

// SetRequestCreator ...
func (s *Sender) SetRequestCreator(rc RequestCreator) {
	s.requestCreator = rc
}
