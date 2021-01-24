package coins

type Roi struct {
	Time       float64 `json:"times"`
	Currency   string  `json:"currency"`
	Percentage float64 `json:"percentage"`
}

type Sparkline struct {
	Price []float64 `json:"price"`
}
