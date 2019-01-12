package options // import "github.com/tgglv/wc-api-go/options"

import (
	"strconv"
	"time"
)

// Default WooCommerce REST API version
const defaultVersion string = "v3"

// Default request timeout.
const timeout int = 15

// Default WP API prefix. Including leading and trailing slashes.
const defaultPrefix string = "/wp-json/"

// DefaultAPIPrefix ...
const DefaultAPIPrefix = "/wc-api/"

// DefaultUserAgent ... No version number.
const DefaultUserAgent string = "WooCommerce API Client-Go"

// Basic options of WooCommerce API Client
type Basic struct {
	URL     string   // Store URL
	Key     string   // API Consumer Key
	Secret  string   // API Consumer Secret
	Options Advanced // Advanced Options (see below)
}

// Advanced options of WooCommerce API Client
type Advanced struct {
	WPAPI                bool   // Allow make requests to the new WP REST API integration
	WPAPIPrefix          string // Custom WP REST API URL prefix
	Version              string // API version, default is v3
	Timeout              int    // Request timeout, default is 15
	FollowRedirects      bool   // Allow the API call to follow redirects
	verifySslInitialized bool
	VerifySsl            bool   // Verify SSL when connect
	QueryStringAuth      bool   // Force Basic Authentication
	OAuthTimestamp       string // Custom oAuth timestamp
	UserAgent            string // Custom user-agent
}

// Version of WooCommerce API which will be used
func (o *Basic) Version() string {
	if "" == o.Options.Version {
		return defaultVersion
	}
	return o.Options.Version
}

// VerifySsl return requirement of SSL certificate check
func (o *Basic) VerifySsl() bool {
	if false == o.Options.verifySslInitialized {
		return true
	}
	return o.Options.VerifySsl
}

// DisableSslVerification in case of debug
func (o *Basic) DisableSslVerification() {
	o.Options.verifySslInitialized = true
	o.Options.VerifySsl = false
}

// Timeout per request to WooCommerce API
func (o *Basic) Timeout() int {
	if 0 == o.Options.Timeout {
		return timeout
	}
	return o.Options.Timeout
}

// QueryStringAuth returns true when credentials will pass through query string
func (o *Basic) QueryStringAuth() bool {
	return o.Options.QueryStringAuth
}

// WPAPI returns if used WordPress API or not
func (o *Basic) WPAPI() bool {
	return o.Options.WPAPI
}

// WPAPIPrefix is a prefix for WordPress API requests
func (o *Basic) WPAPIPrefix() string {
	if "" == o.Options.WPAPIPrefix {
		return defaultPrefix
	}
	return o.Options.WPAPIPrefix
}

// OAuthTimestamp return OAuth specific timestamp
func (o *Basic) OAuthTimestamp() string {
	if "" == o.Options.OAuthTimestamp {
		return strconv.FormatInt(time.Now().Unix(), 10)
	}
	return o.Options.OAuthTimestamp
}

// UserAgent which will be use for requests
func (o *Basic) UserAgent() string {
	if "" == o.Options.UserAgent {
		return DefaultUserAgent
	}
	return o.Options.UserAgent
}

// FollowRedirects during requests
func (o *Basic) FollowRedirects() bool {
	return o.Options.FollowRedirects
}
