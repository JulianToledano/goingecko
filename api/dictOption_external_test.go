package api_test

import (
	"context"
	"testing"

	"github.com/JulianToledano/goingecko/v3/api"
	"github.com/JulianToledano/goingecko/v3/api/coins"
)

func TestWithURLParamsCompilesAcrossPackages(t *testing.T) {
	t.Skip("compile-only coverage")

	var c *coins.CoinsClient
	_, _ = c.CoinsId(context.Background(), "usd", api.WithURLParams(map[string]any{
		"precision": "full",
	}))
}
