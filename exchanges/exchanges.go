package exchanges

type Exchange struct {
	ID                         string  `json:"id"`
	Name                       string  `json:"name"`
	YearEstablished            int32   `json:"year_established"`
	Country                    string  `json:"country"`
	Description                string  `json:"description"`
	Url                        string  `json:"url"`
	Image                      string  `json:"image"`
	HasTradingIncentive        bool    `json:"has_trading_incentive"`
	TrustScore                 int32   `json:"trust_score"`
	TrustScoreRank             int32   `json:"trust_score_rank"`
	TradeVolume24h             float64 `json:"trade_volume_24h_btc"`
	TradeVolume24jBtcNormalice float64 `json:"trade_volume_24h_btc_normalized"`
}

type ExchangesList []Exchange

type ExchangeId struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ExchangeWithTicker struct {
	Name                        string          `json:"name"`
	YearEstablished             int32           `json:"year_established"`
	Country                     string          `json:"country"`
	Description                 string          `json:"description"`
	Url                         string          `json:"url"`
	Image                       string          `json:"image"`
	FacebookUrl                 string          `json:"facebook_url"`
	RedditYrl                   string          `json:"reddit_url"`
	TelegramUrl                 string          `json:"telegram_url"`
	SlackUrl                    string          `json:"slack_url"`
	OtherUrl1                   string          `json:"other_url_1"`
	OtherUrl2                   string          `json:"other_url_2"`
	TwitterHandle               string          `json:"twitter_handle"`
	HasTradingIncentive         bool            `json:"has_trading_incentive"`
	Centralized                 bool            `json:"centralized"`
	PublicNotice                string          `json:"public_notice"`
	AlertNotice                 string          `json:"alert_notice"`
	TrustScore                  int32           `json:"trust_score"`
	TrustScoreRank              int32           `json:"trust_score_rank"`
	TradeVolume24hBtc           float64         `json:"trade_volume_24h_btc"`
	TradeVolume24hBtcNormalized float64         `json:"trade_volume_24h_btc_normalized"`
	Tickers                     []Ticker        `json:"tickers"`
	StatusUpdates               []StatusUpdates `json:"status_updates"`
}

type Tickers struct {
	Name    string   `json:"name"`
	Tickers []Ticker `json:"tickers"`
}

type Ticker struct {
	Base                   string    `json:"base"`
	Target                 string    `json:"target"`
	Market                 Market    `json:"Market"`
	Last                   float64   `json:"Last"`
	Volume                 float64   `json:"volume"`
	ConvertedLast          Converted `json:"converted_last"`
	ConvertedVolume        Converted `json:"converted_volume"`
	TrustScore             string    `json:"trust_score"`
	BidAskSpreadPercentage float64   `json:"bid_ask_spread_percentage"`
	Timestamp              string    `json:"timestamp"`
	LastTradedAt           string    `json:"last_traded_at"`
	LastFetchAt            string    `json:"last_fetch_at"`
	IsAnomaly              bool      `json:"is_anomaly"`
	IsStale                bool      `json:"is_stale"`
	TradeUrl               string    `json:"trade_url"`
	TokenInfoUrl           string    `json:"token_info_url"`
	CoinID                 string    `json:"coin_id"`
	TargetCoinID           string    `json:"target_coin_id"`
}

type StatusUpdates struct {
	// ??
}

type Market struct {
	Name                string `json:"name"`
	Identifier          string `json:"identifier"`
	HasTradingIncentive bool   `json:"has_trading_incentive"`
}

type Converted struct {
	Btc float64 `json:"btc"`
	Eth float64 `json:"eth"`
	Usd float64 `json:"usd"`
}

var VolumeChartValidDays = map[string]struct{}{
	"1":   {},
	"7":   {},
	"14":  {},
	"30":  {},
	"90":  {},
	"180": {},
	"365": {},
}

type Volume []interface{}
