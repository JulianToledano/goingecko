package categories

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/JulianToledano/goingecko/v3/api/categories/types"
)

// CategoriesList allows you to query all the coins categories on CoinGecko.
//
//	👍 Tips
//
//	    You may use this endpoint to query the list of categories for other endpoints that contain params like category
//
//	📘 Notes
//
//	    CoinGecko Equivalent Page: https://www.coingecko.com/en/categories
//	    Cache / Update Frequency: every 5 minutes for all the API plans
//	    CoinGecko categories are different from GeckoTerminal categories
func (c *CategoriesClient) CategoriesList(ctx context.Context) (*types.CategoriesList, error) {
	rUrl := fmt.Sprintf("%s/%s", c.categoriesUrl(), "list")
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}
	var data *types.CategoriesList
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
