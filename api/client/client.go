package client

import (
	"net/http"

	"github.com/JulianToledano/goingecko/api"

	"github.com/JulianToledano/goingecko/api/coins"
	geckohttp "github.com/JulianToledano/goingecko/http"
)

// proApiHeader returns a function that sets the Pro API key header on requests
func proApiHeader(apiKey string) func(r *http.Request) {
	return func(r *http.Request) {
		r.Header.Set("x-cg-pro-api-key", apiKey)
	}
}

// demoApiHeader returns a function that sets the Demo API key header on requests
func demoApiHeader(apiKey string) func(r *http.Request) {
	return func(r *http.Request) {
		r.Header.Set("x-cg-demo-api-key", apiKey)
	}
}

// Client wraps the CoinGecko API client functionality
type Client struct {
	*coins.Client

	url string
}

// NewDefaultClient creates a new Client using the default HTTP client and base URL
func NewDefaultClient() *Client {
	return newClient(
		geckohttp.NewClient(geckohttp.WithHttpClient(http.DefaultClient)),
		api.BaseURL,
	)
}

// NewDemoApiClient creates a new Client configured for the Demo API with the provided API key and HTTP client
func NewDemoApiClient(apiKey string, c *http.Client) *Client {
	return newClient(
		geckohttp.NewClient(geckohttp.WithHttpClient(c), geckohttp.WithApiHeaderFn(demoApiHeader(apiKey))),
		api.BaseURL,
	)
}

// NewProApiClient creates a new Client configured for the Pro API with the provided API key and HTTP client
func NewProApiClient(apiKey string, c *http.Client) *Client {
	return newClient(
		geckohttp.NewClient(geckohttp.WithHttpClient(c), geckohttp.WithApiHeaderFn(proApiHeader(apiKey))),
		api.ProBaseURL,
	)
}

// newClient creates a new Client with the provided HTTP client and base URL
func newClient(c *geckohttp.Client, url string) *Client {
	return &Client{
		Client: coins.NewCoinsClient(c, url),
	}
}