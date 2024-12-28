package contract

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/JulianToledano/goingecko/types"
	"net/url"
)

// ContractMarketChart allows you to get the historical chart data including time in UNIX, price, market cap and 24hrs volume based on asset platform and particular token contract address.
//
//	👍 Tips
//
//	    You may obtain the asset platform and contract address via several ways:
//	        refers to respective coin page and find ‘contract’
//	        refers to /coins/list endpoint (include platform = true)
//
//	📘Notes
//
//	    You may leave the interval params as empty for automatic granularity:
//	        1 day from current time = 5-minutely data
//	        2 - 90 days from current time = hourly data
//	        above 90 days from current time = daily data (00:00 UTC)
//	    For non-Enterprise plan subscribers who would like to get hourly data, please leave the interval params empty for auto granularity
//	    The 5-minutely and hourly interval params are also exclusively available to Enterprise plan subscribers, bypassing auto-granularity:
//	        interval=5m: 5-minutely historical data (responses include information from the past 10 days, up until 2 days ago)
//	        interval=hourly: hourly historical data (responses include information from the past 100 days, up until now)
//	    Cache / Update Frequency:
//	        every 5 minutes for all the API plans
//	        The last completed UTC day (00:00) is available 35 minutes after midnight on the next UTC day (00:35). The cache will always expire at 00:40 UTC
func (c *ContractClient) ContractMarketChart(ctx context.Context, id, contractAddress, vsCurrency, days string, options ...contractOption) (*types.MarketChart, error) {
	params := url.Values{}
	params.Add("vs_currency", vsCurrency)
	params.Add("days", days)

	for _, opt := range options {
		opt.Apply(&params)
	}

	rUrl := fmt.Sprintf("%s/%s/%s/%s/%s?%s", c.contractUrl(), id, "contract", contractAddress, "market_chart", params.Encode())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data *types.MarketChart
	err = json.Unmarshal([]byte(resp), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// ContractMarketChartRange allows you to get the historical chart data within certain time range in UNIX along with
// price, market cap and 24hrs volume based on asset platform and particular token contract address.
//
//	👍Tips
//
//	    You may obtain the asset platform and contract address via several ways:
//	        refers to respective coin page and find ‘contract’
//	        refers to /coins/list endpoint (include platform = true)
//
//	📘Notes
//
//	    You may leave the interval params as empty for automatic granularity:
//	        1 day from current time = 5-minutely data
//	        1 day from any time (except current time) = hourly data
//	        2 - 90 days from any time = hourly data
//	        above 90 days from any time = daily data (00:00 UTC)
//	    For non-Enterprise plan subscribers who would like to get hourly data, please leave the interval params empty for auto granularity
//	    The 5-minutely and hourly interval params are also exclusively available to Enterprise plan subscribers, bypassing auto-granularity:
//	        interval=5m: 5-minutely historical data (responses include information from the past 10 days, up until 2 days ago)
//	        interval=hourly: hourly historical data (responses include information from the past 100 days, up until now)
//	    Data availability:
//	        interval=5m: Available from 9 February 2018 onwards
//	        interval=hourly: Available from 30 Jan 2018 onwards
//	    Cache / Update Frequency:
//	    Based on days range (all the API plans)
//	        1 day = 30 seconds cache
//	        2 -90 days = 30 minutes cache
//	        90 days = 12 hours cache
//	        The last completed UTC day (00:00) is available 35 minutes after midnight on the next UTC day (00:35). The cache will always expire at 00:40 UTC
func (c *ContractClient) ContractMarketChartRange(ctx context.Context, id, contractAddress, vsCurrency, from, to string, options ...contractOption) (*types.MarketChart, error) {
	params := url.Values{}
	params.Add("vs_currency", vsCurrency)
	params.Add("from", from)
	params.Add("to", to)

	for _, opt := range options {
		opt.Apply(&params)
	}

	rUrl := fmt.Sprintf("%s/%s/%s/%s/%s?%s", c.contractUrl(), id, "contract", contractAddress, "market_chart/range", params.Encode())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data *types.MarketChart
	err = json.Unmarshal([]byte(resp), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
