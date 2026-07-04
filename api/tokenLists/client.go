package tokenLists

import (
	"github.com/JulianToledano/goingecko/v3/api/internal"
	geckohttp "github.com/JulianToledano/goingecko/v3/http"
)

type TokenListsClient struct {
	*internal.Client
}

func NewClient(c geckohttp.HttpClient, url string) *TokenListsClient {
	return &TokenListsClient{
		internal.NewClient(c, url),
	}
}

func (c *TokenListsClient) tokenListsUrl() string {
	return c.URL + "/token_lists"
}
