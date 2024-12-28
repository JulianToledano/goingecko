package coins

import (
	geckohttp "github.com/JulianToledano/goingecko/http"
)

type Client struct {
	*geckohttp.Client

	url string
}

func NewCoinsClient(c *geckohttp.Client, url string) *Client {
	return &Client{
		c,
		url,
	}
}

func (c *Client) coinsUrl() string {
	return c.url + "/coins"
}
