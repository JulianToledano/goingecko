package exchanges

import (
	"github.com/JulianToledano/goingecko/v3/api/internal"
	geckohttp "github.com/JulianToledano/goingecko/v3/http"
)

// TODO: https://docs.coingecko.com/reference/exchanges-id-volume-chart-range

type ExchangesClient struct {
	*internal.Client
}

func NewClient(c geckohttp.HttpClient, url string) *ExchangesClient {
	return &ExchangesClient{
		internal.NewClient(c, url),
	}
}

func (c *ExchangesClient) exchangesUrl() string {
	return c.URL + "/exchanges"
}
