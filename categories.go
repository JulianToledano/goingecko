package goingecko

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/JulianToledano/goingecko/v2/categories"
)

// CategoriesList List all categories
// Cache / Update Frequency: every 5 minutes
func (c *Client) CategoriesList(ctx context.Context) (*categories.CategoriesList, error) {
	rUrl := fmt.Sprintf("%s/%s", c.getCategoriesURL(), "list")
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}
	var data *categories.CategoriesList
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Categories List all categories with market data
// Cache / Update Frequency: every 5 minutes
// Parameters:
// order: valid values: market_cap_desc (default), market_cap_asc, name_desc, name_asc, market_cap_change_24h_desc and market_cap_change_24h_asc
func (c *Client) Categories(ctx context.Context, order string) (*categories.CategoriesWithMarketDataList, error) {
	params := url.Values{}

	if order != "" {
		params.Add("order", order)
	}

	rUrl := fmt.Sprintf("%s?%s", c.getCategoriesURL(), params.Encode())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}
	var data *categories.CategoriesWithMarketDataList
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
