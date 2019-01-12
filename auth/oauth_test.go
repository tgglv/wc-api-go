package auth

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestOAuth(t *testing.T) {
	Assert := assert.New(t)
	microtimer := MicroTimerMock{
		result: "0.38716500 1541755274",
	}

	tests := map[string]struct {
		oauth    OAuth
		expected url.Values
	}{
		"Basic": {
			OAuth{
				URL:        "http://127.0.0.1:8000/wp-json/wc/v3/products",
				Key:        "ck_085880997173f0faa8e30616ed37b8eb77f0c467",
				Secret:     "cs_5e9979b3eaeee413961da11113cdec7f89415a23",
				Version:    "wc/v3",
				Method:     "GET",
				Parameters: url.Values{},
				Timestamp:  "1541755201",
			},
			url.Values{
				"oauth_consumer_key":     []string{"ck_085880997173f0faa8e30616ed37b8eb77f0c467"},
				"oauth_nonce":            []string{"aca37845644a99011e8b3922bc1d251f3d9af370"},
				"oauth_signature":        []string{"e5t2rJibFcNY0llLFzhp7JSYN/nW9jJpwHyZYntrLXQ="},
				"oauth_signature_method": []string{"HMAC-SHA256"},
				"oauth_timestamp":        []string{"1541755201"},
			},
		},
	}
	for caseName, test := range tests {
		t.Logf("Test case: %s", caseName)
		test.oauth.SetMicrotimer(&microtimer)
		Assert.Equal(test.expected, test.oauth.GetEnrichedQuery(), "Case: "+caseName)
	}
}
