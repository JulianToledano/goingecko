package coins_test

import (
	"testing"

	"github.com/JulianToledano/goingecko"
)

func TestCoinsList(t *testing.T) {

	cgClient := goingecko.NewClient(nil)

	coinData, _ := cgClient.GetCoinsList()
	if coinData == nil {
		t.Errorf("Error")
	}

}

func TestCoinsMarket(t *testing.T) {

	cgClient := goingecko.NewClient(nil)
	ids := []string{
		"bitcoin",
		"ethereum",
	}
	priceChange := make([]string, 0)

	priceChange = append(priceChange, "24h")
	coinData, _ := cgClient.GetCoinsMarket("usd", ids, "", "", "100", "1", true, priceChange)
	if coinData == nil {
		t.Errorf("Error")
	}

}

func TestCoinsIds(t *testing.T) {

	cgClient := goingecko.NewClient(nil)

	coinData, _ := cgClient.GetCoinId("bitcoin", true, true, true, true, true, true)
	if coinData == nil {
		t.Errorf("Error")
	}

}
