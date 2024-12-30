package simple

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/JulianToledano/goingecko/v3/api/simple/types"
)

// SimpleTokenPrice allows you to query a token price by using token contract address.
//
//	👍 Tips
//
//	    You may obtain the asset platform and contract address via several ways:
//	        refers to respective coin page and find ‘contract’
//	        refers to /coins/list endpoint (include platform = true)
//	    You may flag to include more data such as market cap, 24hr volume, 24hr change, last updated time etc.
//
//	📘 Notes
//
//	    The endpoint returns the global average price of the coin that is aggregated across all active exchanges on CoinGecko
//	    You may cross-check the price on CoinGecko and learn more about our price methodology here
//	    Cache/Update Frequency: every 30 seconds for Pro API (Analyst, Lite, Pro, Enterprise)
func (c *SimpleClient) SimpleTokenPrice(ctx context.Context, id, contractAddresses, vsCurrencies string, options ...priceOption) (types.TokenPrice, error) {
	params := url.Values{}

	params.Add("contract_addresses", contractAddresses)
	params.Add("vs_currencies", vsCurrencies)

	for _, opt := range options {
		opt.Apply(&params)
	}

	rUrl := fmt.Sprintf("%s/%s/%s?%s", c.simpleUrl(), "token_price", id, params.Encode())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data types.TokenPrice
	err = json.Unmarshal([]byte(resp), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
