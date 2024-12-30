package derivatives

import (
	"github.com/JulianToledano/goingecko/api/internal"
	geckohttp "github.com/JulianToledano/goingecko/http"
)

type DerivativesClient struct {
	*internal.Client
}

func NewClient(c *geckohttp.Client, url string) *DerivativesClient {
	return &DerivativesClient{
		internal.NewClient(c, url),
	}
}

func (c *DerivativesClient) derivativesUrl() string {
	return c.URL + "/derivatives"
}
