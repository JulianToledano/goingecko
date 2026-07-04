package types

type MarketCapChart struct {
	MarketCapChart MarketCapChartData `json:"market_cap_chart"`
}

type MarketCapChartData struct {
	MarketCap [][]float64 `json:"market_cap"`
	Volume    [][]float64 `json:"volume"`
}
