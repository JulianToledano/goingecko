package nfts

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/JulianToledano/goingecko/v3/api/nfts/types"
)

// NftsIdTickers allows you to query floor price and 24h volume for an NFT collection across NFT marketplaces.
//
// This is a Pro API endpoint.
func (c *NftsClient) NftsIdTickers(ctx context.Context, id string) (*types.NftTickers, error) {
	rUrl := fmt.Sprintf("%s/%s/%s", c.nftsUrl(), id, "tickers")
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data *types.NftTickers
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
