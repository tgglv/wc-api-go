package auth

import (
	"github.com/stretchr/testify/assert"
	"github.com/tgglv/wc-api-go/options"
	"net/http"
	"testing"
)

func TestEnrichRequest(t *testing.T) {
	Assert := assert.New(t)
	tests := map[string]struct {
		url          string
		expectHeader http.Header
	}{
		"HTTP": {
			url:          "http://woo.dev/",
			expectHeader: http.Header{},
		},
		"HTTPS": {
			url: "https://woo.dev/",
			expectHeader: http.Header{
				"Authorization": []string{"Basic a2V5OnNlY3JldA=="},
			},
		},
	}

	authenticator := Authenticator{
		options: options.Basic{
			Key:    "key",
			Secret: "secret",
		},
	}
	for caseName, test := range tests {
		t.Logf("Test case: %s", caseName)

		request, _ := http.NewRequest("GET", test.url, nil)
		authenticator.EnrichRequest(request, test.url)
		Assert.Equal(test.expectHeader, request.Header)
	}
}

func TestIsSsl(t *testing.T) {
	Assert := assert.New(t)
	tests := map[string]struct {
		url          string
		expectResult bool
	}{
		"HTTP": {
			url:          "http://woo.dev/",
			expectResult: false,
		},
		"HTTPS": {
			url:          "https://woo.dev/",
			expectResult: true,
		},
	}

	authenticator := Authenticator{}
	for caseName, test := range tests {
		t.Logf("Test case: %s", caseName)
		Assert.Equal(test.expectResult, authenticator.IsSsl(test.url))
	}
}
