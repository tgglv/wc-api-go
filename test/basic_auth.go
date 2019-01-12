package test

import "encoding/base64"

// BasicAuthentication calculates expected authorization header value
type BasicAuthentication struct {
}

// GetBasicAuth ...
func (b *BasicAuthentication) GetBasicAuth(username, password string) string {
	auth := username + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}
