package simple

import (
	"github.com/JulianToledano/goingecko/api/internal"
	geckohttp "github.com/JulianToledano/goingecko/http"
)

type Client struct {
	*internal.Client
}

func NewSimpleClient(c *geckohttp.Client, url string) *Client {
	return &Client{
		internal.NewClient(c, url),
	}
}

func (c *Client) simpleUrl() string {
	return c.URL + "/simple"
}
