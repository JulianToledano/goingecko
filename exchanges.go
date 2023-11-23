package goingecko

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/JulianToledano/goingecko/exchanges"
	"net/url"
)

// Exchanges List all exchanges
// Cache / Update Frequency: every 60 seconds
// Parameters
// per_page (string) Total results per page:
//
//	Valid values: 1...250
//	Total results per page
//	Default value:: 100
//
// page (string) page through results
func (c *Client) Exchanges(perPage, page string) (*exchanges.ExchangesList, error) {
	params := url.Values{}
	if perPage != "" {
		params.Add("per_page", perPage)
	}
	if page != "" {
		params.Add("page", page)
	}

	rUrl := fmt.Sprintf("%s?%s", exchangesURL, params.Encode())
	resp, err := c.MakeReq(rUrl)
	if err != nil {
		return nil, err
	}

	var data *exchanges.ExchangesList
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// ExchangesList Use this to obtain all the markets' id in order to make API calls
// Cache / Update Frequency: every 5 minutes
func (c *Client) ExchangesList() ([]exchanges.ExchangeId, error) {
	rUrl := fmt.Sprintf("%s/list", exchangesURL)
	resp, err := c.MakeReq(rUrl)
	if err != nil {
		return nil, err
	}

	var data []exchanges.ExchangeId
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// ExchangesId Get exchange volume in BTC and tickers.
// For derivatives (e.g. bitmex, binance_futures), please use /derivatives/exchange/{id} endpoint.
//
// IMPORTANT:
// Ticker object is limited to 100 items, to get more tickers, use /exchanges/{id}/tickers
// Ticker is_stale is true when ticker that has not been updated/unchanged from the exchange for more than 8 hours.
// Ticker is_anomaly is true if ticker's price is outliered by our system.
// You are responsible for managing how you want to display these information (e.g. footnote, different background, change opacity, hide)
//
// Dictionary:
//
// last: latest unconverted price in the respective pair target currency
// volume: unconverted 24h trading volume in the respective pair target currency
// converted_last: latest converted price in BTC, ETH, and USD
// converted_volume: converted 24h trading volume in BTC, ETH, and USD
// timestamp: returns the last time that the price has changed
// last_traded_at: returns the last time that the price has changed
// last_fetch_at: returns the last time we call the API
// Cache / Update Frequency: every 60 seconds
func (c *Client) ExchangesId(id string) (*exchanges.ExchangeWithTicker, error) {
	rUrl := fmt.Sprintf("%s/%s", exchangesURL, id)
	resp, err := c.MakeReq(rUrl)
	if err != nil {
		return nil, err
	}

	var data *exchanges.ExchangeWithTicker
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// ExchangesIdTickers Get exchange tickers (paginated)
// IMPORTANT:
// Ticker is_stale is true when ticker that has not been updated/unchanged from the exchange for more than 8 hours.
// Ticker is_anomaly is true if ticker's price is outliered by our system.
// You are responsible for managing how you want to display these information (e.g. footnote, different background, change opacity, hide)
//
// Dictionary:
//
// last: latest unconverted price in the respective pair target currency
// volume: unconverted 24h trading volume in the respective pair target currency
// converted_last: latest converted price in BTC, ETH, and USD
// converted_volume: converted 24h trading volume in BTC, ETH, and USD
// timestamp: returns the last time that the price has changed
// last_traded_at: returns the last time that the price has changed
// last_fetch_at: returns the last time we call the API
// Cache / Update Frequency: every 60 seconds
// Parameters:
//
//	id* - string - pass the exchange id (can be obtained from /exchanges/list) eg. binance
//	coinIds - string - filter tickers by coin_ids (ref: v3/coins/list)
//	includeExchangeLogo - string - flag to show exchange_logo. valid values: true, false
//	page - integer - Page through results
//	depth - string - lag to show 2% orderbook depth. i.e., cost_to_move_up_usd and cost_to_move_down_usd. valid values: true, false
//	order - string - valid values: trust_score_desc (default), trust_score_asc and volume_desc
func (c *Client) ExchangesIdTickers(id, coinIds, includeExchangeLogo string, page int32, depth, order string) (*exchanges.Tickers, error) {
	params := url.Values{}
	if coinIds != "" {
		params.Add("coin_ids", coinIds)
	}
	if includeExchangeLogo != "" {
		params.Add("include_exchange_logo", includeExchangeLogo)
	}
	if page > 0 {
		params.Add("page", string(page))
	}
	if depth != "0" {
		params.Add("depth", depth)
	}
	if order != "0" {
		params.Add("order", order)
	}
	rUrl := fmt.Sprintf("%s/%s/tickers?%s", exchangesURL, id, params.Encode())
	resp, err := c.MakeReq(rUrl)
	if err != nil {
		return nil, err
	}

	var data *exchanges.Tickers
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// ExchangesIdVolumeChart Get 24 hour rolling trading volume data (in BTC) for a given exchange.
// Data granularity is automatic (cannot be adjusted)
//
//	1 day = 10-minutely
//	2-90 days = hourly
//	91 days above = daily
//
// Note: exclusive endpoint is available for paid users to query more than 1 year of historical data
//
// Cache / Update Frequency: every 60 seconds
// Parameters:
// id* - string - pass the exchange id (can be obtained from /exchanges/list) eg. binance
// days* - string - Data up to number of days ago (1/7/14/30/90/180/365)
func (c *Client) ExchangesIdVolumeChart(id, days string) ([]exchanges.Volume, error) {
	if _, ok := exchanges.VolumeChartValidDays[days]; !ok {
		return nil, errors.New("not valid day")
	}

	params := url.Values{}
	params.Add("days", days)

	rUrl := fmt.Sprintf("%s/%s/volume_chart?%s", exchangesURL, id, params.Encode())
	resp, err := c.MakeReq(rUrl)
	if err != nil {
		return nil, err
	}

	var data []exchanges.Volume
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
