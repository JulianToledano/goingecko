package exchanges

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/JulianToledano/goingecko/v3/api"

	"github.com/JulianToledano/goingecko/v3/api/exchanges/types"
)

// tickersOption is an interface that extends api.Option to provide options specific to the exchange tickers endpoint.
type tickersOption interface {
	api.Option

	isTickersOption()
}

// WithCoinIdsTickers returns a tickersOption that filters tickers by coin IDs.
// The ids parameter specifies which coin IDs to include in the response.
func WithCoinIdsTickers(ids []string) tickersOption {
	return coinIdsTickersOption{ids}
}

// WithIncludeExchangeLogoTickers returns a tickersOption that controls inclusion of exchange logos.
// When include is true, exchange logos will be included in the response.
func WithIncludeExchangeLogoTickers(include bool) tickersOption {
	return includeExchangeLogoTickersOption{include}
}

// WithPageTickers returns a tickersOption that sets which page of results to return.
// The page parameter specifies which page of tickers to return in the response.
func WithPageTickers(page int64) tickersOption {
	return pageTickersOption{page}
}

// WithDepthTickers returns a tickersOption that controls inclusion of order book depth.
// When depth is true, order book depth data will be included in the response.
func WithDepthTickers(depth bool) tickersOption {
	return depthTickersOption{depth}
}

// WithOrderTickers returns a tickersOption that sets the ordering of results.
// The order parameter specifies how to sort the tickers in the response.
func WithOrderTickers(order string) tickersOption {
	return orderTickersOption{order}
}

// ExchangesIdTickers allows you to query exchange's tickers based on exchange’s id.
//
//	📘 Notes
//
//	    Responses are paginated and limited to 100 tickers per page. You may specify the page number using the page params to retrieve the tickers accordingly
//	    Cache / Update Frequency: every 60 seconds for all the API plans
func (c *ExchangesClient) ExchangesIdTickers(ctx context.Context, id string) (*types.Tickers, error) {
	params := url.Values{}

	rUrl := fmt.Sprintf("%s/%s/tickers?%s", c.exchangesUrl(), id, params.Encode())
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

type coinIdsTickersOption struct{ ids []string }
type includeExchangeLogoTickersOption struct{ include bool }
type pageTickersOption struct{ page int64 }
type depthTickersOption struct{ depth bool }
type orderTickersOption struct{ order string }

func (o coinIdsTickersOption) Apply(v *url.Values) {
	v.Set("coin_ids", strings.Join(o.ids, ","))
}
func (o includeExchangeLogoTickersOption) Apply(v *url.Values) {
	v.Set("include_exchange_logo", strconv.FormatBool(o.include))
}
func (o pageTickersOption) Apply(v *url.Values) {
	v.Set("page", strconv.FormatInt(o.page, 10))
}
func (o depthTickersOption) Apply(v *url.Values) {
	v.Set("depth", strconv.FormatBool(o.depth))
}
func (o orderTickersOption) Apply(v *url.Values) {
	v.Set("order", o.order)
}

func (o coinIdsTickersOption) isTickersOption()             {}
func (o includeExchangeLogoTickersOption) isTickersOption() {}
func (o pageTickersOption) isTickersOption()                {}
func (o depthTickersOption) isTickersOption()               {}
func (o orderTickersOption) isTickersOption()               {}
