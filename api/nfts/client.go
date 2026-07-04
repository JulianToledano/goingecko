package nfts

import (
	"github.com/JulianToledano/goingecko/v3/api/internal"
	geckohttp "github.com/JulianToledano/goingecko/v3/http"
)

type NftsClient struct {
	*internal.Client
}

func NewClient(c geckohttp.HttpClient, url string) *NftsClient {
	return &NftsClient{
		internal.NewClient(c, url),
	}
}

func (c *NftsClient) nftsUrl() string {
	return c.URL + "/nfts"
}
