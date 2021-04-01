package goingecko

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/JulianToledano/goingecko/contract"
	"github.com/JulianToledano/goingecko/types"
)

func (c *Client) ContractInfo(id, contractAddress string) (*contract.ContractAddressInfo, error) {
	rUrl := fmt.Sprintf("%s/%s/%s/%s", contractURL, id, "contract", contractAddress)
	resp, err := c.MakeReq(rUrl)
	if err != nil {
		return nil, err
	}

	var data *contract.ContractAddressInfo
	err = json.Unmarshal([]byte(resp), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) ContractMarketChart(id, contractAddress, vsCurrency, days string) (*types.MarketChart, error) {
	params := url.Values{}
	params.Add("vs_currency", vsCurrency)
	params.Add("days", days)

	rUrl := fmt.Sprintf("%s/%s/%s/%s/%s?%s", contractURL, id, "contract", contractAddress, "market_chart", params.Encode())
	resp, err := c.MakeReq(rUrl)
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

func (c *Client) ContractMarketChartRange(id, contractAddress, vsCurrency, from, to string) (*types.MarketChart, error) {
	params := url.Values{}
	params.Add("vs_currency", vsCurrency)
	params.Add("from", from)
	params.Add("to", to)

	rUrl := fmt.Sprintf("%s/%s/%s/%s/%s?%s", contractURL, id, "contract", contractAddress, "market_chart/range", params.Encode())
	resp, err := c.MakeReq(rUrl)
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
