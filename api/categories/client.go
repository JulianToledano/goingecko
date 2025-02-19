package categories

import (
	"github.com/JulianToledano/goingecko/v3/api/internal"
	geckohttp "github.com/JulianToledano/goingecko/v3/http"
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
	return c.URL + "/coins/categories"
}
