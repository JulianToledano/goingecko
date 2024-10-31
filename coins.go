package goingecko

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/JulianToledano/goingecko/coins"
	"github.com/JulianToledano/goingecko/types"
)

func (c *Client) CoinsList() ([]*coins.CoinInfo, error) {
	rUrl := fmt.Sprintf("%s/%s", c.getCoinsURL(), "list")
	resp, err := c.MakeReq(rUrl)
	if err != nil {
		return nil, err
	}

	var data []*coins.CoinInfo
	err = json.Unmarshal([]byte(resp), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) CoinsMarket(currency string, ids []string, category string, order string, perPage, page string, sparkline bool, priceChange []string) ([]*coins.Market, error) {
	params := url.Values{}
	idsParam := strings.Join(ids[:], ",")
	pChange := strings.Join(priceChange[:], ",")
	// vsCurrenciesParam := strings.Join(vsCurrencies[:], ",")

	params.Add("vs_currency", currency)
	params.Add("ids", idsParam)
	if category != "" {
		params.Add("category", category)
	}
	params.Add("order", order)
	params.Add("per_page", perPage)
	params.Add("page", page)
	params.Add("sparkline", strconv.FormatBool(sparkline))
	params.Add("price_change_percentage", pChange)

	rUrl := fmt.Sprintf("%s/%s?%s", c.getCoinsURL(), "markets", params.Encode())
	resp, err := c.MakeReq(rUrl)
	if err != nil {
		return nil, err
	}

	var data []*coins.Market
	err = json.Unmarshal([]byte(resp), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) CoinsId(id string, localization, tickers, marketData, communityData, developerData, sparkline bool) (*coins.CoinID, error) {
	params := url.Values{}

	params.Add("localization", strconv.FormatBool(localization))
	params.Add("tickers", strconv.FormatBool(tickers))
	params.Add("market_data", strconv.FormatBool(marketData))
	params.Add("community_data", strconv.FormatBool(communityData))
	params.Add("developer_data", strconv.FormatBool(developerData))
	params.Add("sparkline", strconv.FormatBool(sparkline))

	rUrl := fmt.Sprintf("%s/%s?%s", c.getCoinsURL(), id, params.Encode())
	resp, err := c.MakeReq(rUrl)
	if err != nil {
		return nil, err
	}
	var data *coins.CoinID
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) CoinsIdTickers(id, exchangeId, includeExchangeLogo, page, order, depth string) (*coins.Tickers, error) {
	params := url.Values{}

	params.Add("exchange_ids", exchangeId)
	params.Add("include_exchange_logo", includeExchangeLogo)
	params.Add("page", page)
	params.Add("order", order)
	params.Add("depth", depth)

	rUrl := fmt.Sprintf("%s/%s/%s?%s", c.getCoinsURL(), id, "tickers", params.Encode())
	resp, err := c.MakeReq(rUrl)
	if err != nil {
		return nil, err
	}
	var data *coins.Tickers
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) CoinsIdHistory(id, date string, localization bool) (*coins.History, error) {
	params := url.Values{}

	params.Add("date", date)
	params.Add("localization", strconv.FormatBool(localization))

	rUrl := fmt.Sprintf("%s/%s/%s?%s", c.getCoinsURL(), id, "history", params.Encode())
	resp, err := c.MakeReq(rUrl)
	if err != nil {
		return nil, err
	}
	var data *coins.History
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) CoinsIdMarketChart(id, currency, days, interval string) (*types.MarketChart, error) {
	params := url.Values{}
	params.Add("vs_currency", currency)
	params.Add("days", days)

	if interval != "" {
		params.Add("interval", interval)
	}

	rUrl := fmt.Sprintf("%s/%s/%s?%s", c.getCoinsURL(), id, "market_chart", params.Encode())
	resp, err := c.MakeReq(rUrl)
	if err != nil {
		return nil, err
	}
	var data *types.MarketChart
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) CoinsIdMarketChartRange(id, currency, from, to string) (*types.MarketChart, error) {
	params := url.Values{}
	params.Add("vs_currency", currency)
	params.Add("from", from)
	params.Add("to", to)

	rUrl := fmt.Sprintf("%s/%s/%s?%s", c.getCoinsURL(), id, "market_chart/range", params.Encode())
	resp, err := c.MakeReq(rUrl)
	if err != nil {
		return nil, err
	}
	var data *types.MarketChart
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) CoinsOhlc(id, currency, days string) (*coins.Ohlc, error) {
	params := url.Values{}
	params.Add("vs_currency", currency)
	params.Add("days", days)

	rUrl := fmt.Sprintf("%s/%s/%s?%s", c.getCoinsURL(), id, "ohlc", params.Encode())
	resp, err := c.MakeReq(rUrl)
	if err != nil {
		return nil, err
	}
	var data *coins.Ohlc
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
