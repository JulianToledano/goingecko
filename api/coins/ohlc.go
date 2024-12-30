package coins

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/JulianToledano/goingecko/v3/api"
	"github.com/JulianToledano/goingecko/v3/api/coins/types"
)

// ohlcOption is an interface that extends api.Option to provide
// specific options for the CoinsOhlc endpoint. It includes a marker
// method isOhlcOption() to ensure type safety for OHLC-specific options.
type ohlcOption interface {
	api.Option
	isOhlcOption()
}

// WithOhlcIntervalOption sets the interval between data points in the response.
// Valid values: hourly, daily
func WithOhlcIntervalOption(interval string) ohlcOption {
	return intervalOhlcOption{interval: interval}
}

// WithOhlcPrecisionOption sets the number of decimal places in the response data.
// Valid values: from 1 to 18
func WithOhlcPrecisionOption(precision string) ohlcOption {
	return precisionOhlcOption{precision: precision}
}

// CoinsOhlc allows you to get the OHLC chart (Open, High, Low, Close) of a coin based on particular coin id.
//
//	 👍Tips
//
//	    You may obtain the coin id (api id) via several ways:
//	        refers to respective coin page and find ‘api id’
//	        refers to /coins/list endpoint
//	        refers to google sheets here
//	    For historical chart data with better granularity, you may consider using /coins/{id}/market_chart endpoint
//
//	📘 Notes
//
//	    The timestamp displayed in the payload (response) indicates the end (or close) time of the OHLC data
//	    Data granularity (candle's body) is automatic:
//	        1 - 2 days: 30 minutes
//	        3 - 30 days: 4 hours
//	        31 days and beyond: 4 days
//	    Cache / Update Frequency:
//	        every 15 minutes for all the API plans
//	        The last completed UTC day (00:00) is available 35 minutes after midnight on the next UTC day (00:35)
//	    Exclusive daily and hourly candle interval parameter for all paid plan subscribers (interval = daily, interval=hourly)
//	        'daily' interval is available for 1 / 7 / 14 / 30 / 90 / 180 days only.
//	        'hourly' interval is available for 1 /7 / 14 / 30 / 90 days only.
func (c *CoinsClient) CoinsOhlc(ctx context.Context, id, vsCurrency, days string, options ...ohlcOption) (*types.Ohlc, error) {
	params := url.Values{}
	params.Add("vs_currency", vsCurrency)
	params.Add("days", days)

	for _, opt := range options {
		opt.Apply(&params)
	}

	rUrl := fmt.Sprintf("%s/%s/%s?%s", c.coinsUrl(), id, "ohlc", params.Encode())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data *types.Ohlc
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

type intervalOhlcOption struct{ interval string }
type precisionOhlcOption struct{ precision string }

func (o intervalOhlcOption) Apply(v *url.Values)  { v.Set("interval", o.interval) }
func (o precisionOhlcOption) Apply(v *url.Values) { v.Set("precision", o.precision) }

func (o intervalOhlcOption) isOhlcOption()  {}
func (o precisionOhlcOption) isOhlcOption() {}
