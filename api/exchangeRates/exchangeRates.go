package exchangeRates

import (
	"context"
	"encoding/json"

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
	resp, err := c.MakeReq(ctx, c.exchangeRatesUrl())
	if err != nil {
		return nil, err
	}

	var data *types.Rates
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
