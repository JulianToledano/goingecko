package coins

type History struct {
	ID                  string              `json:"id"`
	Symbol              string              `json:"symbol"`
	Name                string              `json:"name"`
	Localization        Localization        `json:"localization"`
	Image               Image               `json:"image"`
	MarketData          HistoryMarketData   `json:"market_data"`
	CommunityData       CommunityData       `json:"community_data"`
	DeveloperData       DeveloperData       `json:"developer_data"`
	PublicInterestStats PublicInterestStats `json:"public_interest_stats"`
}

type HistoryMarketData struct {
	CurrentPrice PriceRates `json:"current_price"`
	MarketCap    PriceRates `json:"market_cap"`
	TotalVolume  PriceRates `json:"total_volume"`
}
