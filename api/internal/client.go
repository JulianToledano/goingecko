package internal

import (
	"net/http"

	geckohttp "github.com/JulianToledano/goingecko/v3/http"
)

type Client struct {
	geckohttp.HttpClient

	URL string
}

func NewClient(httpClient geckohttp.HttpClient, url string) *Client {
	return &Client{
		HttpClient: httpClient,
		URL:        url,
	}
}

// CommonTestClient is a test client for use in tests so rate limiting is not an issue
var CommonTestClient = geckohttp.NewRateLimitedClient(
	geckohttp.WithHttpClient[*geckohttp.RateLimitedClient](http.DefaultClient),
	geckohttp.WithRateLimit[*geckohttp.RateLimitedClient](15),
	geckohttp.WithRetryPolicy[*geckohttp.RateLimitedClient](5, 2),
)
