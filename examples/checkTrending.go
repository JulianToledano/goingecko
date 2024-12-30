package main

import (
	"context"
	"fmt"

	"github.com/JulianToledano/goingecko/v2"
)

func main() {
	cgClient := goingecko.NewClient(nil, "")
	defer cgClient.Close()

	treding, err := cgClient.Trending()
	if err != nil {
		fmt.Print("Something went wrong...")
		return
	}
	for _, coin := range treding.Coins {
		coinData, err := cgClient.CoinsId(context.Background(), coin.Item.ID, false, false, true, false, false, false)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
		fmt.Printf("Name: %s | Symbol: %s | Price: %f\n", coin.Item.Name, coin.Item.Symbol, coinData.MarketData.CurrentPrice.Usd)
	}
}
