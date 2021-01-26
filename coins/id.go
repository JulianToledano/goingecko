package coins

type CoinID struct {
	ID                 string   `json:"id"`
	Symbol             string   `json:"symbol"`
	Name               string   `json:"name"`
	AssetPlatformID    string   `json:"asset_platform_id"`
	BlockTimeInMinutes int      `json:"block_time_in_minutes"`
	HashingAlgorithm   string   `json:"hashing_algorithm"`
	Categories         []string `json:"categories"`
	// PublicNotice ¿? `json:"public_notice"`
	// AdditionalNotices ¿? `json:"public_notices"`
	Localization              Localization        `json:"localization"`
	Description               Description         `json:"description"`
	Links                     Links               `json:"links"`
	Image                     Image               `json:"image"`
	CountryOrigin             string              `json:"country_origin"`
	GenesisData               string              `json:"genesis_date"`
	SentimentVotesUpPercent   float32             `json:"sentiment_votes_up_percentage"`
	SentimentVotesDownPercent float32             `json:"sentiment_votes_down_percentage"`
	MarketCapRank             int                 `json:"market_cap_rank"`
	CoingeckoRank             int                 `json:"coingecko_rank"`
	CoingeckoScore            float32             `json:"coingecko_score"`
	DeveloperScore            float32             `json:"developer_score"`
	CommunityScore            float32             `json:"community_score"`
	LiquidityScore            float32             `json:"liquidity_score"`
	PublicInterestScore       float32             `json:"public_interest_score"`
	MarketData                MarketData          `json:"market_data"`
	CommunityData             CommunityData       `json:"community_data"`
	DeveloperData             DeveloperData       `json:"developer_data"`
	PublicInterestStats       PublicInterestStats `json:"public_interest_stats"`
	Tickers                   []Ticker            `json:"tickers"`
}

type Description struct {
	En   string `json:"en"`
	De   string `json:"de"`
	Es   string `json:"es"`
	Fr   string `json:"fr"`
	It   string `json:"it"`
	Pl   string `json:"pl"`
	Ro   string `json:"ro"`
	Hu   string `json:"hu"`
	Nl   string `json:"nl"`
	Pt   string `json:"pt"`
	Sv   string `json:"sv"`
	Vi   string `json:"vi"`
	Tr   string `json:"tr"`
	Ru   string `json:"ru"`
	Ja   string `json:"ja"`
	Zh   string `json:"zh"`
	Ko   string `json:"ko"`
	Ar   string `json:"ar"`
	Th   string `json:"th"`
	ID   string `json:"id"`
	Zhtw string `json:"zh-tw"`
}

type Links struct {
	Homepage                    []string `json:"homepage"`
	BlockchainSite              []string `json:"blockchain_site"`
	OfficialForumURL            []string `json:"official_forum_url"`
	ChatURL                     []string `json:"chat_url"`
	AnnouncementURL             []string `json:"announcement_url"`
	TwitterScreenName           string   `json:"twitter_screen_name"`
	FacebookUsername            string   `json:"facebook_username"`
	BitcointalkThreadIdentifier string   `json:"bitcointalk_thread_identifier"`
	TelegramChannelIdentifier   string   `json:"telegram_channel_identifier"`
	SubredditURL                string   `json:"subreddit_url"`
	ReposURL                    ReposUrl `json:"repos_url"`
}

type ReposUrl struct {
	Github    []string `json:"github"`
	Bitbucket []string `json:"bitbucket"`
}

type MarketData struct {
	CurrentPrice                          PriceRates `json:"current_price"`
	Roi                                   Roi        `json:"roi"`
	Ath                                   PriceRates `json:"ath"`
	AthChangePercentage                   PriceRates `json:"ath_change_percentage"`
	AthDate                               DateRates  `json:"ath_date"`
	Atl                                   PriceRates `json:"atl"`
	AtlChangePercentage                   PriceRates `json:"atl_change_percentage"`
	AtlDate                               DateRates  `json:"atl_date"`
	MarketCap                             CapRates   `json:"market_cap"`
	MarketCapRank                         int16      `json:"market_cap_rank"`
	FullyDilutedValuation                 CapRates   `json:"fully_diluted_valuation"`
	TotalVolume                           CapRates   `json:"total_volume"`
	High24                                PriceRates `json:"high_24h"`
	Low24                                 PriceRates `json:"low_24h"`
	PriceChangeDay                        float64    `json:"price_change_24h"`
	PriceChangePercentDay                 float64    `json:"price_change_percentage_24h"`
	PriceChangePercentWeek                float64    `json:"price_change_percentage_7d"`
	PriceChangePercentTwoWeeks            float64    `json:"price_change_percentage_14d"`
	PriceChangePercentMonth               float64    `json:"price_change_percentage_30d"`
	PriceChangePercentTwoMonths           float64    `json:"price_change_percentage_60d"`
	PriceChangePercentHalfYear            float64    `json:"price_change_percentage_200d"`
	PriceChangePercentYear                float64    `json:"price_change_percentage_1y"`
	MarketCapChangeDay                    float64    `json:"market_cap_change_24h"`
	MarketCapChangePercentDay             float64    `json:"market_cap_change_percentage_24h"`
	MarketCapChangeDayRates               PriceRates `json:"price_change_24h_in_currency"`
	PriceChangePercentageHourCurrency     PriceRates `json:"price_change_percentage_1h_in_currency"`
	PriceChangePercentageDayCurrency      PriceRates `json:"price_change_percentage_24h_in_currency"`
	PriceChangePercentageWeekCurrency     PriceRates `json:"price_change_percentage_7d_in_currency"`
	PriceChangePercentageTwoWeekCurrency  PriceRates `json:"price_change_percentage_14d_in_currency"`
	PriceChangePercentageMonthCurrency    PriceRates `json:"price_change_percentage_30d_in_currency"`
	PriceChangePercentageTwoMonthCurrency PriceRates `json:"price_change_percentage_60d_in_currency"`
	PriceChangePercentageHalfYearCurrency PriceRates `json:"price_change_percentage_200d_in_currency"`
	PriceChangePercentageYearCurrency     PriceRates `json:"price_change_percentage_1y_in_currency"`
	MarketCapChangeDayCurrency            PriceRates `json:"market_cap_change_24h_in_currency"`
	MarketCapChangePercentageDayCurrency  PriceRates `json:"market_cap_change_percentage_24h_in_currency"`
	TotalSupply                           float64    `json:"total_supply"`
	MaxSupply                             float64    `json:"max_supply"`
	CirculatingSupply                     float64    `json:"circulating_supply"`
	Sparkline                             Sparkline  `json:""sparkline_7d`
	LastUpdated                           string     `json:"last_updated"`
}
