package goingecko

import (
	"encoding/json"
	"fmt"

	"github.com/JulianToledano/goingecko/exchangeRates"
)

func (c *Client) ExchangeRates() (*exchangeRates.Rates, error) {
	rUrl := fmt.Sprintf("%s", exchangeRatesURL)
	resp, err := c.MakeReq(rUrl)
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
