package types

type Search struct {
	Coins []Item `json:"coins"`
}

type Item struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	ApiSymbol     string `json:"api_symbol"`
	Symbol        string `json:"symbol"`
	MarketCapRank int64  `json:"market_cap_rank"`
	Thumb         string `json:"thumb"`
	Large         string `json:"large"`
}
