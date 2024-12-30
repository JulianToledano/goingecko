package types

type Treasury struct {
	TotalHoldings      float64     `json:"total_holdings"`
	TotalValueUsd      float64     `json:"total_value_usd"`
	MarketCapDominance float64     `json:"market_cap_dominance"`
	Companies          []Companies `json:"companies"`
}

type Companies struct {
	Name                    string  `json:"name"`
	Symbol                  string  `json:"symbol"`
	Country                 string  `json:"country"`
	TotalHoldings           int64   `json:"total_holdings"`
	TotalEntryValueUsd      int64   `json:"total_entry_value_usd"`
	TotalCurrentValueUsd    int64   `json:"total_current_value_usd"`
	PercentageOfTotalSupply float64 `json:"percentage_of_total_supply"`
}
