package test

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// Response structure creates new http.Response with specified body
type Response struct {
}

// GetWithBody creates response
func (r *Response) GetWithBody(body string) *http.Response {
	return &http.Response{
		StatusCode:    http.StatusOK,
		Proto:         "HTTP/1.0",
		Body:          ioutil.NopCloser(bytes.NewReader([]byte(body))),
		ContentLength: int64(len(body)),
	}
}
