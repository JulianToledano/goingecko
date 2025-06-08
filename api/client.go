package api

import (
	"net/http"
	"time"

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

const (
	proApiKeyHeaderKey  = "x-cg-pro-api-key"
	demoApiKeyHeaderKey = "x-cg-demo-api-key"
)

// proApiHeader returns a function that sets the Pro API key header on requests
func proApiHeader(apiKey string) geckohttp.ApiHeaderFn {
	return func(r *http.Request) {
		r.Header.Set(proApiKeyHeaderKey, apiKey)
	}
}

// demoApiHeader returns a function that sets the Demo API key header on requests
func demoApiHeader(apiKey string) geckohttp.ApiHeaderFn {
	return func(r *http.Request) {
		r.Header.Set(demoApiKeyHeaderKey, apiKey)
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

// NewClient creates a new Client with the given gecko HTTP client and base URL.
func NewClient(c geckohttp.HttpClient, url string) *Client {
	return newClient(c, url)
}

// NewDefaultClient creates a new Client using the default HTTP client and base URL
func NewDefaultClient() *Client {
	return newClient(
		geckohttp.NewClient(
			geckohttp.WithHttpClient[*geckohttp.Client](http.DefaultClient),
		),
		BaseURL,
	)
}

// NewDefaultRateLimitedClient creates a new Client using the default HTTP client and base URL
func NewDefaultRateLimitedClient() *Client {
	return newClient(
		geckohttp.NewRateLimitedClient(
			geckohttp.WithHttpClient[*geckohttp.RateLimitedClient](http.DefaultClient),
			geckohttp.WithRateLimit[*geckohttp.RateLimitedClient](15),
			geckohttp.WithRetryPolicy[*geckohttp.RateLimitedClient](5, 2),
		),
		BaseURL,
	)
}

// NewDemoApiClient creates a new Client configured for the Demo API with the provided API key and HTTP client
func NewDemoApiClient(apiKey string, c *http.Client) *Client {
	return newClient(
		geckohttp.NewClient(
			geckohttp.WithHttpClient[*geckohttp.Client](c),
			geckohttp.WithApiHeaderFn[*geckohttp.Client](demoApiHeader(apiKey)),
		),
		BaseURL,
	)
}

// NewDemoApiRateLimitedClient creates a new Client configured for the Demo API with rate limiting.
// It takes an API key, HTTP client, and rate limiting configuration:
//   - apiKey: The API key for authentication
//   - c: The HTTP client to use for requests
//   - reqPerMinute: Maximum number of requests allowed per minute
//
// The client is configured with 5 retry attempts and a base delay of 2 seconds between retries.
func NewDemoApiRateLimitedClient(apiKey string, c *http.Client, reqPerMinute int) *Client {
	return newClient(
		geckohttp.NewRateLimitedClient(
			geckohttp.WithHttpClient[*geckohttp.RateLimitedClient](c),
			geckohttp.WithApiHeaderFn[*geckohttp.RateLimitedClient](demoApiHeader(apiKey)),
			geckohttp.WithRateLimit[*geckohttp.RateLimitedClient](reqPerMinute),
			geckohttp.WithRetryPolicy[*geckohttp.RateLimitedClient](5, 2)),
		BaseURL,
	)
}

// NewProApiClient creates a new Client configured for the Pro API with the provided API key and HTTP client
func NewProApiClient(apiKey string, c *http.Client) *Client {
	return newClient(
		geckohttp.NewClient(
			geckohttp.WithHttpClient[*geckohttp.Client](c),
			geckohttp.WithApiHeaderFn[*geckohttp.Client](proApiHeader(apiKey)),
		),
		ProBaseURL,
	)
}

// NewProApiRateLimitedClient creates a new Client configured for the Pro API with rate limiting.
// It takes an API key, HTTP client, and rate limiting configuration parameters:
//   - reqPerMinute: Maximum number of requests allowed per minute
//   - maxRetries: Maximum number of retry attempts for failed requests
//   - baseDelay: Initial delay between retries, which will be exponentially increased
func NewProApiRateLimitedClient(apiKey string, c *http.Client, reqPerMinute, maxRetries int, baseDelay time.Duration) *Client {
	return newClient(
		geckohttp.NewRateLimitedClient(
			geckohttp.WithHttpClient[*geckohttp.RateLimitedClient](c),
			geckohttp.WithApiHeaderFn[*geckohttp.RateLimitedClient](proApiHeader(apiKey)),
			geckohttp.WithRateLimit[*geckohttp.RateLimitedClient](reqPerMinute),
			geckohttp.WithRetryPolicy[*geckohttp.RateLimitedClient](maxRetries, baseDelay),
		),
		ProBaseURL,
	)
}

// newClient creates a new Client with the provided HTTP client and base URL
func newClient(c geckohttp.HttpClient, url string) *Client {
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
