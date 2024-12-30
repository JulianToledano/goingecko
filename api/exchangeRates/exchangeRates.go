package exchangeRates

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/JulianToledano/goingecko/v3/api/exchangeRates/types"
)

// ExchangeRates allows you to query BTC exchange rates with other currencies.
//
//	👍 Tips
//
//	    You may use this endpoint to convert the response data, which is originally in BTC, to other currencies
//
//	📘Notes
//
//	    Cache / Update Frequency: every 5 minutes for all the API plans
func (c *ExchangeRatesClient) ExchangeRates(ctx context.Context) (*types.Rates, error) {
	rUrl := fmt.Sprintf("%s", c.exchangeRatesUrl())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data *types.Rates
	err = json.Unmarshal([]byte(resp), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
