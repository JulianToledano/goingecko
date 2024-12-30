package client

import (
	"net/http"

	"github.com/JulianToledano/goingecko/api"
	"github.com/JulianToledano/goingecko/api/assetPlatforms"
	"github.com/JulianToledano/goingecko/api/categories"
	"github.com/JulianToledano/goingecko/api/coins"
	"github.com/JulianToledano/goingecko/api/contract"
	"github.com/JulianToledano/goingecko/api/derivatives"
	"github.com/JulianToledano/goingecko/api/exchanges"
	"github.com/JulianToledano/goingecko/api/ping"
	"github.com/JulianToledano/goingecko/api/simple"
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
	*ping.PingClient
	*simple.SimpleClient
	*coins.CoinsClient
	*contract.ContractClient
	*assetPlatforms.AssetPlatformsClient
	*categories.CategoriesClient
	*exchanges.ExchangesClient
	*derivatives.DerivativesClient

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
		PingClient:           ping.NewClient(c, url),
		SimpleClient:         simple.NewClient(c, url),
		CoinsClient:          coins.NewClient(c, url),
		ContractClient:       contract.NewClient(c, url),
		AssetPlatformsClient: assetPlatforms.NewClient(c, url),
		CategoriesClient:     categories.NewClient(c, url),
		ExchangesClient:      exchanges.NewClient(c, url),
		DerivativesClient:    derivatives.NewClient(c, url),
	}
}
