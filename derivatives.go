package goingecko

import (
	"encoding/json"
	"fmt"
	"github.com/JulianToledano/goingecko/derivatives"
	"net/url"
)

// Derivatives List all derivative tickers.
// Note: 'open_interest' and 'volume_24h' data are in USD
// Cache / Update Frequency: every 30 seconds
func (c *Client) Derivatives() ([]derivatives.Derivative, error) {

	rUrl := fmt.Sprintf("%s", derivativesURL)
	resp, err := c.MakeReq(rUrl)
	if err != nil {
		return nil, err
	}

	var data []derivatives.Derivative
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// DerivativesExchanges List all derivative exchanges.
//
// Cache / Update Frequency: every 60 seconds
// Parameters:
//
//	order(string) - order results using following params name_asc，name_desc，open_interest_btc_asc，
//					open_interest_btc_desc，trade_volume_24h_btc_asc，trade_volume_24h_btc_desc
//
// perPage(integer) - Total results per page
// page(integer) - Page through results
func (c *Client) DerivativesExchanges(order string, perPage, page int32) ([]derivatives.Exchange, error) {
	params := url.Values{}
	if order != "" {
		params.Add("order", order)
	}
	if perPage > 0 {
		params.Add("per_page", string(perPage))
	}
	if page > 0 {
		params.Add("page", string(page))
	}

	rUrl := fmt.Sprintf("%s/exchanges?%s", derivativesURL, params.Encode())
	resp, err := c.MakeReq(rUrl)
	if err != nil {
		return nil, err
	}

	var data []derivatives.Exchange
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// DerivativesExchangesId show derivative exchange data
//
// Dictionary:
//
//	last: latest unconverted price in the respective pair target currency
//	volume: unconverted 24h trading volume in the respective pair target currency
//	converted_last: latest converted price in BTC, ETH, and USD
//	converted_volume: converted 24h trading volume in BTC, ETH, and USD
//
// Cache / Update Frequency: every 30 seconds
// Parameters:
//
//	id*(string) - pass the exchange id (can be obtained from derivatives/exchanges/list) eg. bitmex
//	includeTickers(string) - ['all', 'unexpired'] - expired to show unexpired tickers, all to list all tickers,
//							 leave blank to omit tickers data in response
func (c *Client) DerivativesExchangesId(id, includeTickers string) (*derivatives.ExchangeId, error) {
	params := url.Values{}
	if includeTickers != "" {
		params.Add("include_tickers", includeTickers)
	}

	rUrl := fmt.Sprintf("%s/exchanges/%s?%s", derivativesURL, id, params.Encode())
	resp, err := c.MakeReq(rUrl)
	if err != nil {
		return nil, err
	}

	var data *derivatives.ExchangeId
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// DerivativesExchangesList List all derivative exchanges name and identifier.
//
// Cache / Update Frequency: every 5 minutes
func (c *Client) DerivativesExchangesList() ([]derivatives.DerivativesListItem, error) {
	rUrl := fmt.Sprintf("%s/exchanges/list", derivativesURL)
	resp, err := c.MakeReq(rUrl)
	if err != nil {
		return nil, err
	}

	var data []derivatives.DerivativesListItem
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
