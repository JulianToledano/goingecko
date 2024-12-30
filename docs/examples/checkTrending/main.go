package main

import (
	"context"
	"fmt"

	"github.com/JulianToledano/goingecko/v3/api"
)

func main() {
	cgClient := api.NewDefaultClient()

	treding, err := cgClient.Trending(context.Background())
	if err != nil {
		fmt.Print("Something went wrong...")
		return
	}
	for _, coin := range treding.Coins {
		coinData, err := cgClient.CoinsId(context.Background(), coin.Item.ID)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
		fmt.Printf("Name: %s | Symbol: %s | Price: %f\n", coin.Item.Name, coin.Item.Symbol, coinData.MarketData.CurrentPrice.Usd)
	}
}
