package api

import (
	"github.com/JulianToledano/goingecko/v3/api/internal"
)

var (
	// Version represents the CoinGecko API version used by this client
	Version = internal.Version
	// BaseURL is the base URL for the CoinGecko public API
	BaseURL = internal.BaseURL
	// ProBaseURL is the base URL for the CoinGecko Pro API which requires authentication
	ProBaseURL = internal.ProBaseURL
)

// GeckoApiVersion returns the version of the CoinGecko API client library
func GeckoApiVersion() string {
	return internal.CoinGeckoAPIVersion
}
