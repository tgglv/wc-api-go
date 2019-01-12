package net

import (
	"net/http"
)

// ClientMock imitates sending requests via standard library
type ClientMock struct {
	request  http.Request
	response *http.Response
	err      error
}

// Do method imitates request and returns previously saved response and error
func (c *ClientMock) Do(req *http.Request) (*http.Response, error) {
	c.request = *req
	return c.response, c.err
}
