package main

import (
	"fmt"

	"github.com/JulianToledano/goingecko"
)

func main() {
	cgClient := goingecko.NewClient(nil)
	defer cgClient.Close()

	treding, err := cgClient.Trending()
	if err != nil {
		fmt.Print("Something went wrong...")
		return
	}
	for _, coin := range treding.Coins {
		coinData, err := cgClient.CoinsId(coin.Item.ID, false, false, true, false, false, false)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
		fmt.Printf("Name: %s | Symbol: %s | Price: %f\n", coin.Item.Name, coin.Item.Symbol, coinData.MarketData.CurrentPrice.Usd)
	}
}
