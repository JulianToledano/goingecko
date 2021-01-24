package coins

type Market struct {
	ID                                         string    `json:"ID"`
	Symbol                                     string    `json:"symbol"`
	Name                                       string    `json:"name"`
	Image                                      string    `json:"image"`
	CurrentPrice                               float64   `json:"current_price"`
	MarketCap                                  float64   `json:"market_cap"`
	MarkeCcapRank                              int32     `json:"market_cap_rank"`
	FullyDilutedValuation                      float64   `json:"fully_diluted_valuation"`
	TotalVolume                                float64   `json:"total_volume"`
	HighDay                                    float64   `json:"high_24h"`
	LowDay                                     float64   `json:"low_24h"`
	PriceChangeDay                             float64   `json:"price_change_24h"`
	PriceChangePercentageDay                   float64   `json:"price_change_percentage_24h"`
	MarketCapChangeDay                         float64   `json:"market_cap_change_24h"`
	MarketCapChangePercentageDay               float64   `json:"market_cap_change_percentage_24h"`
	CirculatingSupply                          float64   `json:"circulating_supply"`
	TotalSupply                                float64   `json:"total_supply"`
	MaxSupply                                  float64   `json:"max_supply"`
	Ath                                        float64   `json:"ath"`
	AthChangePercentage                        float64   `json:"ath_change_percentage"`
	AthDate                                    string    `json:"ath_date"`
	Atl                                        float64   `json:"atl"`
	AtlChangePercentage                        float64   `json:"atl_change_percentage"`
	AtlDate                                    string    `json:"atl_date"`
	Roi                                        Roi       `json:"roi"`
	LastUpdated                                string    `json:"last_updated"`
	SparklineInWeek                            Sparkline `json:"sparkline_in_7d"`
	PriceChangePercentageTwoWeeksInCurrency    float64   `json:"price_change_percentage_14d_in_currency"`
	PriceChangePercentageHourInCurrency        float64   `json:"price_change_percentage_1h_in_currency"`
	PriceChangePercentageThreeMonthsInCurrency float64   `json:"price_change_percentage_200d_in_currency"`
	PriceChangePercentageDayInCurrency         float64   `json:"price_change_percentage_24h_in_currency"`
	PriceChangePercentageMonthInCurrency       float64   `json:"price_change_percentage_30d_in_currency"`
	PriceChangePercentageWeekInCurrency        float64   `json:"price_change_percentage_7d_in_currency"`
}
