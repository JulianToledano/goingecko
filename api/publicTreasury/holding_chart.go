package publicTreasury

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/JulianToledano/goingecko/v3/api/internal"
	"github.com/JulianToledano/goingecko/v3/api/publicTreasury/types"
)

// holdingChartOption is an interface that extends internal.Option to provide options specific to the public treasury holding chart endpoint.
type holdingChartOption interface {
	internal.Option

	IsHoldingChartOption()
}

// WithHoldingChartDays returns a holdingChartOption that sets the number of days of historical data to include.
func WithHoldingChartDays(days string) holdingChartOption {
	return holdingChartDaysOption{days}
}

// WithHoldingChartIncludeEmptyIntervals returns a holdingChartOption that sets whether empty intervals with no transaction data are included.
func WithHoldingChartIncludeEmptyIntervals(includeEmptyIntervals bool) holdingChartOption {
	return holdingChartIncludeEmptyIntervalsOption{includeEmptyIntervals}
}

// PublicTreasuryHoldingChart allows you to query historical cryptocurrency holdings chart data by entity ID and coin ID.
func (c *PublicTreasuryClient) PublicTreasuryHoldingChart(ctx context.Context, entityID, coinID string, options ...holdingChartOption) (*types.HoldingChart, error) {
	params := url.Values{}
	for _, opt := range options {
		opt.Apply(&params)
	}

	rUrl := fmt.Sprintf("%s/public_treasury/%s/%s/holding_chart?%s", c.URL, entityID, coinID, params.Encode())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data *types.HoldingChart
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type holdingChartDaysOption struct{ days string }

func (o holdingChartDaysOption) Apply(v *url.Values) {
	v.Set("days", o.days)
}

func (o holdingChartDaysOption) IsHoldingChartOption() {}

type holdingChartIncludeEmptyIntervalsOption struct{ includeEmptyIntervals bool }

func (o holdingChartIncludeEmptyIntervalsOption) Apply(v *url.Values) {
	v.Set("include_empty_intervals", strconv.FormatBool(o.includeEmptyIntervals))
}

func (o holdingChartIncludeEmptyIntervalsOption) IsHoldingChartOption() {}
