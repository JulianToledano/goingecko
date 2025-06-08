package global

import (
	"github.com/JulianToledano/goingecko/v3/api/internal"
	geckohttp "github.com/JulianToledano/goingecko/v3/http"
)

type GlobalClient struct {
	*internal.Client
}

func NewClient(c geckohttp.HttpClient, url string) *GlobalClient {
	return &GlobalClient{
		internal.NewClient(c, url),
	}
}

func (c *GlobalClient) globalUrl() string {
	return c.URL + "/global"
}
