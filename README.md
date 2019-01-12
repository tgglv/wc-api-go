# WooCommerce API - Golang Client

A Golang wrapper for the WooCommerce REST API. Easily interact with the WooCommerce REST API securely using this library. If using a HTTPS connection this library uses BasicAuth, else it uses Oauth to provide a secure connection to WooCommerce.

[![Build Status](https://travis-ci.com/tgglv/wc-api-go.svg?branch=master)](https://travis-ci.com/tgglv/wc-api-go)

## Installation

To install this WooCommerce REST API Golang Wrapper, use `go get`:
```
go get github.com/tgglv/wc-api-go
```

## Staying up to date

To update WooCommerce REST API Golang Wrapper to the latest version, use
```
go get -u github.com/tgglv/wc-api-go
```

## Getting started

Generate API credentials (Consumer Key & Consumer Secret) following this instructions <http://docs.woocommerce.com/document/woocommerce-rest-api/>.

Check out the WooCommerce API endpoints and data that can be manipulated in <https://woocommerce.github.io/woocommerce-rest-api-docs/>.

## Setup

Setup for the new WP REST API integration:

```go
import "github.com/tgglv/wc-api-go/client"

factory := client.Factory{}
c := factory.NewClient()

c.SetOptions(client.BasicOptions{
    URL:    "http://example.com",
    Key:    "ck_XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
    Secret: "cs_XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
    Options: client.AdvancedOptions{
        WPAPI:       true,
        WPAPIPrefix: "/wp-json/",
        Version:     "wc/v3",
    },
})
```

### Options

|       Option      |   Type   | Required |                Description                 |
| ----------------- | -------- | -------- | ------------------------------------------ |
| `URL`             | `string` | yes      | Your Store URL, example: http://woo.dev/   |
| `Key`    | `string` | yes      | Your API consumer key                      |
| `Secret` | `string` | yes      | Your API consumer secret                   |
| `Options`         | `struct`  | no       | Extra arguments (see client options table) |

#### Advanced options

|        Option       |   Type   | Required |                                                      Description                                                       |
|---------------------|----------|----------|------------------------------------------------------------------------------------------------------------------------|
| `WPAPI`            | `bool`   | no       | Allow make requests to the new WP REST API integration (WooCommerce 2.6 or later)                                      |
| `WPAPIPrefix`     | `string` | no       | Custom WP REST API URL prefix, used to support custom prefixes created with the `rest_url_prefix` filter               |
| `Version`           | `string` | no       | API version, default is `v3`                                                                                           |
| `Timeout`           | `int`    | no       | Request timeout, default is `15`                                                                                       |
| `FollowRedirects`  | `bool`   | no       | Allow the API call to follow redirects                                                                                 |
| `VerifySsl`        | `bool`   | no       | Verify SSL when connect, use this option as `false` when need to test with self-signed certificates, default is `true` |
| `QueryStringAuth` | `bool`   | no       | Force Basic Authentication as query string when `true` and using under HTTPS, default is `false`                       |
| `OAuthTimestamp`   | `string` | no       | Custom oAuth timestamp, default is `time()`                                                                            |
| `UserAgent`        | `string` | no       | Custom user-agent, default is `WooCommerce API Client-PHP`                                                             |

## Methods

|    Params    |   Type   |                         Description                          |
| ------------ | -------- | ------------------------------------------------------------ |
| `endpoint`   | `string` | WooCommerce API endpoint, example: `customers` or `order/12` |
| `data`       | `array`  | Only for POST and PUT, data that will be converted to JSON   |
| `parameters` | `array`  | Only for GET and DELETE, request query string                |

### GET

```go
c.Get(endpoint, parameters)
```

### POST

```go
c.Post(endpoint, data)
```

### PUT

```go
c.Put(endpoint, data)
```

### DELETE

```go
c.Delete(endpoint, parameters)
```

### OPTIONS

```go
c.Options(endpoint)
```

#### Response

All methods will return `*http.Response` on success and returning an  `error` on failure.


```go
package main

import (
	"fmt"
	"github.com/tgglv/wc-api-go/client"
	"github.com/tgglv/wc-api-go/options"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	factory := client.Factory{}
	c := factory.NewClient(options.Basic{
		URL:    "http://woo.dev/",
		Key:    "ck_XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		Secret: "cs_XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		Options: options.Advanced{
			WPAPI:       true,
			WPAPIPrefix: "/wp-json/",
			Version:     "wc/v3",
		},
	})

	if r, err := c.Get("products", nil); err != nil {
		log.Fatal(err)
	} else if r.StatusCode != http.StatusOK {
		log.Fatal("Unexpected StatusCode:", r)
	} else {
		defer r.Body.Close()
		if bodyBytes, err := ioutil.ReadAll(r.Body); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(string(bodyBytes))
		}
	}
}
```

## Release History

- 2019-01-12 - 1.0.0 - First Release