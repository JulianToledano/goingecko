package nfts

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/JulianToledano/goingecko/v3/api/nfts/types"
)

// NftsContract allows you to query all the NFT data (name, floor price, 24 hr volume....) based on the nft collection contract address and respective asset platform.
//
//	👍 Tips
//
//	    You may also obtain the asset platform id and contract address through /nfts/list endpoint
//
//	📘 Notes
//
//	    Solana NFT & Art Blocks are not supported for this endpoint, please use /nfts/{id} endpoint instead
//	    Cache / Update Frequency: every 60 seconds for all the API plans
func (c *NftsClient) NftsContract(ctx context.Context, assetPlatform, contract string) (*types.NftId, error) {
	rUrl := fmt.Sprintf("%s/%s/contract/%s", c.nftsUrl(), assetPlatform, contract)
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
