package types

import "github.com/JulianToledano/goingecko/types"

type History struct {
	ID                  string                    `json:"id"`
	Symbol              string                    `json:"symbol"`
	Name                string                    `json:"name"`
	Localization        types.Localization        `json:"localization"`
	Image               types.Image               `json:"image"`
	MarketData          HistoryMarketData         `json:"market_data"`
	CommunityData       types.CommunityData       `json:"community_data"`
	DeveloperData       types.DeveloperData       `json:"developer_data"`
	PublicInterestStats types.PublicInterestStats `json:"public_interest_stats"`
}

type HistoryMarketData struct {
	CurrentPrice types.PriceRates `json:"current_price"`
	MarketCap    types.PriceRates `json:"market_cap"`
	TotalVolume  types.PriceRates `json:"total_volume"`
}
