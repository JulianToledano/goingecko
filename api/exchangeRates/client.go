package exchangeRates

import (
	"github.com/JulianToledano/goingecko/v3/api/internal"
	geckohttp "github.com/JulianToledano/goingecko/v3/http"
)

type ExchangeRatesClient struct {
	*internal.Client
}

func NewClient(c *geckohttp.Client, url string) *ExchangeRatesClient {
	return &ExchangeRatesClient{
		internal.NewClient(c, url),
	}
}

func (c *ExchangeRatesClient) exchangeRatesUrl() string {
	return c.URL + "/exchange_rates"
}
