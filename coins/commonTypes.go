package coins

type Roi struct {
	Time       float64 `json:"times"`
	Currency   string  `json:"currency"`
	Percentage float64 `json:"percentage"`
}

type Sparkline struct {
	Price []float64 `json:"price"`
}

type Ticker struct {
	Base                   string       `json:"base"`
	Target                 string       `json:"target"`
	Market                 MarketTicker `json:"market"`
	Last                   float64      `json:"last"`
	Volume                 float64      `json:"volume"`
	ConvertedLast          Converted    `json:"converted_last"`
	ConvertedVolume        Converted    `json:"converted_volume"`
	TrustScore             string       `json:"trust_score"`
	BidAskSpreadPercentage float64      `json:"bid_ask_spread_percentage"`
	Timestamp              string       `json:"timestamp"`
	LastTradedAt           string       `json:"last_traded_at"`
	LastFetchAt            string       `json:"last_fetch_at"`
	IsAnomaly              bool         `json:"is_anomaly"`
	IsStale                bool         `json:"is_stale"`
	TradeURL               string       `json:"trade_url"`
	// token_info_url `json:"token_info_url"`	null
	CoinID string `json:"coin_id"`
}

type MarketTicker struct {
	Name                string `json:"name"`
	Identifier          string `json:"identifier"`
	HasTradingIncentive bool   `json:"has_trading_incentive"`
}

type Converted struct {
	Btc float64 `json:"btc"`
	Eth float64 `json:"eth"`
	Usd float64 `json:"usd"`
}

type Image struct {
	Thumb string `json:"thumb"`
	Small string `json:"small"`
	Large string `json:"large"`
}

type Localization struct {
	En  string `json:"en"`
	De  string `json:"de"`
	Es  string `json:"es"`
	Fr  string `json:"fr"`
	It  string `json:"it"`
	Pl  string `json:"pl"`
	Ro  string `json:"ro"`
	Hu  string `json:"hu"`
	Nl  string `json:"nl"`
	Pt  string `json:"pt"`
	Sv  string `json:"sv"`
	Vi  string `json:"vi"`
	Tr  string `json:"tr"`
	Ru  string `json:"ru"`
	Ja  string `json:"ja"`
	Zh  string `json:"zh"`
	Zhw string `json:"zhw"`
	Ko  string `json:"ko"`
	Ar  string `json:"ar"`
	Th  string `json:"th"`
	ID  string `json:"id"`
}

type CommunityData struct {
	FacebookLikes                float64 `json:"facebook_likes"`
	TwitterFollowers             float64 `json:"twitter_followers"`
	RedditAveragePosts           float64 `json:"reddit_average_posts_48h"`
	RedditAverageCommentsTwoDays float64 `json:"reddit_average_comments_48h"`
	RedditSubscribers            float64 `json:"reddit_subscribers"`
	RedditAccountsActiveTwoDays  float64 `json:"reddit_accounts_active_48h"`
	TelegramChannelUserCount     float64 `json:"telegram_channel_user_count"`
}

type DeveloperData struct {
	Forks                       float32   `json:"forks"`
	Stars                       float32   `json:"stars"`
	Subscribers                 float32   `json:"subscribers"`
	TotalIssues                 float32   `json:"total_issues"`
	ClosedIssues                float32   `json:"closed_issues"`
	PullRequestsMerged          float32   `json:"pull_requests_merged"`
	PullRequestContributors     float32   `json:"pull_request_contributors"`
	CodeAdditionsDeletionsMonth CodeStats `json:"code_additions_deletions_4_weeks"`
	CommitCountMonth            float32   `json:"commit_count_4_weeks"`
	MonthCommitActivitySeries   []float32 `json:"last_4_weeks_commit_activity_series"`
}

type CodeStats struct {
	Additions int16 `json:"additions"`
	Deletions int16 `json:"deletions"`
}

type PublicInterestStats struct {
	AlexaRank float64 `json:"alexa_rank"`
	// bing_matches `json:"bing_matches"`
}

