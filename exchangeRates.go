package goingecko

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/JulianToledano/goingecko/exchangeRates"
)

func (c *Client) ExchangeRates(ctx context.Context) (*exchangeRates.Rates, error) {
	rUrl := fmt.Sprintf("%s", c.getExchangeRatesURL())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data *exchangeRates.Rates
	err = json.Unmarshal([]byte(resp), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
