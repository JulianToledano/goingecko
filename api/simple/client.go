package simple

import (
	"github.com/JulianToledano/goingecko/v3/api/internal"
	geckohttp "github.com/JulianToledano/goingecko/v3/http"
)

type SimpleClient struct {
	*internal.Client
}

func NewClient(c geckohttp.HttpClient, url string) *SimpleClient {
	return &SimpleClient{
		internal.NewClient(c, url),
	}
}

func (c *SimpleClient) simpleUrl() string {
	return c.URL + "/simple"
}
