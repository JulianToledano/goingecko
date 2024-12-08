package test

import (
	"context"
	"testing"

	"github.com/JulianToledano/goingecko"
)

func TestCoinsList(t *testing.T) {

	cgClient := goingecko.NewClient(nil, "")

	coinData, _ := cgClient.CoinsList(context.Background())
	if coinData == nil {
		t.Errorf("Error")
	}
}

func TestCoinsMarket(t *testing.T) {

	cgClient := goingecko.NewClient(nil, "")
	ids := []string{
		"bitcoin",
		"ethereum",
	}
	priceChange := make([]string, 0)

	priceChange = append(priceChange, "24h")
	coinData, _ := cgClient.CoinsMarket(context.Background(), "usd", ids, "", "", "100", "1", true, priceChange)
	if coinData == nil {
		t.Errorf("Error")
	}
}

func TestCoinsIds(t *testing.T) {

	cgClient := goingecko.NewClient(nil, "")

	coinData, err := cgClient.CoinsId(context.Background(), "bitcoin", true, true, true, true, true, true)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if coinData == nil {
		t.Errorf("Error")
	}

}

func TestCoinsIdsTickers(t *testing.T) {

	cgClient := goingecko.NewClient(nil, "")

	coinData, err := cgClient.CoinsIdTickers(context.Background(), "bitcoin", "", "", "", "", "")
	if coinData == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestCoinsIdsHistory(t *testing.T) {

	cgClient := goingecko.NewClient(nil, "")

	coinData, err := cgClient.CoinsIdHistory(context.Background(), "bitcoin", "30-12-17", true)
	if coinData == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestCoinsIdsMarketChart(t *testing.T) {

	cgClient := goingecko.NewClient(nil, "")

	coinData, err := cgClient.CoinsIdMarketChart(context.Background(), "bitcoin", "usd", "10", "daily")
	if coinData == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestCoinsIdsMarketChartRange(t *testing.T) {

	cgClient := goingecko.NewClient(nil, "")

	coinData, err := cgClient.CoinsIdMarketChartRange(context.Background(), "bitcoin", "usd", "1392577232", "1422577232")
	if coinData == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestCoinsIdsOhlc(t *testing.T) {

	cgClient := goingecko.NewClient(nil, "")

	coinData, err := cgClient.CoinsOhlc(context.Background(), "bitcoin", "usd", "7")
	if coinData == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
