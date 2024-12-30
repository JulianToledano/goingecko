package types

type Defi struct {
	Data defiInfo `json:"data"`
}

type defiInfo struct {
	DefiMarketCap        string  `json:"defi_market_cap"`
	EthMarketCap         string  `json:"eth_market_cap"`
	DefiToEthRatio       string  `json:"defi_to_eth_ratio"`
	TradingVolume24h     string  `json:"trading_volume_24h"`
	Defi_dominance       string  `json:"defi_dominance"`
	TopCoinName          string  `json:"top_coin_name"`
	TopCoinDefiDominance float64 `json:"top_coin_defi_dominance"`
}
