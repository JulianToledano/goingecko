package search

import (
	"github.com/JulianToledano/goingecko/api/internal"
	geckohttp "github.com/JulianToledano/goingecko/http"
)

type SearchClient struct {
	*internal.Client
}

func NewClient(c *geckohttp.Client, url string) *SearchClient {
	return &SearchClient{
		internal.NewClient(c, url),
	}
}

func (c *SearchClient) searchUrl() string {
	return c.URL + "/search"
}
