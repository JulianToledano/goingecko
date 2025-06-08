package search

import (
	"github.com/JulianToledano/goingecko/v3/api/internal"
	geckohttp "github.com/JulianToledano/goingecko/v3/http"
)

type SearchClient struct {
	*internal.Client
}

func NewClient(c geckohttp.HttpClient, url string) *SearchClient {
	return &SearchClient{
		internal.NewClient(c, url),
	}
}

func (c *SearchClient) searchUrl() string {
	return c.URL + "/search"
}
