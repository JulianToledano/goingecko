package assetPlatforms

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/JulianToledano/goingecko/v3/api/assetPlatforms/types"
	"github.com/JulianToledano/goingecko/v3/api/internal"
)

// assetPlatformsOption is an interface that extends internal.Option to provide
// asset platform-specific options
type assetPlatformsOption interface {
	internal.Option

	isAssetPlatformsOption()
}

// WithFilter returns an assetPlatformsOption that sets the filter parameter
// for asset platform requests. The filter parameter can be used to filter
// asset platforms by name or ID.
// Current supported filters: [nft]
func WithFilter(filter string) assetPlatformsOption {
	return &filterOption{filter}
}

// AssetPlatforms allows you to query all the asset platforms on CoinGecko.
//
//	👍 Tips
//
//	    You may use this endpoint to query the list of asset platforms for other endpoints that contain params like id orids(asset platforms)
//	    You may include NFT at the filter params to get the list of NFT-support asset platforms on CoinGecko
func (c *AssetPlatformsClient) AssetPlatforms(ctx context.Context, opts ...assetPlatformsOption) (*types.AssetPlatforms, error) {
	params := url.Values{}

	for _, opt := range opts {
		opt.Apply(&params)
	}

	rUrl := fmt.Sprintf("%s?%s", c.assetPlatformsUrl(), params.Encode())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}
	var data *types.AssetPlatforms
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type filterOption struct{ filter string }

func (o filterOption) Apply(v *url.Values) {
	v.Set("filter", o.filter)
}

func (o filterOption) isAssetPlatformsOption() {}
