package types

import "github.com/JulianToledano/goingecko/api/types"

type Global struct {
	Data data `json:"data"`
}

type data struct {
	ActiveCryptocurrencies          int32              `json:"active_cryptocurrencies"`
	UpcomingIcos                    int32              `json:"upcoming_icos"`
	OngoingIcos                     int32              `json:"ongoing_icos"`
	EndedIcos                       int32              `json:"ended_icos"`
	Markets                         int32              `json:"markets"`
	TotalMarketCap                  types.PriceRates   `json:"total_market_cap"`
	TotalVolume                     types.PriceRates   `json:"total_volume"`
	MarketCapPercentage             map[string]float64 `json:"market_cap_percentage"`
	MarketCapChangePercentage24hUsd float64            `json:"market_cap_change_percentage_24h_usd"`
	UpdatedAt                       float64            `json:"updated_at"`
}
