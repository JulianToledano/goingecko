package goingecko

import (
	"encoding/json"
	"fmt"
	"github.com/JulianToledano/goingecko/search"
	"net/url"
)

// Search for coins, categories and markets listed on CoinGecko ordered by largest Market Cap first.
//
// Cache / Update Frequency: every 15 minutes
func (c *Client) Search(query string) (*search.Search, error) {
	params := url.Values{}
	if query != "" {
		params.Add("query", query)
	}

	rUrl := fmt.Sprintf("%s?%s", searchURL, params.Encode())
	resp, err := c.MakeReq(rUrl)
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
