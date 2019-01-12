package auth // import "github.com/tgglv/wc-api-go/auth"

import (
	"github.com/tgglv/wc-api-go/options"
	"github.com/tgglv/wc-api-go/request"
	"net/http"
	"net/url"
	"strings"
)

// BasicAuthenticationQueryEnricher ...
type BasicAuthenticationQueryEnricher interface {
	GetEnrichedQuery(p url.Values, o options.Basic) url.Values
}

type OAuthQueryEnricher interface {
}

// Authenticator ...
type Authenticator struct {
	// GetQuery(URL string, query url.Values, options options.Basic) url.Values
	options   options.Basic
	basicAuth BasicAuthentication
	oauth     OAuth
}

// EnrichRequest ...
func (a *Authenticator) EnrichRequest(r *http.Request, URL string) {
	if a.IsSsl(URL) {
		r.SetBasicAuth(a.options.Key, a.options.Secret)
	}
}

// GetEnrichedQuery ...
func (a *Authenticator) GetEnrichedQuery(url string, query url.Values, req request.Request) url.Values {
	if a.IsSsl(url) {
		return a.basicAuth.GetEnrichedQuery(query, a.options)
	}

	a.oauth.SetMethod(req.Method)
	a.oauth.SetParameters(query)
	a.oauth.SetURL(url)
	return a.oauth.GetEnrichedQuery()
}

// IsSsl method determines HTTPS protocol URL
func (a *Authenticator) IsSsl(url string) bool {
	return 0 == strings.Index(url, "https://")
}

// SetOptions ...
func (a *Authenticator) SetOptions(o options.Basic) {
	a.options = o
	a.oauth.SetOptions(o)
}

// SetOAuth method sets OAuth object to structure's inner variable
func (a *Authenticator) SetOAuth(o OAuth) {
	a.oauth = o
}

// SetBasicAuth method sets BasicAuthentication object to structure's inner variable
func (a *Authenticator) SetBasicAuth(ba BasicAuthentication) {
	a.basicAuth = ba
}
