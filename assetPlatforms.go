package goingecko

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/JulianToledano/goingecko/v2/assetPlatforms"
)

func (c *Client) AssetPlatforms(ctx context.Context, filter string) (*assetPlatforms.AssetPlatforms, error) {
	params := url.Values{}

	if filter != "" {
		params.Add("filter", filter)
	}

	rUrl := fmt.Sprintf("%s?%s", c.getAssetPlatformsURL(), params.Encode())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}
	var data *assetPlatforms.AssetPlatforms
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
