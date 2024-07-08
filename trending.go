package goingecko

import (
	"encoding/json"

	"github.com/JulianToledano/goingecko/trending"
)

func (c *Client) Trending() (*trending.Trending, error) {
	resp, err := c.MakeReq(c.getTrendingURL())
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
