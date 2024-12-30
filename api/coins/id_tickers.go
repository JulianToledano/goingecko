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

// idTickersOption is an interface that extends internal.Option to provide
// specific options for the CoinsIdTickers endpoint. It includes a marker
// method isCoinsIdTickersOption() to ensure type safety for tickers-specific options.
type idTickersOption interface {
	internal.Option
	isCoinsIdTickersOption()
}

// WithExchangeId filters tickers by exchange id.
// Multiple exchange ids can be provided as a comma-separated string.
// Refers to /exchanges/list.
func WithExchangeId(exchangeIds string) idTickersOption { return exchangeIdsOption{exchangeIds} }

// WithIncludeExchangeLogo includes exchange logo URLs in the response if true.
// Default: false
func WithIncludeExchangeLogo(includeLogo bool) idTickersOption {
	return includeExchangeLogoOption{includeLogo}
}

// WithIdTickersPage specifies which page of results to return.
func WithIdTickersPage(page int64) idTickersOption { return pageIdTickersOption{page} }

// WithIdTickersOrder specifies the ordering of results.
// Valid values: "trust_score_desc", "trust_score_asc", "volume_desc", "volume_asc
func WithIdTickersOrder(order string) idTickersOption { return orderIdTickersOption{order} }

// WithDepth includes 2% orderbook depth info if "cost_to_move_up_usd" or "cost_to_move_down_usd".
// Valid values: "", "cost_to_move_up_usd", "cost_to_move_down_usd"
func WithDepth(depth string) idTickersOption { return depthOption{depth} }

// CoinsIdTickers allows you to query the coin tickers on both centralized exchange (cex) and decentralized exchange
// (dex) based on a particular coin id.
//
// 👍 Tips
//
// You may obtain the coin id (api id) via several ways:
// refers to respective coin page and find ‘api id’
// refers to /coins/list endpoint
// refers to google sheets here
// You may specify the exchange_ids if you want to retrieve tickers for specific exchange only
// You may include values such as page to specify which page of responses you would like to show
// You may also flag to include more data such as exchange logo and depth
//
// 📘 Notes
//
// The tickers are paginated to 100 items
// Cache / Update Frequency: every 2 minutes for all the API plans
// When order is sorted by 'volume', converted_volume will be used instead of volume
func (c *CoinsClient) CoinsIdTickers(ctx context.Context, id string, options ...idTickersOption) (*types.Tickers, error) {
	params := url.Values{}

	// Apply all the options
	for _, opt := range options {
		opt.Apply(&params)
	}

	rUrl := fmt.Sprintf("%s/%s/%s?%s", c.coinsUrl(), id, "tickers", params.Encode())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data *types.Tickers
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Define option types
type exchangeIdsOption struct{ exchangeIds string }
type includeExchangeLogoOption struct{ includeLogo bool }
type pageIdTickersOption struct{ page int64 }
type orderIdTickersOption struct{ order string }
type depthOption struct{ depth string }

// Implement Option interface
func (o exchangeIdsOption) Apply(v *url.Values) { v.Set("exchange_ids", o.exchangeIds) }
func (o includeExchangeLogoOption) Apply(v *url.Values) {
	v.Set("include_exchange_logo", strconv.FormatBool(o.includeLogo))
}
func (o pageIdTickersOption) Apply(v *url.Values)  { v.Set("page", strconv.FormatInt(o.page, 10)) }
func (o orderIdTickersOption) Apply(v *url.Values) { v.Set("order", o.order) }
func (o depthOption) Apply(v *url.Values)          { v.Set("depth", o.depth) }

// Implement CoinsIdTickersOption interface
func (exchangeIdsOption) isCoinsIdTickersOption()         {}
func (includeExchangeLogoOption) isCoinsIdTickersOption() {}
func (pageIdTickersOption) isCoinsIdTickersOption()       {}
func (orderIdTickersOption) isCoinsIdTickersOption()      {}
func (depthOption) isCoinsIdTickersOption()               {}
