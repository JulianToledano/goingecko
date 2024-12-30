package global

import (
	"github.com/JulianToledano/goingecko/api/internal"
	geckohttp "github.com/JulianToledano/goingecko/http"
)

type GlobalClient struct {
	*internal.Client
}

func NewClient(c *geckohttp.Client, url string) *GlobalClient {
	return &GlobalClient{
		internal.NewClient(c, url),
	}
}

func (c *GlobalClient) globalUrl() string {
	return c.URL + "/global"
}
