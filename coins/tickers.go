package coins

import "github.com/JulianToledano/goingecko/types"

type Tickers struct {
	Name    string         `json:"name"`
	Tickers []types.Ticker `json:"tickers"`
}
