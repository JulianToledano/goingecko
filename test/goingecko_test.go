package test

import (
	"fmt"
	"testing"

	"github.com/JulianToledano/goingecko"
)

func TestCoins(t *testing.T) {
	coin := "bitcoin"

	cgClient := goingecko.NewClient(nil)

	coinData, _ := cgClient.GetCoinsId(coin, true, true, true, true, true, true)
	fmt.Println(coinData.MarketData.CurrentPrice.Usd)

}
