package client

import (
	"github.com/JulianToledano/goingecko/api"
	"net/http"

	"github.com/JulianToledano/goingecko/api/coins"
	geckohttp "github.com/JulianToledano/goingecko/http"
)

func proApiHeader(apiKey string) func(r *http.Request) {
	return func(r *http.Request) {
		r.Header.Set("x-cg-pro-api-key", apiKey)
	}
}

func demoApiHeader(apiKey string) func(r *http.Request) {
	return func(r *http.Request) {
		r.Header.Set("x-cg-demo-api-key", apiKey)
	}
}

type Client struct {
	*coins.Client

	url string
}

func NewDefaultClient() *Client {
	return newClient(
		geckohttp.NewClient(geckohttp.WithHttpClient(http.DefaultClient)),
		api.BaseURL,
	)
}

func NewDemoApiClient(apiKey string, c *http.Client) *Client {
	return newClient(
		geckohttp.NewClient(geckohttp.WithHttpClient(c), geckohttp.WithApiHeaderFn(demoApiHeader(apiKey))),
		api.ProBaseURL,
	)
}

func NewProApiClient(apiKey string, c *http.Client) *Client {
	return newClient(
		geckohttp.NewClient(geckohttp.WithHttpClient(c), geckohttp.WithApiHeaderFn(proApiHeader(apiKey))),
		api.ProBaseURL,
	)
}

func newClient(c *geckohttp.Client, url string) *Client {
	return &Client{
		Client: coins.NewCoinsClient(c, url),
	}
}
