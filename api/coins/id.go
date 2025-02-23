package coins

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/JulianToledano/goingecko/v3/api/coins/types"
	"github.com/JulianToledano/goingecko/v3/api/internal"
)

// coinsIdOption is specific to the CoinsId function
type coinsIdOption interface {
	internal.Option
	isCoinsIdOption()
}

// WithLocalization includes localized data in the response if true.
// Default: true
func WithLocalization(localization bool) coinsIdOption { return localizationOption{localization} }

// WithTickers includes tickers data in the response if true.
// Default: true
func WithTickers(tickers bool) coinsIdOption { return tickersOption{tickers} }

// WithMarketData includes market data in the response if true.
// Default: true
func WithMarketData(marketData bool) coinsIdOption { return marketDataOption{marketData} }

// WithCommunityData includes community data in the response if true.
// Default: true
func WithCommunityData(communityData bool) coinsIdOption {
	return communityDataOption{communityData}
}

// WithDeveloperData includes developer data in the response if true.
// Default: true
func WithDeveloperData(developerData bool) coinsIdOption {
	return developerDataOption{developerData}
}

// WithCoinSparkline includes sparkline data in the response if true.
// Default: false
func WithCoinSparkline(sparkline bool) coinsIdOption { return coinSparklineOption{sparkline} }

// CoinsId allows you to query all the coin data of a coin (name, price, market .... including exchange tickers) on
// CoinGecko coin page based on a particular coin id.
//
// 👍 Tips
//
// You may obtain the coin id (api id) via several ways:
// refers to respective coin page and find ‘api id’
// refers to /coins/list endpoint
// refers to google sheets here
// You may also flag to include more data such as tickers, market data, community data, developer data and sparkline
// You may refer to last_updated in the endpoint response to check whether the price is stale
//
// 📘 Notes
//
// Tickers are limited to 100 items, to get more tickers, please go to /coins/{id}/tickers
// Cache/Update Frequency:
// Every 60 seconds for all the API plans
// Community data for Twitter and Telegram will be updated on weekly basis (Reddit community data is no longer supported)
func (c *CoinsClient) CoinsId(ctx context.Context, id string, options ...coinsIdOption) (*types.CoinID, error) {
	params := url.Values{}

	// Apply all the options
	for _, opt := range options {
		opt.Apply(&params)
	}

	rUrl := fmt.Sprintf("%s/%s?%s", c.coinsUrl(), id, params.Encode())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data *types.CoinID
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Define option types
type (
	localizationOption  struct{ localization bool }
	tickersOption       struct{ tickers bool }
	marketDataOption    struct{ marketData bool }
	communityDataOption struct{ communityData bool }
	developerDataOption struct{ developerData bool }
	coinSparklineOption struct{ sparkline bool }
)

// Implement Option interface
func (o localizationOption) Apply(v *url.Values) {
	v.Set("localization", strconv.FormatBool(o.localization))
}
func (o tickersOption) Apply(v *url.Values) { v.Set("tickers", strconv.FormatBool(o.tickers)) }
func (o marketDataOption) Apply(v *url.Values) {
	v.Set("market_data", strconv.FormatBool(o.marketData))
}

func (o communityDataOption) Apply(v *url.Values) {
	v.Set("community_data", strconv.FormatBool(o.communityData))
}

func (o developerDataOption) Apply(v *url.Values) {
	v.Set("developer_data", strconv.FormatBool(o.developerData))
}

func (o coinSparklineOption) Apply(v *url.Values) {
	v.Set("sparkline", strconv.FormatBool(o.sparkline))
}

// Implement CoinsIdOption interface
func (localizationOption) isCoinsIdOption()  {}
func (tickersOption) isCoinsIdOption()       {}
func (marketDataOption) isCoinsIdOption()    {}
func (communityDataOption) isCoinsIdOption() {}
func (developerDataOption) isCoinsIdOption() {}
func (coinSparklineOption) isCoinsIdOption() {}
