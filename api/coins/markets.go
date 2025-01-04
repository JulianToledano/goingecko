package coins

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/JulianToledano/goingecko/v3/api/coins/types"
	"github.com/JulianToledano/goingecko/v3/api/internal"
)

// marketsOption is an interface that extends internal.Option to provide
// specific options for the CoinsMarket endpoint. It includes a marker
// method isCoinsMarketOption() to ensure type safety for market-specific options.
type marketsOption interface {
	internal.Option
	isMarketsOption()
}

// WithIDs specifies the coin ids to filter results by.
// Multiple ids can be provided as a string slice.
func WithIDs(ids []string) marketsOption { return idsOption{ids} }

// WithCategory filters results by coin category.
// Valid values include: "decentralized_finance_defi", "stablecoins", etc.
func WithCategory(category string) marketsOption { return categoryOption{category} }

// WithOrder specifies the ordering of results.
// Valid values: "market_cap_desc", "market_cap_asc", "volume_desc", "volume_asc",
// "id_desc", "id_asc", "gecko_desc", "gecko_asc"
func WithOrder(order string) marketsOption { return orderOption{order} }

// WithPerPage sets the number of results per page.
// Valid values: 1-250, default: 100
func WithPerPage(perPage int) marketsOption { return perPageOption{perPage} }

// WithPage specifies which page of results to return.
// Default: 1
func WithPage(page int) marketsOption { return pageOption{page} }

// WithSparkline includes sparkline data in results if true.
// Default: false
func WithSparkline(sparkline bool) marketsOption { return sparklineOption{sparkline} }

// WithPriceChangePercentage includes price change percentage for specified intervals.
// Valid intervals: "1h", "24h", "7d", "14d", "30d", "200d", "1y"
func WithPriceChangePercentage(intervals []string) marketsOption {
	return priceChangePercentageOption{intervals}
}

// CoinsMarket allows you to query all the supported coins with price, market cap, volume and market related data.
// 👍 Tips
// You may specify the coins’ ids in ids parameter if you want to retrieve market data for specific coins only instead of the whole list
// You may also provide value in category to filter the responses based on coin's category
// You can use per_page and page values to control the number of results per page and specify which page of results you want to display in the responses
//
// 📘 Notes
// If you provide values for both category and ids parameters, the category parameter will be prioritized over the ids parameter
// Cache/Update Frequency: Every 45 seconds for all the API plans
func (c *CoinsClient) CoinsMarket(ctx context.Context, currency string, options ...marketsOption) ([]*types.Market, error) {
	params := url.Values{}
	params.Add("vs_currency", currency)

	// Apply all the options
	for _, opt := range options {
		opt.Apply(&params)
	}

	rUrl := fmt.Sprintf("%s/%s?%s", c.coinsUrl(), "markets", params.Encode())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data []*types.Market
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Define option types
type (
	idsOption                   struct{ ids []string }
	categoryOption              struct{ category string }
	orderOption                 struct{ order string }
	perPageOption               struct{ perPage int }
	pageOption                  struct{ page int }
	sparklineOption             struct{ sparkline bool }
	priceChangePercentageOption struct{ intervals []string }
)

// Implement Option interface
func (o idsOption) Apply(v *url.Values)       { v.Set("ids", strings.Join(o.ids, ",")) }
func (o categoryOption) Apply(v *url.Values)  { v.Set("category", o.category) }
func (o orderOption) Apply(v *url.Values)     { v.Set("order", o.order) }
func (o perPageOption) Apply(v *url.Values)   { v.Set("per_page", strconv.Itoa(o.perPage)) }
func (o pageOption) Apply(v *url.Values)      { v.Set("page", strconv.Itoa(o.page)) }
func (o sparklineOption) Apply(v *url.Values) { v.Set("sparkline", strconv.FormatBool(o.sparkline)) }
func (o priceChangePercentageOption) Apply(v *url.Values) {
	v.Set("price_change_percentage", strings.Join(o.intervals, ","))
}

// Implement CoinsMarketOption interface
func (idsOption) isMarketsOption()                   {}
func (categoryOption) isMarketsOption()              {}
func (orderOption) isMarketsOption()                 {}
func (perPageOption) isMarketsOption()               {}
func (pageOption) isMarketsOption()                  {}
func (sparklineOption) isMarketsOption()             {}
func (priceChangePercentageOption) isMarketsOption() {}
