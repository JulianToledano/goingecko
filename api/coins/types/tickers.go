package types

import "github.com/JulianToledano/goingecko/api/types"

type Tickers struct {
	Name    string         `json:"name"`
	Tickers []types.Ticker `json:"tickers"`
}
