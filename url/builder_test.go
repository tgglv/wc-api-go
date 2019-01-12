package url

import (
	"github.com/stretchr/testify/assert"
	"github.com/tgglv/wc-api-go/options"
	"github.com/tgglv/wc-api-go/request"
	URL "net/url"
	"testing"
)

func TestGetURL(t *testing.T) {
	Assert := assert.New(t)

	defaultOptions := options.Basic{
		URL:     "http://woo.dev/",
		Options: options.Advanced{},
	}

	tests := map[string]struct {
		method               string
		endpoint             string
		values               URL.Values
		useValuesForEnricher bool
		options              options.Basic
		expectedURL          string
	}{
		"Searchings for Products": {
			method:   "GET",
			endpoint: "products",
			values: URL.Values{
				"search": []string{"awesome"},
			},
			useValuesForEnricher: true,
			options:              defaultOptions,
			expectedURL:          "http://woo.dev/wc-api/v3/products?search=awesome",
		},
		"Create a Coupon": {
			method:   "POST",
			endpoint: "coupons",
			values: URL.Values{
				"code": []string{"50off"},
			},
			useValuesForEnricher: false,
			options:              defaultOptions,
			expectedURL:          "http://woo.dev/wc-api/v3/coupons?",
		},
		"Update a product review": {
			method:   "PUT",
			endpoint: "products/reviews",
			values: URL.Values{
				"rating": []string{"10"},
			},
			useValuesForEnricher: false,
			options:              defaultOptions,
			expectedURL:          "http://woo.dev/wc-api/v3/products/reviews?",
		},
		"Delete an Order": {
			method:   "GET",
			endpoint: "orders/123",
			values: URL.Values{
				"force": []string{"true"},
			},
			useValuesForEnricher: true,
			options:              defaultOptions,
			expectedURL:          "http://woo.dev/wc-api/v3/orders/123?force=true",
		},
	}

	for name, test := range tests {
		t.Logf("Test case: %s", name)

		request := request.Request{
			Method:   test.method,
			Endpoint: test.endpoint,
			Values:   test.values,
		}

		var queryForEnricher URL.Values
		if test.useValuesForEnricher {
			queryForEnricher = test.values
		} else {
			queryForEnricher = URL.Values{}
		}
		queryEnricher := QueryEnricherMock{
			query: queryForEnricher,
		}

		builder := Builder{}
		builder.SetOptions(test.options)
		builder.SetQueryEnricher(&queryEnricher)
		Assert.Equal(test.expectedURL, builder.GetURL(request))
	}
}

func TestGetFilteredQuery(t *testing.T) {
	Assert := assert.New(t)

	sampleValues := URL.Values{}
	sampleValues.Set("foo", "bar")

	tests := map[string]struct {
		req           request.Request
		expectedQuery URL.Values
	}{
		"GET": {
			req: request.Request{
				Method: "GET",
				Values: sampleValues,
			},
			expectedQuery: sampleValues,
		},
		"GET with empty values": {
			req: request.Request{
				Method: "GET",
				Values: nil,
			},
			expectedQuery: URL.Values{},
		},
		"DELETE": {
			req: request.Request{
				Method: "DELETE",
				Values: sampleValues,
			},
			expectedQuery: sampleValues,
		},
		"POST": {
			req: request.Request{
				Method: "POST",
				Values: sampleValues,
			},
			expectedQuery: URL.Values{},
		},
		"PUT": {
			req: request.Request{
				Method: "PUT",
				Values: sampleValues,
			},
			expectedQuery: URL.Values{},
		},
		"OPTIONS": {
			req: request.Request{
				Method: "OPTIONS",
				Values: sampleValues,
			},
			expectedQuery: URL.Values{},
		},
	}

	for name, test := range tests {
		t.Logf("Test case: %s", name)

		builder := Builder{
			options: options.Basic{
				URL: "http://woo.dev/",
			},
		}
		Assert.Equal(test.expectedQuery, builder.getFilteredQuery(test.req))
	}
}

func TestGetBaseURL(t *testing.T) {
	Assert := assert.New(t)
	tests := map[string]struct {
		WPAPI           bool
		WPAPIPrefix     string
		version         string
		expectedBaseURL string
	}{
		"Default API Prefix": {
			expectedBaseURL: "http://woo.dev/wc-api/v3/",
		},
	}

	for name, test := range tests {
		t.Logf("Test case: %s", name)

		builder := Builder{
			options: options.Basic{
				URL: "http://woo.dev/",
			},
		}
		Assert.Equal(test.expectedBaseURL, builder.getBaseURL())
	}
}

func TestGetAPIPrefix(t *testing.T) {
	Assert := assert.New(t)
	tests := map[string]struct {
		WPAPI          bool
		WPAPIPrefix    string
		expectedPrefix string
	}{
		"WordPress API wasn't set up": {
			WPAPI:          false,
			expectedPrefix: "/wc-api/",
		},
		"WordPress API was set up": {
			WPAPI:          true,
			WPAPIPrefix:    "/wc-test/",
			expectedPrefix: "/wc-test/",
		},
		"Default API Prefix": {
			WPAPI:          false,
			WPAPIPrefix:    "/wc-test/",
			expectedPrefix: "/wc-api/",
		},
	}

	for name, test := range tests {
		t.Logf("Test case: %s", name)

		builder := Builder{}
		builder.SetOptions(options.Basic{
			Options: options.Advanced{
				WPAPI:       test.WPAPI,
				WPAPIPrefix: test.WPAPIPrefix,
			},
		})

		Assert.Equal(test.expectedPrefix, builder.getAPIPrefix())
	}
}
