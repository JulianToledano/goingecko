package coins

type Tickers struct {
	Name    string   `json:"name"`
	Tickers []Ticker `json:"tickers"`
}