type PriceRates struct {
	Aed  float64 `json:"aed"`
	Ars  float64 `json:"ars"`
	Aud  float64 `json:"aud"`
	Bch  float64 `json:"bch"`
	Bdt  float64 `json:"bdt"`
	Bhd  float64 `json:"bhd"`
	Bmd  float64 `json:"bmd"`
	Bnb  float64 `json:"bnb"`
	Brl  float64 `json:"brl"`
	Btc  float64 `json:"btc"`
	Cad  float64 `json:"cad"`
	Chf  float64 `json:"chf"`
	Clp  float64 `json:"clp"`
	Cny  float64 `json:"cny"`
	Czk  float64 `json:"czk"`
	Dkk  float64 `json:"dkk"`
	Dot  float64 `json:"dot"`
	Eos  float64 `json:"eos"`
	Eth  float64 `json:"eth"`
	Eur  float64 `json:"eur"`
	Gbp  float64 `json:"gbp"`
	Hkd  float64 `json:"hkd"`
	Huf  float64 `json:"huf"`
	Idr  float64 `json:"idr"`
	Ils  float64 `json:"ils"`
	Inr  float64 `json:"inr"`
	Jpy  float64 `json:"jpy"`
	Krw  float64 `json:"krw"`
	Kwd  float64 `json:"kwd"`
	Lkr  float64 `json:"lkr"`
	Ltc  float64 `json:"ltc"`
	Mmk  float64 `json:"mmk"`
	Mxn  float64 `json:"mxn"`
	Myr  float64 `json:"myr"`
	Ngn  float64 `json:"ngn"`
	Nok  float64 `json:"nok"`
	Nzd  float64 `json:"nzd"`
	Php  float64 `json:"php"`
	Pkr  float64 `json:"pkr"`
	Pln  float64 `json:"pln"`
	Rub  float64 `json:"rub"`
	Sar  float64 `json:"sar"`
	Sek  float64 `json:"sek"`
	Sgd  float64 `json:"sgd"`
	Thb  float64 `json:"thb"`
	Try  float64 `json:"try"`
	Twd  float64 `json:"twd"`
	Uah  float64 `json:"uah"`
	Usd  float64 `json:"usd"`
	Vef  float64 `json:"vef"`
	Vnd  float64 `json:"vnd"`
	Xag  float64 `json:"xag"`
	Xau  float64 `json:"xau"`
	Xdr  float64 `json:"xdr"`
	Xlm  float64 `json:"xlm"`
	Xrp  float64 `json:"xrp"`
	Yfi  float64 `json:"yfi"`
	Zar  float64 `json:"zar"`
	Bits float64 `json:"bits"`
	Link float64 `json:"link"`
	Sats float64 `json:"sats"`
}

