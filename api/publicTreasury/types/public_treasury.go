package types

type Entity struct {
	ID      string `json:"id"`
	Symbol  string `json:"symbol"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

type Treasury struct {
	TotalHoldings      float64           `json:"total_holdings"`
	TotalValueUsd      float64           `json:"total_value_usd"`
	MarketCapDominance float64           `json:"market_cap_dominance"`
	Companies          []TreasuryHolding `json:"companies"`
	Governments        []TreasuryHolding `json:"governments"`
}

type TreasuryHolding struct {
	Name                    string  `json:"name"`
	Symbol                  string  `json:"symbol"`
	Country                 string  `json:"country"`
	TotalHoldings           float64 `json:"total_holdings"`
	TotalEntryValueUsd      float64 `json:"total_entry_value_usd"`
	TotalCurrentValueUsd    float64 `json:"total_current_value_usd"`
	PercentageOfTotalSupply float64 `json:"percentage_of_total_supply"`
}

type EntityTreasury struct {
	Name                       string                  `json:"name"`
	ID                         string                  `json:"id"`
	Type                       string                  `json:"type"`
	Symbol                     *string                 `json:"symbol"`
	Country                    string                  `json:"country"`
	WebsiteURL                 string                  `json:"website_url"`
	TwitterScreenName          string                  `json:"twitter_screen_name"`
	TotalTreasuryValueUsd      float64                 `json:"total_treasury_value_usd"`
	UnrealizedPnl              float64                 `json:"unrealized_pnl"`
	MNav                       float64                 `json:"m_nav"`
	TotalAssetValuePerShareUsd float64                 `json:"total_asset_value_per_share_usd"`
	Holdings                   []EntityTreasuryHolding `json:"holdings"`
}

type EntityTreasuryHolding struct {
	CoinID                   string             `json:"coin_id"`
	Amount                   float64            `json:"amount"`
	PercentageOfTotalSupply  float64            `json:"percentage_of_total_supply"`
	AmountPerShare           float64            `json:"amount_per_share"`
	EntityValueUsdPercentage float64            `json:"entity_value_usd_percentage"`
	CurrentValueUsd          float64            `json:"current_value_usd"`
	TotalEntryValueUsd       float64            `json:"total_entry_value_usd"`
	AverageEntryValueUsd     float64            `json:"average_entry_value_usd"`
	UnrealizedPnl            float64            `json:"unrealized_pnl"`
	HoldingAmountChange      map[string]float64 `json:"holding_amount_change"`
	HoldingChangePercentage  map[string]float64 `json:"holding_change_percentage"`
}

type HoldingChart struct {
	Holdings          [][]float64 `json:"holdings"`
	HoldingValueInUsd [][]float64 `json:"holding_value_in_usd"`
}

type TransactionHistory struct {
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	Date                 int64   `json:"date"`
	SourceURL            string  `json:"source_url"`
	CoinID               string  `json:"coin_id"`
	Type                 string  `json:"type"`
	HoldingNetChange     float64 `json:"holding_net_change"`
	TransactionValueUsd  float64 `json:"transaction_value_usd"`
	HoldingBalance       float64 `json:"holding_balance"`
	AverageEntryValueUsd float64 `json:"average_entry_value_usd"`
}
