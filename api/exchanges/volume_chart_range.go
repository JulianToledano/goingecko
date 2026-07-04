package exchanges

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/JulianToledano/goingecko/v3/api/exchanges/types"
)

// ExchangesIdVolumeChartRange allows you to query historical volume chart data in BTC by date range based on an exchange ID.
//
// This is a Pro API endpoint.
func (c *ExchangesClient) ExchangesIdVolumeChartRange(ctx context.Context, id, from, to string) ([]types.Volume, error) {
	params := url.Values{}
	params.Add("from", from)
	params.Add("to", to)

	rUrl := fmt.Sprintf("%s/%s/%s?%s", c.exchangesUrl(), id, "volume_chart/range", params.Encode())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data []types.Volume
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
