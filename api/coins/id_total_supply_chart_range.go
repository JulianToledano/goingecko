package coins

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/JulianToledano/goingecko/v3/api/coins/types"
)

// CoinsIdTotalSupplyChartRange allows you to query historical total supply of a coin within a time range.
//
// This is an Enterprise API endpoint.
func (c *CoinsClient) CoinsIdTotalSupplyChartRange(ctx context.Context, id, from, to string) (*types.SupplyChart, error) {
	params := url.Values{}
	params.Add("from", from)
	params.Add("to", to)

	rUrl := fmt.Sprintf("%s/%s/%s?%s", c.coinsUrl(), id, "total_supply_chart/range", params.Encode())
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
