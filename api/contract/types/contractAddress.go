package types

import "github.com/JulianToledano/goingecko/v3/api/types"

// TODO: missing a lot of info
type ContractAddressInfo struct {
	ID                 string            `json:"id"`
	Symbol             string            `json:"symbol"`
	Name               string            `json:"name"`
	AssetPlatformId    string            `json:"asset_platform_id"`
	Platforms          map[string]string `json:"platforms"`
	BlockTimeInMinutes int64             `json:"block_time_in_minutes"`
	Hashing_algorithm  string            `json:"hashing_algorithm"`
	Categories         []string          `json:"categories"`
	// public_notice `json:"//`"
	// additional_notices [] `json:"//`"
	Localization                    types.Localization        `json:"localization"`
	Jescription                     types.Description         `json:"description"`
	Links                           types.Links               `json:"links"`
	Image                           types.Image               `json:"image"`
	CountryOrigin                   string                    `json:"country_origin"`
	GenesisDate                     string                    `json:"genesis_date"`
	Contract_address                string                    `json:"contract_address"`
	Sentiment_votes_up_percentage   float64                   `json:"sentiment_votes_up_percentage"`
	Sentiment_votes_down_percentage float64                   `json:"sentiment_votes_down_percentage"`
	IcoData                         IcoData                   `json:"ico_data"`
	MarketCapRank                   int64                     `json:"market_cap_rank"`
	CoingeckoRank                   int64                     `json:"coingecko_rank"`
	CoingeckoScore                  float64                   `json:"coingecko_score"`
	DeveloperScore                  float64                   `json:"developer_score"`
	CommunityScore                  float64                   `json:"community_score"`
	LiquidityScore                  float64                   `json:"liquidity_score"`
	PublicInterestScore             float64                   `json:"public_interest_score"`
	MarketData                      types.MarketData          `json:"market_data"`
	CommunityData                   types.CommunityData       `json:"community_data"`
	DeveloperData                   types.DeveloperData       `json:"developer_data"`
	PublicInterestStats             types.PublicInterestStats `json:"public_interest_stats"`
	StatusUpdates                   []string                  `json:"status_updates"`
	LastUpdated                     string                    `json:"last_updated"`
	Tickers                         []types.Ticker            `json:"tickers"`
}

type IcoData struct {
	IcoStartDate        string `json:"ico_start_date"`
	IcoEndDate          string `json:"ico_end_date"`
	ShortDesc           string `json:"short_desc"`
	Description         string `json:"description"`
	Links               Links  `json:"Linkslinks"`
	SoftcapCurrency     string `json:"softcap_currency"`
	HardcapCurrency     string `json:"hardcap_currency"`
	TotalRaisedCurrency string `json:"total_raised_currency"`
	// softcap_amount	`json:"nullsoftcap_amount"`
	// hardcap_amount	`json:"nullhardcap_amount"`
	// total_raised	`json:"nulltotal_raised"`
	QuotePreSaleCurrency string `json:"quote_pre_sale_currency"`
	// base_pre_sale_amount	`json:"nullbase_pre_sale_amount"`
	// quote_pre_sale_amount	`json:"nullquote_pre_sale_amount"`
	QuotePublicSaleCurrency string  `json:"quote_public_sale_currency"`
	BasePublicSaleAmount    float64 `json:"base_public_sale_amount"`
	QuotePublicSaleAmount   float64 `json:"quote_public_sale_amount"`
	AcceptingCurrencies     string  `json:"accepting_currencies"`
	CountryOrigin           string  `json:"country_origin"`
	// pre_sale_start_date	`json:"nullpre_sale_start_date"`
	// pre_sale_end_date	`json:"nullpre_sale_end_date"`
	// bounty_detail_url
	// amount_for_sale
	KycRequired bool `json:"kyc_required"`
	// whitelist_available
	// pre_sale_available
	PreSaleEnded bool `json:"pre_sale_ended"`
}

type Links struct {
	Web        string `json:"web"`
	Blog       string `json:"blog"`
	Github     string `json:"github"`
	Twitter    string `json:"twitter"`
	Facebook   string `json:"facebook"`
	Telegram   string `json:"telegram"`
	Whitepaper string `json:"whitepaper"`
}
