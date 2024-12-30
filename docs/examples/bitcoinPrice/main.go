package main

import (
	"context"
	"fmt"

	"github.com/JulianToledano/goingecko/v3/api"
	"github.com/JulianToledano/goingecko/v3/api/coins"
)

func main() {
	cgClient := api.NewDefaultClient()

	data, err := cgClient.CoinsId(context.Background(), "bitcoin", coins.WithTickers(false))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Bitcoin price is: %f$", data.MarketData.CurrentPrice.Usd)
}
