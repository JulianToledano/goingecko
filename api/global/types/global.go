package types

import "github.com/JulianToledano/goingecko/v3/api/types"

type Global struct {
	Data data `json:"data"`
}

type data struct {
	ActiveCryptocurrencies          int64              `json:"active_cryptocurrencies"`
	UpcomingIcos                    int64              `json:"upcoming_icos"`
	OngoingIcos                     int64              `json:"ongoing_icos"`
	EndedIcos                       int64              `json:"ended_icos"`
	Markets                         int64              `json:"markets"`
	TotalMarketCap                  types.PriceRates   `json:"total_market_cap"`
	TotalVolume                     types.PriceRates   `json:"total_volume"`
	MarketCapPercentage             map[string]float64 `json:"market_cap_percentage"`
	MarketCapChangePercentage24hUsd float64            `json:"market_cap_change_percentage_24h_usd"`
	UpdatedAt                       float64            `json:"updated_at"`
}
