package coins

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/JulianToledano/goingecko/v3/api/coins/types"
	"github.com/JulianToledano/goingecko/v3/api/internal"
)

// idCirculatingSupplyChartOption is an interface that extends internal.Option to provide
// specific options for the CoinsIdCirculatingSupplyChart endpoint.
type idCirculatingSupplyChartOption interface {
	internal.Option
	isIdCirculatingSupplyChartOption()
}

// WithIntervalIdCirculatingSupplyChart sets the interval between data points in the response.
// Valid values: 5m, hourly, daily.
func WithIntervalIdCirculatingSupplyChart(interval string) idCirculatingSupplyChartOption {
	return intervalIdCirculatingSupplyChartOption{interval: interval}
}

// CoinsIdCirculatingSupplyChart allows you to query historical circulating supply of a coin by number of days from now.
//
// This is an Enterprise API endpoint.
func (c *CoinsClient) CoinsIdCirculatingSupplyChart(ctx context.Context, id, days string, options ...idCirculatingSupplyChartOption) (*types.SupplyChart, error) {
	params := url.Values{}
	params.Add("days", days)

	for _, opt := range options {
		opt.Apply(&params)
	}

	rUrl := fmt.Sprintf("%s/%s/%s?%s", c.coinsUrl(), id, "circulating_supply_chart", params.Encode())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data *types.SupplyChart
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

type intervalIdCirculatingSupplyChartOption struct {
	interval string
}

func (o intervalIdCirculatingSupplyChartOption) Apply(v *url.Values) {
	v.Set("interval", o.interval)
}

func (intervalIdCirculatingSupplyChartOption) isIdCirculatingSupplyChartOption() {}