type DateRates struct {
	Aed  string `json:"aed"`
	Ars  string `json:"ars"`
	Aud  string `json:"aud"`
	Bch  string `json:"bch"`
	Bdt  string `json:"bdt"`
	Bhd  string `json:"bhd"`
	Bmd  string `json:"bmd"`
	Bnb  string `json:"bnb"`
	Brl  string `json:"brl"`
	Btc  string `json:"btc"`
	Cad  string `json:"cad"`
	Chf  string `json:"chf"`
	Clp  string `json:"clp"`
	Cny  string `json:"cny"`
	Czk  string `json:"czk"`
	Dkk  string `json:"dkk"`
	Dot  string `json:"dot"`
	Eos  string `json:"eos"`
	Eth  string `json:"eth"`
	Eur  string `json:"eur"`
	Gbp  string `json:"gbp"`
	Hkd  string `json:"hkd"`
	Huf  string `json:"huf"`
	Idr  string `json:"idr"`
	Ils  string `json:"ils"`
	Inr  string `json:"inr"`
	Jpy  string `json:"jpy"`
	Krw  string `json:"krw"`
	Kwd  string `json:"kwd"`
	Lkr  string `json:"lkr"`
	Ltc  string `json:"ltc"`
	Mmk  string `json:"mmk"`
	Mxn  string `json:"mxn"`
	Myr  string `json:"myr"`
	Ngn  string `json:"ngn"`
	Nok  string `json:"nok"`
	Nzd  string `json:"nzd"`
	Php  string `json:"php"`
	Pkr  string `json:"pkr"`
	Pln  string `json:"pln"`
	Rub  string `json:"rub"`
	Sar  string `json:"sar"`
	Sek  string `json:"sek"`
	Sgd  string `json:"sgd"`
	Thb  string `json:"thb"`
	Try  string `json:"try"`
	Twd  string `json:"twd"`
	Uah  string `json:"uah"`
	Usd  string `json:"usd"`
	Vef  string `json:"vef"`
	Vnd  string `json:"vnd"`
	Xag  string `json:"xag"`
	Xau  string `json:"xau"`
	Xdr  string `json:"xdr"`
	Xlm  string `json:"xlm"`
	Xrp  string `json:"xrp"`
	Yfi  string `json:"yfi"`
	Zar  string `json:"zar"`
	Bits string `json:"bits"`
	Link string `json:"link"`
	Sats string `json:"sats"`
}

type CapRates struct {
	Aed  int64 `json:"aed"`
	Ars  int64 `json:"ars"`
	Aud  int64 `json:"aud"`
	Bch  int64 `json:"bch"`
	Bdt  int64 `json:"bdt"`
	Bhd  int64 `json:"bhd"`
	Bmd  int64 `json:"bmd"`
	Bnb  int64 `json:"bnb"`
	Brl  int64 `json:"brl"`
	Btc  int64 `json:"btc"`
	Cad  int64 `json:"cad"`
	Chf  int64 `json:"chf"`
	Clp  int64 `json:"clp"`
	Cny  int64 `json:"cny"`
	Czk  int64 `json:"czk"`
	Dkk  int64 `json:"dkk"`
	Dot  int64 `json:"dot"`
	Eos  int64 `json:"eos"`
	Eth  int64 `json:"eth"`
	Eur  int64 `json:"eur"`
	Gbp  int64 `json:"gbp"`
	Hkd  int64 `json:"hkd"`
	Huf  int64 `json:"huf"`
	Idr  int64 `json:"idr"`
	Ils  int64 `json:"ils"`
	Inr  int64 `json:"inr"`
	Jpy  int64 `json:"jpy"`
	Krw  int64 `json:"krw"`
	Kwd  int64 `json:"kwd"`
	Lkr  int64 `json:"lkr"`
	Ltc  int64 `json:"ltc"`
	Mmk  int64 `json:"mmk"`
	Mxn  int64 `json:"mxn"`
	Myr  int64 `json:"myr"`
	Ngn  int64 `json:"ngn"`
	Nok  int64 `json:"nok"`
	Nzd  int64 `json:"nzd"`
	Php  int64 `json:"php"`
	Pkr  int64 `json:"pkr"`
	Pln  int64 `json:"pln"`
	Rub  int64 `json:"rub"`
	Sar  int64 `json:"sar"`
	Sek  int64 `json:"sek"`
	Sgd  int64 `json:"sgd"`
	Thb  int64 `json:"thb"`
	Try  int64 `json:"try"`
	Twd  int64 `json:"twd"`
	Uah  int64 `json:"uah"`
	Usd  int64 `json:"usd"`
	Vef  int64 `json:"vef"`
	Vnd  int64 `json:"vnd"`
	Xag  int64 `json:"xag"`
	Xau  int64 `json:"xau"`
	Xdr  int64 `json:"xdr"`
	Xlm  int64 `json:"xlm"`
	Xrp  int64 `json:"xrp"`
	Yfi  int64 `json:"yfi"`
	Zar  int64 `json:"zar"`
	Bits int64 `json:"bits"`
	Link int64 `json:"link"`
	Sats int64 `json:"sats"`
}
