package goingecko

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/JulianToledano/goingecko/v2/nfts"
)

// NftsList Use this to obtain all the NFT ids in order to make API calls, paginated to 100 items.
//
// Cache / Update Frequency: every 5 minutes
// Parameters:
// order(string) - valid values: h24_volume_native_asc, h24_volume_native_desc, floor_price_native_asc, floor_price_native_desc, market_cap_native_asc, market_cap_native_desc, market_cap_usd_asc, market_cap_usd_desc
// assetPlatformId(string) - The id of the platform issuing tokens (See asset_platforms endpoint for list of options)
// per_page(integer) - Valid values: 1..250. Total results per page
// page(integer) - Page through results
func (c *Client) NftsList(ctx context.Context, order, assetPlatformId string, perPage, page int32) ([]nfts.Nft, error) {
	params := url.Values{}
	if order != "" {
		params.Add("order", order)
	}
	if assetPlatformId != "" {
		params.Add("asset_platform_id", assetPlatformId)
	}
	if perPage > 0 {
		params.Add("per_page", string(perPage))
	}
	if page > 0 {
		params.Add("page", string(page))
	}

	rUrl := fmt.Sprintf("%s/list?%s", c.getNftsURL(), params.Encode())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data []nfts.Nft
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NftsId Get current data (name, price_floor, volume_24h ...) for an NFT collection. native_currency (string) is only a representative of the currency.
//
// Cache / Update Frequency: every 60 seconds
// Parameters:
// id*(string) - id of nft collection (can be obtained from /nfts/list)
func (c *Client) NftsId(ctx context.Context, id string) (*nfts.NftId, error) {
	rUrl := fmt.Sprintf("%s/%s", c.getNftsURL(), id)
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data *nfts.NftId
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NftsContract
func (c *Client) NftsContract(ctx context.Context, assetPlatform, contract string) (*nfts.NftId, error) {
	rUrl := fmt.Sprintf("%s/%s/contract/%s", c.getNftsURL(), assetPlatform, contract)
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data *nfts.NftId
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
