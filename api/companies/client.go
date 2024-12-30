package companies

import (
	"github.com/JulianToledano/goingecko/api/internal"
	geckohttp "github.com/JulianToledano/goingecko/http"
)

type CompaniesClient struct {
	*internal.Client
}

func NewClient(c *geckohttp.Client, url string) *CompaniesClient {
	return &CompaniesClient{
		internal.NewClient(c, url),
	}
}

func (c *CompaniesClient) companiesUrl() string {
	return c.URL + "/companies"
}
