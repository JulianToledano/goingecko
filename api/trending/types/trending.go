package types

type Trending struct {
	Coins []coin `json:"coins"`
}

type coin struct {
	Item item `json:"item"`
}

type item struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Symbol        string  `json:"symbol"`
	MarketCapRank int64   `json:"market_cap_rank"`
	Thumb         string  `json:"thumb"`
	Large         string  `json:"large"`
	Score         float64 `json:"score"`
}
