package coins

import (
	"github.com/JulianToledano/goingecko/api/internal"
	geckohttp "github.com/JulianToledano/goingecko/http"
)

type Client struct {
	*internal.Client
}

func NewCoinsClient(c *geckohttp.Client, url string) *Client {
	return &Client{
		internal.NewClient(c, url),
	}
}

func (c *Client) coinsUrl() string {
	return c.URL + "/coins"
}
