package auth

import (
	"github.com/stretchr/testify/assert"
	"github.com/tgglv/wc-api-go/options"
	"net/url"
	"testing"
)

const KEY = "key"
const SECRET = "secret"

func TestGetParameters(t *testing.T) {
	Assert := assert.New(t)
	tests := map[string]struct {
		options                 options.Basic
		parameterValuesExpected bool
	}{
		"Non Query String Based Auth": {
			options:                 getOptions(false),
			parameterValuesExpected: false,
		},
		"Query String Based Auth": {
			options:                 getOptions(true),
			parameterValuesExpected: true,
		},
	}

	for caseName, test := range tests {
		t.Logf("Test case: %s", caseName)
		b := BasicAuthentication{}
		p := b.GetEnrichedQuery(getParameters(), test.options)
		if true == test.parameterValuesExpected {
			assertURLValuesHasExactValue(Assert, p, "consumer_key", KEY)
			assertURLValuesHasExactValue(Assert, p, "consumer_secret", SECRET)
		} else {
			assertURLValuesHasNoKey(Assert, p, "consumer_key")
			assertURLValuesHasNoKey(Assert, p, "consumer_secret")
		}
	}
}

func assertURLValuesHasExactValue(Assert *assert.Assertions, p url.Values, key string, expectedValue string) {
	value := p.Get(key)
	Assert.NotEqual("", value)
	Assert.Equal(expectedValue, value)
}

func assertURLValuesHasNoKey(Assert *assert.Assertions, p url.Values, key string) {
	Assert.Equal("", p.Get(key))
}

func getParameters() url.Values {
	parameters := url.Values{}
	parameters.Set("foo", "bar")
	return parameters
}

func getOptions(queryStringAuth bool) options.Basic {
	return options.Basic{
		Key:    KEY,
		Secret: SECRET,
		Options: options.Advanced{
			QueryStringAuth: queryStringAuth,
		},
	}
}
