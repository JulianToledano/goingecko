package nfts

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/JulianToledano/goingecko/v3/api/nfts/types"
)

// NftsIdMarketChart allows you to query historical market data of an NFT collection by number of days from now.
//
// This is a Pro API endpoint.
func (c *NftsClient) NftsIdMarketChart(ctx context.Context, id, days string) (*types.NftMarketChart, error) {
	params := url.Values{}
	params.Add("days", days)

	rUrl := fmt.Sprintf("%s/%s/%s?%s", c.nftsUrl(), id, "market_chart", params.Encode())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data *types.NftMarketChart
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
