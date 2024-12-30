package trending

import (
	"context"
	"encoding/json"

	"github.com/JulianToledano/goingecko/api/trending/types"
)

// Trending allows you query trending search coins, nfts and categories on CoinGecko in the last 24 hours.
//
//	📘Notes
//
//	    The endpoint currently supports:
//	        Top 15 trending coins (sorted by the most popular user searches)
//	        Top 7 trending NFTs (sorted by the highest percentage change in floor prices)
//	        Top 6 trending categories (sorted by the most popular user searches)
//	    Cache / Update Frequency: every 10 minutes for all the API plans
func (c *TrendingClient) Trending(ctx context.Context) (*types.Trending, error) {
	resp, err := c.MakeReq(ctx, c.trendingUrl())
	if err != nil {
		return nil, err
	}

	var data *types.Trending
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
