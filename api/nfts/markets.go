package nfts

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/JulianToledano/goingecko/v3/api/internal"
	"github.com/JulianToledano/goingecko/v3/api/nfts/types"
)

// marketsOption is an interface that extends internal.Option to provide
// specific options for the NftsMarkets endpoint.
type marketsOption interface {
	internal.Option
	IsMarketsOption()
}

// WithNftsMarketsAssetPlatformID filters results by asset platform.
func WithNftsMarketsAssetPlatformID(assetPlatformID string) marketsOption {
	return assetPlatformIDMarketsOption{assetPlatformID: assetPlatformID}
}

// WithNftsMarketsOrder specifies the ordering of results.
// Valid values: h24_volume_native_asc, h24_volume_native_desc, h24_volume_usd_asc,
// h24_volume_usd_desc, market_cap_usd_asc, market_cap_usd_desc.
func WithNftsMarketsOrder(order string) marketsOption {
	return orderMarketsOption{order: order}
}

// WithNftsMarketsPerPage specifies the number of results per page.
// Valid values: 1..250.
func WithNftsMarketsPerPage(perPage int64) marketsOption {
	return perPageMarketsOption{perPage: perPage}
}

// WithNftsMarketsPage specifies which page of results to return.
func WithNftsMarketsPage(page int64) marketsOption {
	return pageMarketsOption{page: page}
}

// NftsMarkets allows you to query supported NFT collections with floor price, market cap, volume and market data.
//
// This is a Pro API endpoint.
func (c *NftsClient) NftsMarkets(ctx context.Context, options ...marketsOption) ([]types.NftMarket, error) {
	params := url.Values{}

	for _, opt := range options {
		opt.Apply(&params)
	}

	rUrl := fmt.Sprintf("%s/markets?%s", c.nftsUrl(), params.Encode())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data []types.NftMarket
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

type (
	assetPlatformIDMarketsOption struct{ assetPlatformID string }
	orderMarketsOption           struct{ order string }
	perPageMarketsOption         struct{ perPage int64 }
	pageMarketsOption            struct{ page int64 }
)

func (o assetPlatformIDMarketsOption) Apply(v *url.Values) {
	v.Set("asset_platform_id", o.assetPlatformID)
}
func (o orderMarketsOption) Apply(v *url.Values) { v.Set("order", o.order) }
func (o perPageMarketsOption) Apply(v *url.Values) {
	v.Set("per_page", strconv.FormatInt(o.perPage, 10))
}
func (o pageMarketsOption) Apply(v *url.Values) { v.Set("page", strconv.FormatInt(o.page, 10)) }

func (assetPlatformIDMarketsOption) IsMarketsOption() {}
func (orderMarketsOption) IsMarketsOption()           {}
func (perPageMarketsOption) IsMarketsOption()         {}
func (pageMarketsOption) IsMarketsOption()            {}
