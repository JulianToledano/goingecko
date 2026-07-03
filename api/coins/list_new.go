package coins

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/JulianToledano/goingecko/v3/api/coins/types"
)

// CoinsListNew allows you to query the latest 200 coins recently listed on CoinGecko.
//
// This is a Pro API endpoint.
func (c *CoinsClient) CoinsListNew(ctx context.Context) ([]*types.RecentlyAddedCoin, error) {
	rUrl := fmt.Sprintf("%s/%s", c.coinsUrl(), "list/new")
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data []*types.RecentlyAddedCoin
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
