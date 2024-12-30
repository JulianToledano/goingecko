package coins

import "github.com/JulianToledano/goingecko/v2/types"

type Tickers struct {
	Name    string         `json:"name"`
	Tickers []types.Ticker `json:"tickers"`
}
