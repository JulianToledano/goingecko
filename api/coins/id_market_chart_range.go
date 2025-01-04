package coins

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/JulianToledano/goingecko/v3/api/internal"
	"github.com/JulianToledano/goingecko/v3/api/types"
)

// idMarketChartRangeOption is an interface that extends internal.Option to provide
// specific options for the CoinsIdMarketChartRange endpoint. It includes a marker
// method isIdMarketChartRangeOption() to ensure type safety for market chart range-specific options.
type idMarketChartRangeOption interface {
	internal.Option
	isIdMarketChartRangeOption()
}

// WithIntervalIdMarketChartRange sets the interval between data points in the response.
// Valid values: 5m, hourly, daily
func WithIntervalIdMarketChartRange(interval string) idMarketChartRangeOption {
	return intervalIdMarketChartRangeOption{interval}
}

// WithPrecisionIdMarketChartRange sets the number of decimal places in the response data.
// Valid values: from 1 to 18
func WithPrecisionIdMarketChartRange(precision string) idMarketChartRangeOption {
	return precisionIdMarketChartRangeOption{precision}
}

func (c *CoinsClient) CoinsIdMarketChartRange(ctx context.Context, id, currency, from, to string, options ...idMarketChartRangeOption) (*types.MarketChart, error) {
	params := url.Values{}
	params.Add("vs_currency", currency)
	params.Add("from", from)
	params.Add("to", to)

	for _, opt := range options {
		opt.Apply(&params)
	}

	rUrl := fmt.Sprintf("%s/%s/%s?%s", c.coinsUrl(), id, "market_chart/range", params.Encode())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data *types.MarketChart
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

type (
	intervalIdMarketChartRangeOption  struct{ interval string }
	precisionIdMarketChartRangeOption struct{ precision string }
)

func (o intervalIdMarketChartRangeOption) Apply(v *url.Values)  { v.Set("interval", o.interval) }
func (o precisionIdMarketChartRangeOption) Apply(v *url.Values) { v.Set("precision", o.precision) }

func (o intervalIdMarketChartRangeOption) isIdMarketChartRangeOption()  {}
func (o precisionIdMarketChartRangeOption) isIdMarketChartRangeOption() {}
