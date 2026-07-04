package global

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/JulianToledano/goingecko/v3/api/global/types"
	"github.com/JulianToledano/goingecko/v3/api/internal"
)

// marketCapChartOption is an interface that extends internal.Option to provide
// specific options for the GlobalMarketCapChart endpoint.
type marketCapChartOption interface {
	internal.Option
	IsMarketCapChartOption()
}

// WithGlobalMarketCapChartVsCurrency sets the target currency for market cap and volume values.
func WithGlobalMarketCapChartVsCurrency(vsCurrency string) marketCapChartOption {
	return vsCurrencyMarketCapChartOption{vsCurrency: vsCurrency}
}

// GlobalMarketCapChart allows you to query historical global market cap and volume data by number of days from now.
//
// This is a Pro API endpoint.
func (c *GlobalClient) GlobalMarketCapChart(ctx context.Context, days string, options ...marketCapChartOption) (*types.MarketCapChart, error) {
	params := url.Values{}
	params.Add("days", days)

	for _, opt := range options {
		opt.Apply(&params)
	}

	rUrl := fmt.Sprintf("%s/%s?%s", c.globalUrl(), "market_cap_chart", params.Encode())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data *types.MarketCapChart
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

type vsCurrencyMarketCapChartOption struct {
	vsCurrency string
}

func (o vsCurrencyMarketCapChartOption) Apply(v *url.Values) {
	v.Set("vs_currency", o.vsCurrency)
}

func (vsCurrencyMarketCapChartOption) IsMarketCapChartOption() {}
