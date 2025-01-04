package api

import (
	"net/http"

	"github.com/JulianToledano/goingecko/v3/api/assetPlatforms"
	"github.com/JulianToledano/goingecko/v3/api/categories"
	"github.com/JulianToledano/goingecko/v3/api/coins"
	"github.com/JulianToledano/goingecko/v3/api/companies"
	"github.com/JulianToledano/goingecko/v3/api/contract"
	"github.com/JulianToledano/goingecko/v3/api/derivatives"
	"github.com/JulianToledano/goingecko/v3/api/exchangeRates"
	"github.com/JulianToledano/goingecko/v3/api/exchanges"
	"github.com/JulianToledano/goingecko/v3/api/global"
	"github.com/JulianToledano/goingecko/v3/api/nfts"
	"github.com/JulianToledano/goingecko/v3/api/ping"
	"github.com/JulianToledano/goingecko/v3/api/search"
	"github.com/JulianToledano/goingecko/v3/api/simple"
	"github.com/JulianToledano/goingecko/v3/api/trending"
	geckohttp "github.com/JulianToledano/goingecko/v3/http"
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
	*nfts.NftsClient
	*exchangeRates.ExchangeRatesClient
	*search.SearchClient
	*trending.TrendingClient
	*global.GlobalClient
	*companies.CompaniesClient
}

// NewDefaultClient creates a new Client using the default HTTP client and base URL
func NewDefaultClient() *Client {
	return newClient(
		geckohttp.NewClient(geckohttp.WithHttpClient(http.DefaultClient)),
		BaseURL,
	)
}

// NewDemoApiClient creates a new Client configured for the Demo API with the provided API key and HTTP client
func NewDemoApiClient(apiKey string, c *http.Client) *Client {
	return newClient(
		geckohttp.NewClient(geckohttp.WithHttpClient(c), geckohttp.WithApiHeaderFn(demoApiHeader(apiKey))),
		BaseURL,
	)
}

// NewProApiClient creates a new Client configured for the Pro API with the provided API key and HTTP client
func NewProApiClient(apiKey string, c *http.Client) *Client {
	return newClient(
		geckohttp.NewClient(geckohttp.WithHttpClient(c), geckohttp.WithApiHeaderFn(proApiHeader(apiKey))),
		ProBaseURL,
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
		NftsClient:           nfts.NewClient(c, url),
		ExchangeRatesClient:  exchangeRates.NewClient(c, url),
		SearchClient:         search.NewClient(c, url),
		TrendingClient:       trending.NewClient(c, url),
		GlobalClient:         global.NewClient(c, url),
		CompaniesClient:      companies.NewClient(c, url),
	}
}
