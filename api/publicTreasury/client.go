package publicTreasury

import (
	"github.com/JulianToledano/goingecko/v3/api/internal"
	geckohttp "github.com/JulianToledano/goingecko/v3/http"
)

type PublicTreasuryClient struct {
	*internal.Client
}

func NewClient(c geckohttp.HttpClient, url string) *PublicTreasuryClient {
	return &PublicTreasuryClient{
		internal.NewClient(c, url),
	}
}
