package exchanges

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/JulianToledano/goingecko/api/exchanges/types"
)

// ExchangesId endpoint allows you to query exchange’s data (name, year established, country, .... etc), exchange volume in BTC and top 100 tickers based on exchange’s id.
//
//	📘 Notes
//
//	    The exchange volume in the response is provided in BTC. To convert it to other currencies, please use /exchange_rates endpoint
//	    For derivatives (e.g. bitmex, binance_futures), to get derivatives exchanges data, please go to /derivatives/exchange/{id} endpoint
//	    Tickers are limited to 100 items, to get more tickers, please go to /exchanges/{id}/tickers endpoint
//	    Cache / Update Frequency: every 60 seconds for all the API plans
func (c *ExchangesClient) ExchangesId(ctx context.Context, id string) (*types.ExchangeWithTicker, error) {
	rUrl := fmt.Sprintf("%s/%s", c.exchangesUrl(), id)
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data *types.ExchangeWithTicker
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
