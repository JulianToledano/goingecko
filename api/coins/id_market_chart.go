package coins

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/JulianToledano/goingecko/api"
	"github.com/JulianToledano/goingecko/types"
)

// idMarketChartOption is an interface that extends api.Option to provide
// specific options for the CoinsIdMarketChart endpoint. It includes a marker
// method isIdMarketChartOptions() to ensure type safety for market chart-specific options.
type idMarketChartOption interface {
	api.Option
	isIdMarketChartOption()
}

// WithIntervalIdMarketChart sets the interval between data points in the response.
// Valid values: 5m, hourly, daily
func WithIntervalIdMarketChart(interval string) idMarketChartOption {
	return intervalIdMarketChartOptions{interval}
}

// WithPrecisionIdMarketChart sets the number of decimal places in the response data.
// Valid values: from 1 to 18
func WithPrecisionIdMarketChart(precision string) idMarketChartOption {
	return precisionIdMarketChartOptions{precision}
}

// CoinsIdMarketChart allows you to get the historical chart data of a coin including time in UNIX, price, market cap
// and 24hrs volume based on particular coin id.
//
// 👍Tips
//
//	You may obtain the coin id (api id) via several ways:
//		refers to respective coin page and find ‘api id’
//	    refers to /coins/list endpoint
//	        refers to google sheets here
//	    You may use tools like epoch converter to convert human readable date to UNIX timestamp
//
//	📘Notes
//	    You may leave the interval params as empty for automatic granularity:
//	        1 day from current time = 5-minutely data
//	        2 - 90 days from current time = hourly data
//	        above 90 days from current time = daily data (00:00 UTC)
//	    For non-Enterprise plan subscribers who would like to get hourly data, please leave the interval params empty for auto granularity
//	    The 5-minutely and hourly interval params are also exclusively available to Enterprise plan subscribers, bypassing auto-granularity:
//	        interval=5m: 5-minutely historical data (responses include information from the past 10 days, up until 2 days ago)
//	        interval=hourly: hourly historical data (responses include information from the past 100 days, up until now)
//	    Cache / Update Frequency:
//	        every 30 seconds for all the API plans (for last data point)
//	        The last completed UTC day (00:00) data is available 10 minutes after midnight on the next UTC day (00:10).
func (c *CoinsClient) CoinsIdMarketChart(ctx context.Context, id, vsCurrency, days string, options ...idMarketChartOption) (*types.MarketChart, error) {
	params := url.Values{}
	params.Add("vs_currency", vsCurrency)
	params.Add("days", days)
	params.Add("interval", "daily")

	for _, opt := range options {
		opt.Apply(&params)
	}

	rUrl := fmt.Sprintf("%s/%s/%s?%s", c.coinsUrl(), id, "market_chart", params.Encode())
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

type intervalIdMarketChartOptions struct{ interval string }
type precisionIdMarketChartOptions struct{ precision string }

func (o intervalIdMarketChartOptions) Apply(v *url.Values)  { v.Set("interval", o.interval) }
func (o precisionIdMarketChartOptions) Apply(v *url.Values) { v.Set("precision", o.precision) }

func (o intervalIdMarketChartOptions) isIdMarketChartOption()  {}
func (o precisionIdMarketChartOptions) isIdMarketChartOption() {}
