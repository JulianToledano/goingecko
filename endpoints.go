package goingecko

import "fmt"

var (
	version          = "v3"
	baseURL          = fmt.Sprintf("https://api.coingecko.com/api/%s", version)
	pingURL          = fmt.Sprintf("%s/ping", baseURL)
	simpleURL        = fmt.Sprintf("%s/simple", baseURL)
	coinsURL         = fmt.Sprintf("%s/coins", baseURL)
	contractURL      = fmt.Sprintf("%s/coins", baseURL)
	exchangesURL     = fmt.Sprintf("%s/exchanges", baseURL)
	financeURL       = fmt.Sprintf("%s/finance", baseURL)
	indexesURL       = fmt.Sprintf("%s/indexes", baseURL)
	derivativesURL   = fmt.Sprintf("%s/derivatices", baseURL)
	statusUpdatesURL = fmt.Sprintf("%s/status_updates", baseURL)
	eventsURL        = fmt.Sprintf("%s/events", baseURL)
	exchangeRatesURL = fmt.Sprintf("%s/exchange_rates", baseURL)
	trendingURL      = fmt.Sprintf("%s/search/trending", baseURL)
	globalURL        = fmt.Sprintf("%s/global", baseURL)
)
