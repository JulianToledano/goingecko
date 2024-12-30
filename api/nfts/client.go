package nfts

import (
	"github.com/JulianToledano/goingecko/api/internal"
	geckohttp "github.com/JulianToledano/goingecko/http"
)

// TODO: to implement
//   - https://docs.coingecko.com/reference/nfts-markets
//   - https://docs.coingecko.com/reference/nfts-id-market-chart
//   - https://docs.coingecko.com/reference/nfts-contract-address-market-chart
//   - https://docs.coingecko.com/reference/nfts-id-tickers

type NftsClient struct {
	*internal.Client
}

func NewClient(c *geckohttp.Client, url string) *NftsClient {
	return &NftsClient{
		internal.NewClient(c, url),
	}
}

func (c *NftsClient) nftsUrl() string {
	return c.URL + "/nfts"
}
