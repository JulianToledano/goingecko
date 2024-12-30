package categories

import (
	"github.com/JulianToledano/goingecko/api/internal"
	geckohttp "github.com/JulianToledano/goingecko/http"
)

type CategoriesClient struct {
	*internal.Client
}

func NewClient(c *geckohttp.Client, url string) *CategoriesClient {
	return &CategoriesClient{
		internal.NewClient(c, url),
	}
}

func (c *CategoriesClient) categoriesUrl() string {
	return c.URL + "/categories"
}
