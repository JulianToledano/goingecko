package coins

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/JulianToledano/goingecko/v3/api/coins/types"
	"github.com/JulianToledano/goingecko/v3/api/internal"
)

// idTotalSupplyChartOption is an interface that extends internal.Option to provide
// specific options for the CoinsIdTotalSupplyChart endpoint.
type idTotalSupplyChartOption interface {
	internal.Option
	IsIdTotalSupplyChartOption()
}

// WithIntervalIdTotalSupplyChart sets the interval between data points in the response.
// Valid value: daily.
func WithIntervalIdTotalSupplyChart(interval string) idTotalSupplyChartOption {
	return intervalIdTotalSupplyChartOption{interval: interval}
}

// CoinsIdTotalSupplyChart allows you to query historical total supply of a coin by number of days from now.
//
// This is an Enterprise API endpoint.
func (c *CoinsClient) CoinsIdTotalSupplyChart(ctx context.Context, id, days string, options ...idTotalSupplyChartOption) (*types.SupplyChart, error) {
	params := url.Values{}
	params.Add("days", days)

	for _, opt := range options {
		opt.Apply(&params)
	}

	rUrl := fmt.Sprintf("%s/%s/%s?%s", c.coinsUrl(), id, "total_supply_chart", params.Encode())
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

type intervalIdTotalSupplyChartOption struct {
	interval string
}

func (o intervalIdTotalSupplyChartOption) Apply(v *url.Values) {
	v.Set("interval", o.interval)
}

func (intervalIdTotalSupplyChartOption) IsIdTotalSupplyChartOption() {}
