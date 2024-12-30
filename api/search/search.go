package search

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/JulianToledano/goingecko/v3/api/search/types"
)

// Search allows you to query BTC exchange rates with other currencies.
//
//	👍 Tips
//
//	    You may use this endpoint to convert the response data, which is originally in BTC, to other currencies
//
//	📘Notes
//
//	    Cache / Update Frequency: every 5 minutes for all the API plans
func (c *SearchClient) Search(ctx context.Context, query string) (*types.Search, error) {
	params := url.Values{}
	if query != "" {
		params.Add("query", query)
	}

	rUrl := fmt.Sprintf("%s?%s", c.searchUrl(), params.Encode())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data *types.Search
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
