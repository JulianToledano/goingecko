package goingecko

import (
	"encoding/json"
	"fmt"
	"github.com/JulianToledano/goingecko/categories"
	"net/url"
)

// CategoriesList List all categories
// Cache / Update Frequency: every 5 minutes
func (c *Client) CategoriesList() (*categories.CategoriesList, error) {
	rUrl := fmt.Sprintf("%s/%s", categoriesURL, "list")
	resp, err := c.MakeReq(rUrl)
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
func (c *Client) Categories(order string) (*categories.CategoriesWithMarketDataList, error) {
	params := url.Values{}

	if order != "" {
		params.Add("order", order)
	}

	rUrl := fmt.Sprintf("%s?%s", categoriesURL, params.Encode())
	resp, err := c.MakeReq(rUrl)
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
