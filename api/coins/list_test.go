package coins

import (
	"context"
	"net/http"
	"testing"

	"github.com/JulianToledano/goingecko/v3/api/internal"
	geckohttp "github.com/JulianToledano/goingecko/v3/http"
)

func TestCoinsClient_CoinsList(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CoinsClient{
				internal.NewClient(
					geckohttp.NewClient(geckohttp.WithHttpClient(http.DefaultClient)),
					internal.BaseURL,
				),
			}
			got, err := c.CoinsList(context.Background(), WithIncludePlatform(true))
			if err != nil {
				t.Errorf("CoinsList() error = %v", err)
			}
			if got == nil {
				t.Errorf("CoinsList() got = nil, want not nil")
			}
		})
	}
}
