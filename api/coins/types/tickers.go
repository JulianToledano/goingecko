package types

import "github.com/JulianToledano/goingecko/v3/api/types"

type Tickers struct {
	Name    string         `json:"name"`
	Tickers []types.Ticker `json:"tickers"`
}
