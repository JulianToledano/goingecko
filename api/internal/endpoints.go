package internal

import "fmt"

var (
	// Version represents the CoinGecko API version used by this client
	Version = "v3"
	// BaseURL is the base URL for the CoinGecko public API
	BaseURL = fmt.Sprintf("https://api.coingecko.com/api/%s", Version)
	// ProBaseURL is the base URL for the CoinGecko Pro API which requires authentication
	ProBaseURL = fmt.Sprintf("https://pro-api.coingecko.com/api/%s", Version)
	// CoinGeckoAPIVersion represents the version of the CoinGecko API implementation in this client
	CoinGeckoAPIVersion = "3.1.1"
)
