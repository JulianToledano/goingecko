package nfts

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/JulianToledano/goingecko/api"

	"github.com/JulianToledano/goingecko/api/nfts/types"
)

// listOption is an interface that extends api.Option to provide
// specific options for the NftsList endpoint. It includes a marker
// method isListOption() to ensure type safety for list-specific options.
type listOption interface {
	api.Option
	isListOption()
}

// WithOrder specifies the ordering of results.
// Valid values: h24_volume_native_asc, h24_volume_native_desc, floor_price_native_asc,
// floor_price_native_desc, market_cap_native_asc, market_cap_native_desc,
// market_cap_usd_asc, market_cap_usd_desc, etc.
func WithOrder(order string) listOption {
	return orderListOption{order: order}
}

// WithPerPage specifies the number of results per page.
// Valid values: 1..250
func WithPerPage(perPage int64) listOption {
	return perPageListOption{perPage: perPage}
}

// WithPage specifies which page of results to return.
func WithPage(page int64) listOption {
	return pageListOption{page: page}
}

// NftsList allows you to query all supported NFTs with id, contract address, name, asset platform id and symbol on CoinGecko.
//
//	👍 Tips
//
//	    You may use this endpoint to query the list of nfts for other endpoints that contain params like id (nft collection’s id) as well as asset_platform_id and contract_address
//	    You may include values such as per_page and page to specify how many results you would like to show in the responses per page and which page of responses you would like to show
//
//	📘 Notes
//
//	    The responses are paginated to 100 items
//	    Cache / Update Frequency: every 5 minutes for all the API plans
func (c *NftsClient) NftsList(ctx context.Context, options ...listOption) ([]types.Nft, error) {
	params := url.Values{}

	for _, opt := range options {
		opt.Apply(&params)
	}

	rUrl := fmt.Sprintf("%s/list?%s", c.nftsUrl(), params.Encode())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data []types.Nft
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

type orderListOption struct{ order string }
type perPageListOption struct{ perPage int64 }
type pageListOption struct{ page int64 }

func (o orderListOption) Apply(v *url.Values)   { v.Set("order", o.order) }
func (o perPageListOption) Apply(v *url.Values) { v.Set("per_page", strconv.FormatInt(o.perPage, 10)) }
func (o pageListOption) Apply(v *url.Values)    { v.Set("page", strconv.FormatInt(o.page, 10)) }

func (o orderListOption) isListOption()   {}
func (o perPageListOption) isListOption() {}
func (o pageListOption) isListOption()    {}
