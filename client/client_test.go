package client

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/tgglv/wc-api-go/request"
	"github.com/tgglv/wc-api-go/test"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestRequest(t *testing.T) {
	parameters := url.Values{}
	parameters.Set("foo", "bar")

	methods := []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}

	Assert := assert.New(t)

	for _, method := range methods {
		t.Logf("Test method: %s", method)
		request := request.Request{
			Method:   method,
			Endpoint: "products",
			Values:   parameters,
		}

		sender := getSenderMock(request, getResponseMock(method))
		client := Client{
			sender: sender,
		}

		r, _ := executeRequest(client, &request)

		body, _ := ioutil.ReadAll(r.Body)
		Assert.Equal(getResponseBody(method), string(body))

		err := r.Body.Close()
		if err != nil {
			t.Errorf("Failed to close body of response")
		}
	}
}

func getSenderMock(request request.Request, response *http.Response) *SenderMock {
	sender := SenderMock{
		response: *response,
	}
	return &sender
}

func executeRequest(c Client, r *request.Request) (*http.Response, error) {
	switch r.Method {
	case "GET":
		return c.Get(r.Endpoint, r.Values)
	case "POST":
		return c.Post(r.Endpoint, r.Values)
	case "PUT":
		return c.Put(r.Endpoint, r.Values)
	case "DELETE":
		return c.Delete(r.Endpoint, r.Values)
	case "OPTIONS":
		return c.Options(r.Endpoint)
	default:
		return nil, errors.New("incorrect request method")
	}
}

func getResponseMock(method string) *http.Response {
	body := getResponseBody(method)
	response := test.Response{}
	return response.GetWithBody(body)
}

func getResponseBody(method string) string {
	return "Hello " + method + "!"
}
