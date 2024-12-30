package main

import (
	"context"
	"fmt"

	"github.com/JulianToledano/goingecko/v2"
)

func main() {
	cgClient := goingecko.NewClient(nil, "pro api key", true)
	defer cgClient.Close()

	data, err := cgClient.CoinsId(context.Background(), "bitcoin", true, true, true, false, false, false)
	if err != nil {
		fmt.Print("Somethig went wrong...")
		return
	}
	fmt.Printf("Bitcoin price is: %f$", data.MarketData.CurrentPrice.Usd)
}
