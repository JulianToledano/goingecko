package goingecko

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/JulianToledano/goingecko/search"
)

// Search for coins, categories and markets listed on CoinGecko ordered by largest Market Cap first.
//
// Cache / Update Frequency: every 15 minutes
func (c *Client) Search(ctx context.Context, query string) (*search.Search, error) {
	params := url.Values{}
	if query != "" {
		params.Add("query", query)
	}

	rUrl := fmt.Sprintf("%s?%s", c.getSearchURL(), params.Encode())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data *search.Search
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
