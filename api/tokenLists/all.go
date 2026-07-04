package tokenLists

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/JulianToledano/goingecko/v3/api/tokenLists/types"
)

// TokenListsAll allows you to query all tokens for an asset platform supported by the Ethereum token list standard.
func (c *TokenListsClient) TokenListsAll(ctx context.Context, assetPlatformID string) (*types.TokenList, error) {
	rUrl := fmt.Sprintf("%s/%s/all.json", c.tokenListsUrl(), assetPlatformID)
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data *types.TokenList
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
