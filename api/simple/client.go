package simple

import (
	"github.com/JulianToledano/goingecko/api/internal"
	geckohttp "github.com/JulianToledano/goingecko/http"
)

type SimpleClient struct {
	*internal.Client
}

func NewClient(c *geckohttp.Client, url string) *SimpleClient {
	return &SimpleClient{
		internal.NewClient(c, url),
	}
}

func (c *SimpleClient) simpleUrl() string {
	return c.URL + "/simple"
}