package exchanges

import (
	"github.com/JulianToledano/goingecko/api/internal"
	geckohttp "github.com/JulianToledano/goingecko/http"
)

// TODO: https://docs.coingecko.com/reference/exchanges-id-volume-chart-range

type ExchangesClient struct {
	*internal.Client
}

func NewClient(c *geckohttp.Client, url string) *ExchangesClient {
	return &ExchangesClient{
		internal.NewClient(c, url),
	}
}

func (c *ExchangesClient) exchangesUrl() string {
	return c.URL + "/exchanges"
}
