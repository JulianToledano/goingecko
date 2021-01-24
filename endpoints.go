package goingecko

import "fmt"

var baseURL = "https://api.coingecko.com/api/v3"
var pingURL = fmt.Sprintf("%s/ping", baseURL)
var simpleURL = fmt.Sprintf("%s/simple", baseURL)
var coinsURL = fmt.Sprintf("%s/coins", baseURL)
var contractURL = fmt.Sprintf("%s/contract", baseURL)
var exchangesURL = fmt.Sprintf("%s/exchanges", baseURL)
var financeURL = fmt.Sprintf("%s/finance", baseURL)
var indexesURL = fmt.Sprintf("%s/indexes", baseURL)
var derivativesURL = fmt.Sprintf("%s/derivatices", baseURL)
var statusUpdatesURL = fmt.Sprintf("%s/status_updates", baseURL)
var eventsUpdatesURL = fmt.Sprintf("%s/events", baseURL)
var exchangeRatesURL = fmt.Sprintf("%s/exchange_rates", baseURL)
var trendingURL = fmt.Sprintf("%s/trending", baseURL)
var globalURL = fmt.Sprintf("%s/global", baseURL)
