package goingecko

import (
	"context"
	"encoding/json"

	"github.com/JulianToledano/goingecko/trending"
)

func (c *Client) Trending(ctx context.Context) (*trending.Trending, error) {
	resp, err := c.MakeReq(ctx, c.getTrendingURL())
	if err != nil {
		return nil, err
	}

	var data *trending.Trending
	err = json.Unmarshal([]byte(resp), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
