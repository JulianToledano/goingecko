package categories

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/JulianToledano/goingecko/api"
	"github.com/JulianToledano/goingecko/api/categories/types"
)

// categoriesOption is an interface that extends api.Option to provide
// category-specific options
type categoriesOption interface {
	api.Option

	isCategoryOptions()
}

// WithOrderOption returns a categoriesOption that sets the order parameter
// for category requests. The order parameter can be used to sort categories
// by market cap, volume, etc.
func WithOrderOption(order string) categoriesOption {
	return &orderOption{order}
}

// Categories allows you to query all the coins categories with market data (market cap, volume, etc.) on CoinGecko.
//
//	📘Notes
//
//	    CoinGecko Equivalent Page: https://www.coingecko.com/en/categories
//	    Cache / Update Frequency: every 5 minutes for all the API plans
//	    CoinGecko categories are different from GeckoTerminal categories
func (c *CategoriesClient) Categories(ctx context.Context, options ...categoriesOption) (*types.CategoriesWithMarketDataList, error) {
	params := url.Values{}

	for _, opt := range options {
		opt.Apply(&params)
	}

	rUrl := fmt.Sprintf("%s?%s", c.categoriesUrl(), params.Encode())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}
	var data *types.CategoriesWithMarketDataList
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type orderOption struct{ order string }

func (o orderOption) Apply(v *url.Values) {
	v.Set("order", o.order)
}
func (o orderOption) isCategoryOptions() {}
