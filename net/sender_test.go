package net

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/tgglv/wc-api-go/options"
	"github.com/tgglv/wc-api-go/request"
	"github.com/tgglv/wc-api-go/test"
	"net/url"
	"testing"
)

const defaultURL = "http://woo.dev"
const defaultEndpoint = "products"

type Expected struct {
	form url.Values
}

func TestRequest(t *testing.T) {
	Assert := assert.New(t)
	baseOptions := options.Basic{
		Key:    "key",
		Secret: "secret",
	}
	requestCreator := HTTP{}
	response := test.Response{}
	Response := response.GetWithBody("Hello!")

	OAuthURLBuilderMock := URLBuilderMock{
		url:         defaultURL,
		isBasicAuth: false,
	}

	baseClientMock := ClientMock{
		response: Response,
		err:      nil,
	}

	getRequest := request.Request{
		Method:   "GET",
		Endpoint: defaultEndpoint,
		Values:   nil,
	}

	ba := test.BasicAuthentication{}
	baseFormValues := url.Values{}
	baseFormValues.Set("foo", "bar")

	requestEnricher := RequestEnricherMock{}

	tests := map[string]struct {
		urlMock  URLBuilderMock
		client   ClientMock
		request  request.Request
		expected Expected
	}{
		"GET, Basic Authentication": {
			urlMock: URLBuilderMock{
				url:         "https://woo.dev",
				isBasicAuth: true,
			},
			client:  baseClientMock,
			request: getRequest,
			expected: Expected{
				form: nil,
			},
		},
		"GET without query, OAuth": {
			urlMock: OAuthURLBuilderMock,
			client:  baseClientMock,
			request: getRequest,
			expected: Expected{
				form: nil,
			},
		},
		"POST without data, OAuth": {
			urlMock: OAuthURLBuilderMock,
			client:  baseClientMock,
			request: request.Request{
				Method:   "POST",
				Endpoint: defaultEndpoint,
				Values:   nil,
			},
			expected: Expected{
				form: nil,
			},
		},
		"POST with data, OAuth": {
			urlMock: OAuthURLBuilderMock,
			client:  baseClientMock,
			request: request.Request{
				Method:   "POST",
				Endpoint: defaultEndpoint,
				Values:   baseFormValues,
			},
			expected: Expected{
				form: baseFormValues,
			},
		},
		"PUT with data, OAuth": {
			urlMock: OAuthURLBuilderMock,
			client:  baseClientMock,
			request: request.Request{
				Method:   "PUT",
				Endpoint: defaultEndpoint,
				Values:   baseFormValues,
			},
			expected: Expected{
				form: baseFormValues,
			},
		},
		"GET with data, OAuth": {
			urlMock: OAuthURLBuilderMock,
			client:  baseClientMock,
			request: request.Request{
				Method:   "GET",
				Endpoint: defaultEndpoint,
				Values:   baseFormValues,
			},
			expected: Expected{
				form: nil,
			},
		},
		"Network Error": {
			urlMock: OAuthURLBuilderMock,
			client: ClientMock{
				response: Response,
				err:      errors.New("Network Error"),
			},
		},
	}

	for caseName, testDetails := range tests {
		t.Logf("Test case: %s", caseName)

		sender := Sender{}
		sender.SetRequestEnricher(&requestEnricher)
		sender.SetURLBuilder(&testDetails.urlMock)
		sender.SetRequestCreator(&requestCreator)
		sender.SetHTTPClient(&testDetails.client)

		response, err := sender.Send(testDetails.request)
		request := testDetails.client.request

		authHeader := request.Header.Get("Authorization")
		if testDetails.urlMock.IsBasicAuth() {
			Assert.True(len(authHeader) > 0)
			expectedHeaderValue := ba.GetBasicAuth(baseOptions.Key, baseOptions.Secret)
			Assert.Equal(expectedHeaderValue, authHeader)
		} else {
			Assert.True(len(authHeader) == 0)
		}

		if testDetails.client.err != nil {
			Assert.Equal(testDetails.client.err, err)
		}

		if testDetails.expected.form != nil {
			Assert.Equal(testDetails.expected.form, request.Form)
		} else {
			Assert.Nil(request.Form)
		}

		Assert.Equal(Response, response)
	}
}
