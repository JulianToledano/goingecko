package types

import "github.com/JulianToledano/goingecko/v3/api/types"

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
	Localization              types.Localization        `json:"localization"`
	Description               types.Description         `json:"description"`
	Links                     types.Links               `json:"links"`
	Image                     types.Image               `json:"image"`
	CountryOrigin             string                    `json:"country_origin"`
	GenesisData               string                    `json:"genesis_date"`
	SentimentVotesUpPercent   float32                   `json:"sentiment_votes_up_percentage"`
	SentimentVotesDownPercent float32                   `json:"sentiment_votes_down_percentage"`
	MarketCapRank             int                       `json:"market_cap_rank"`
	CoingeckoRank             int                       `json:"coingecko_rank"`
	CoingeckoScore            float32                   `json:"coingecko_score"`
	DeveloperScore            float32                   `json:"developer_score"`
	CommunityScore            float32                   `json:"community_score"`
	LiquidityScore            float32                   `json:"liquidity_score"`
	PublicInterestScore       float32                   `json:"public_interest_score"`
	MarketData                types.MarketData          `json:"market_data"`
	CommunityData             types.CommunityData       `json:"community_data"`
	DeveloperData             types.DeveloperData       `json:"developer_data"`
	PublicInterestStats       types.PublicInterestStats `json:"public_interest_stats"`
	Tickers                   []types.Ticker            `json:"tickers"`
}

type ReposUrl struct {
	Github    []string `json:"github"`
	Bitbucket []string `json:"bitbucket"`
}
