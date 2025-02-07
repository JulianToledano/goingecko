package types

type Derivative struct {
	Market                   string  `json:"market"`
	Symbol                   string  `json:"symbol"`
	IndexId                  string  `json:"index_id"`
	Price                    string  `json:"price"`
	PricePercentageChange24h float64 `json:"price_percentage_change_24h"`
	ContractType             string  `json:"contract_type"`
	Index                    float64 `json:"index"`
	Basis                    float64 `json:"basis"`
	Spread                   float64 `json:"spread"`
	FundingRate              float64 `json:"funding_rate"`
	OpenInterest             float64 `json:"open_interest"`
	Volume24h                float64 `json:"volume_24h"`
	LastTradedAt             int64   `json:"last_traded_at"`
	ExpiredAt                int64   `json:"expired_at"`
}

type Exchange struct {
	Name                   string  `json:"name"`
	ID                     string  `json:"id"`
	OpenInterestBtc        float64 `json:"open_interest_btc"`
	TradeVolume24hBtc      string  `json:"trade_volume_24h_btc"`
	NumberOfPerpetualPairs int64   `json:"number_of_perpetual_pairs"`
	NumberOfFuturesPairs   int64   `json:"number_of_futures_pairs"`
	Image                  string  `json:"image"`
	YearEstablished        int64   `json:"year_established"`
	Country                string  `json:"country"`
	Description            string  `json:"description"`
	Url                    string  `json:"url"`
}

type ExchangeId struct {
	Name                   string   `json:"name"`
	OpenInterestBtc        float64  `json:"open_interest_btc"`
	TradeVolume24hBtc      string   `json:"trade_volume_24h_btc"`
	NumberOfPerpetualPairs int64    `json:"number_of_perpetual_pairs"`
	NumberOfPuturesPairs   int64    `json:"number_of_futures_pairs"`
	Image                  string   `json:"image"`
	YearEstablished        int64    `json:"year_established"`
	Country                string   `json:"country"`
	Description            string   `json:"description"`
	Url                    string   `json:"url"`
	Tickers                []Ticker `json:"tickers"`
}

type Ticker struct {
	Symbol               string    `json:"symbol"`
	Base                 string    `json:"base"`
	Target               string    `json:"target"`
	TradeURL             string    `json:"trade_url"`
	ContractType         string    `json:"contract_type"`
	Last                 float64   `json:"last"`
	H24PercentageChange  float64   `json:"h24_percentage_change"`
	Index                float64   `json:"index"`
	IndexBasicPercentage float64   `json:"index_basic_percentage"`
	BidAskSpread         float64   `json:"bid_ask_spread"`
	FundingRate          float64   `json:"funding_rate"`
	OpenInterestUsd      float64   `json:"open_interest_usd"`
	H24Volume            float64   `json:"h24_volume"`
	ConvertedVolume      Converted `json:"converted_volume"`
	ConvertedLast        Converted `json:"converted_last"`
	LastTradedAt         string    `json:"last_traded_at"`
	ExpiredAt            float64   `json:"expired_at"`
}

type Converted struct {
	Btc string `json:"btc"`
	Eth string `json:"eth"`
	Usd string `json:"usd"`
}

type DerivativesListItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
