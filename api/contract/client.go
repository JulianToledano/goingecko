package contract

import (
	"github.com/JulianToledano/goingecko/api/internal"
	geckohttp "github.com/JulianToledano/goingecko/http"
)

type ContractClient struct {
	*internal.Client
}

func NewClient(c *geckohttp.Client, url string) *ContractClient {
	return &ContractClient{
		internal.NewClient(c, url),
	}
}

func (c *ContractClient) contractUrl() string {
	return c.URL + "/coins"
}
