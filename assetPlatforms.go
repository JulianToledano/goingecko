package goingecko

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/JulianToledano/goingecko/assetPlatforms"
)

func (c *Client) AssetPlatforms(filter string) (*assetPlatforms.AssetPlatforms, error) {
	params := url.Values{}

	if filter != "" {
		params.Add("filter", filter)
	}

	rUrl := fmt.Sprintf("%s?%s", c.getAssetPlatformsURL(), params.Encode())
	resp, err := c.MakeReq(rUrl)
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
