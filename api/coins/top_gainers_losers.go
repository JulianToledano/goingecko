package coins

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/JulianToledano/goingecko/v3/api/coins/types"
	"github.com/JulianToledano/goingecko/v3/api/internal"
)

// topGainersLosersOption is an interface that extends internal.Option to provide
// specific options for the CoinsTopGainersLosers endpoint.
type topGainersLosersOption interface {
	internal.Option
	isTopGainersLosersOption()
}

// WithTopGainersLosersDuration filters the result by time range.
// Valid values: "1h", "24h", "7d", "14d", "30d", "60d", "1y".
func WithTopGainersLosersDuration(duration string) topGainersLosersOption {
	return topGainersLosersDurationOption{duration: duration}
}

// WithTopGainersLosersPriceChangePercentage includes price change percentage for specified intervals.
// Valid intervals: "1h", "24h", "7d", "14d", "30d", "60d", "200d", "1y".
func WithTopGainersLosersPriceChangePercentage(intervals []string) topGainersLosersOption {
	return topGainersLosersPriceChangePercentageOption{intervals: intervals}
}

// WithTopCoins filters the result by market cap ranking or all coins.
// Valid values: "300", "500", "1000", "all".
func WithTopCoins(topCoins string) topGainersLosersOption {
	return topCoinsOption{topCoins: topCoins}
}

// CoinsTopGainersLosers allows you to query the top 30 coins with largest price gain and loss by a specific time duration.
//
// This is a Pro API endpoint.
func (c *CoinsClient) CoinsTopGainersLosers(ctx context.Context, currency string, options ...topGainersLosersOption) (*types.TopGainersLosers, error) {
	params := url.Values{}
	params.Add("vs_currency", currency)

	for _, opt := range options {
		opt.Apply(&params)
	}

	rUrl := fmt.Sprintf("%s/%s?%s", c.coinsUrl(), "top_gainers_losers", params.Encode())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data *types.TopGainersLosers
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

type (
	topGainersLosersDurationOption              struct{ duration string }
	topGainersLosersPriceChangePercentageOption struct{ intervals []string }
	topCoinsOption                              struct{ topCoins string }
)

func (o topGainersLosersDurationOption) Apply(v *url.Values) {
	v.Set("duration", o.duration)
}

func (o topGainersLosersPriceChangePercentageOption) Apply(v *url.Values) {
	v.Set("price_change_percentage", strings.Join(o.intervals, ","))
}

func (o topCoinsOption) Apply(v *url.Values) {
	v.Set("top_coins", o.topCoins)
}

func (topGainersLosersDurationOption) isTopGainersLosersOption()              {}
func (topGainersLosersPriceChangePercentageOption) isTopGainersLosersOption() {}
func (topCoinsOption) isTopGainersLosersOption()                              {}
