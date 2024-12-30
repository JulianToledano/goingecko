package nfts

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/JulianToledano/goingecko/api/nfts/types"
)

// NftsId allows you to query all the NFT data (name, floor price, 24 hr volume....) based on the nft collection id.
//
//	📘 Notes
//
//	    Cache / Update Frequency: every 60 seconds for all the API plans
func (c *NftsClient) NftsId(ctx context.Context, id string) (*types.NftId, error) {
	rUrl := fmt.Sprintf("%s/%s", c.nftsUrl(), id)
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data *types.NftId
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
