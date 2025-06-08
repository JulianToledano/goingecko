package main

import (
	"context"
	"fmt"

	"github.com/JulianToledano/goingecko/v3/api"
	"github.com/JulianToledano/goingecko/v3/api/coins"
)

func main() {
	cgClient := api.NewDefaultRateLimitedClient()

	for i := 0; i < 100; i++ {
		data, err := cgClient.CoinsId(context.Background(), "bitcoin", coins.WithTickers(false))
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		fmt.Printf("%d | Bitcoin price is: %f$\n", i, data.MarketData.CurrentPrice.Usd)
	}
}
