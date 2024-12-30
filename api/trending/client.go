package trending

import (
	"github.com/JulianToledano/goingecko/v3/api/internal"
	geckohttp "github.com/JulianToledano/goingecko/v3/http"
)

type TrendingClient struct {
	*internal.Client
}

func NewClient(c *geckohttp.Client, url string) *TrendingClient {
	return &TrendingClient{
		internal.NewClient(c, url),
	}
}

func (c *TrendingClient) trendingUrl() string {
	return c.URL + "/search/trending"
}
