package goingecko

import "fmt"

var (
	version           = "v3"
	baseURL           = fmt.Sprintf("https://api.coingecko.com/api/%s", version)
	pingURL           = fmt.Sprintf("%s/ping", baseURL)
	simpleURL         = fmt.Sprintf("%s/simple", baseURL)
	coinsURL          = fmt.Sprintf("%s/coins", baseURL)
	assetPlatformsURL = fmt.Sprintf("%s/asset_platforms", baseURL)
	categoriesURL     = fmt.Sprintf("%s/categories", coinsURL)
	exchangesURL      = fmt.Sprintf("%s/exchanges", baseURL)
	derivativesURL    = fmt.Sprintf("%s/derivatives", baseURL)
	contractURL       = fmt.Sprintf("%s/coins", baseURL)
	exchangeRatesURL  = fmt.Sprintf("%s/exchange_rates", baseURL)
	trendingURL       = fmt.Sprintf("%s/search/trending", baseURL)
	globalURL         = fmt.Sprintf("%s/global", baseURL)
)
