package goingecko

import (
	"encoding/json"
	"fmt"
	"github.com/JulianToledano/goingecko/assetPlatforms"
	"net/url"
)

func (c *Client) AssetPlatforms(filter string) (*assetPlatforms.AssetPlatforms, error) {
	params := url.Values{}

	if filter != "" {
		params.Add("filter", filter)
	}

	rUrl := fmt.Sprintf("%s?%s", assetPlatformsURL, params.Encode())
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
