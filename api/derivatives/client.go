package derivatives

import (
	"github.com/JulianToledano/goingecko/v3/api/internal"
	geckohttp "github.com/JulianToledano/goingecko/v3/http"
)

type DerivativesClient struct {
	*internal.Client
}

func NewClient(c geckohttp.HttpClient, url string) *DerivativesClient {
	return &DerivativesClient{
		internal.NewClient(c, url),
	}
}

func (c *DerivativesClient) derivativesUrl() string {
	return c.URL + "/derivatives"
}
