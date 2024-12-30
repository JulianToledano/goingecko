package coins

import (
	"github.com/JulianToledano/goingecko/api/internal"
	geckohttp "github.com/JulianToledano/goingecko/http"
)

type CoinsClient struct {
	*internal.Client
}

func NewClient(c *geckohttp.Client, url string) *CoinsClient {
	return &CoinsClient{
		internal.NewClient(c, url),
	}
}

func (c *CoinsClient) coinsUrl() string {
	return c.URL + "/coins"
}
