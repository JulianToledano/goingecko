package coins

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/JulianToledano/goingecko/v3/api/coins/types"
)

// CoinsIdSupplyBreakdown allows you to query the supply breakdown of a coin based on a provided coin ID.
//
// This is a Pro API endpoint.
func (c *CoinsClient) CoinsIdSupplyBreakdown(ctx context.Context, id string) (*types.SupplyBreakdown, error) {
	rUrl := fmt.Sprintf("%s/%s/%s", c.coinsUrl(), id, "supply_breakdown")
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data *types.SupplyBreakdown
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
