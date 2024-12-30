package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/JulianToledano/goingecko/v2"
)

func TestCoins(t *testing.T) {
	coin := "bitcoin"

	cgClient := goingecko.NewClient(nil, "")

	coinData, _ := cgClient.CoinsId(context.Background(), coin, true, true, true, true, true, true)
	fmt.Println(coinData.MarketData.CurrentPrice.Usd)

}
