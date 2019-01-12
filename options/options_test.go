package options

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestOptionsInit(t *testing.T) {
	tests := map[string]struct {
		options         Basic
		version         string
		verifySsl       bool
		timeout         int
		queryStringAuth bool
		wpAPI           bool
		prefix          string
		timestampFrom   string
		timestampTo     string
		userAgent       string
		followRedirects bool
	}{
		"Empty constructor init": {
			options:         Basic{},
			version:         "v3",
			verifySsl:       true,
			timeout:         15,
			queryStringAuth: false,
			wpAPI:           false,
			prefix:          "/wp-json/",
			timestampFrom:   getTime(0),
			timestampTo:     getTime(1),
			userAgent:       DefaultUserAgent,
			followRedirects: false,
		},
		"Simple case": {
			options:         getSimpleCaseOptions(),
			version:         "v99",
			verifySsl:       false,
			timeout:         15,
			queryStringAuth: false,
			wpAPI:           true,
			prefix:          "/wp-api-json/",
			timestampFrom:   "1234567890",
			timestampTo:     "1234567890",
			userAgent:       "Tester",
			followRedirects: false,
		},
	}

	assert := assert.New(t)
	for caseName, test := range tests {
		t.Logf("Test case: %s", caseName)
		assert.Equal(test.version, test.options.Version())
		assert.Equal(test.verifySsl, test.options.VerifySsl())
		assert.Equal(test.timeout, test.options.Timeout())

		assert.Equal(test.queryStringAuth, test.options.QueryStringAuth())
		assert.Equal(test.wpAPI, test.options.WPAPI())
		assert.Equal(test.prefix, test.options.WPAPIPrefix())

		actualTimestamp := test.options.OAuthTimestamp()
		compareFrom := strings.Compare(test.timestampFrom, actualTimestamp)
		compareTo := strings.Compare(actualTimestamp, test.timestampTo)
		assert.True(compareFrom != 1 && compareTo != 1)

		assert.Equal(test.userAgent, test.options.UserAgent())
		assert.Equal(test.followRedirects, test.options.FollowRedirects())
	}
}

func getTime(offset int) string {
	return strconv.FormatInt(time.Now().Add(time.Second*time.Duration(offset)).Unix(), 10)
}

func getSimpleCaseOptions() Basic {
	simpleCaseOptions := Basic{
		URL:    "http://127.0.0.1",
		Key:    "ck_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		Secret: "cs_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		Options: Advanced{
			WPAPI:          true,
			WPAPIPrefix:    "/wp-api-json/",
			Version:        "v99",
			OAuthTimestamp: "1234567890",
			UserAgent:      "Tester",
		},
	}
	simpleCaseOptions.DisableSslVerification()
	return simpleCaseOptions
}
