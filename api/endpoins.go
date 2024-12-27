package api

import "fmt"

var (
	Version    = "v3"
	BaseURL    = fmt.Sprintf("https://api.coingecko.com/api/%s", Version)
	ProBaseURL = fmt.Sprintf("https://pro-api.coingecko.com/api/%s", Version)
)
