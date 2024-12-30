package assetPlatforms

import (
	"github.com/JulianToledano/goingecko/api/internal"
	geckohttp "github.com/JulianToledano/goingecko/http"
)

type AssetPlatformsClient struct {
	*internal.Client
}

func NewClient(c *geckohttp.Client, url string) *AssetPlatformsClient {
	return &AssetPlatformsClient{
		internal.NewClient(c, url),
	}
}

func (c *AssetPlatformsClient) assetPlatformsUrl() string {
	return c.URL + "/asset_platforms"
}

// TODO: https://docs.coingecko.com/reference/token-lists