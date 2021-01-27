package goingecko

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/JulianToledano/goingecko/coins"
	"github.com/JulianToledano/goingecko/ping"
)

type Client struct {
	httpClient *http.Client
	baseUrl    string
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &Client{
		httpClient: httpClient,
		baseUrl:    baseURL,
	}
}

func doReq(req *http.Request, client *http.Client) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}

// MakeReq HTTP request helper
func (c *Client) MakeReq(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	resp, err := doReq(req, c.httpClient)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Ping /ping endpoint
func (c *Client) Ping() (*ping.Ping, error) {
	resp, err := c.MakeReq(pingURL)
	if err != nil {
		return nil, err
	}
	var data *ping.Ping
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) GetCoinsList() ([]*coins.CoinInfo, error) {
	url := fmt.Sprintf("%s/%s", coinsURL, "list")
	resp, err := c.MakeReq(url)
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

func (c *Client) GetCoinsMarket(currency string, ids []string, category string, order string, perPage, page string, sparkline bool, priceChange []string) ([]*coins.Market, error) {
	params := url.Values{}
	idsParam := strings.Join(ids[:], ",")
	pChange := strings.Join(priceChange[:], ",")
	// vsCurrenciesParam := strings.Join(vsCurrencies[:], ",")

	params.Add("vs_currency", currency)
	params.Add("ids", idsParam)
	params.Add("category", category)
	params.Add("order", order)
	params.Add("per_page", perPage)
	params.Add("page", page)
	params.Add("sparkline", strconv.FormatBool(sparkline))
	params.Add("price_change_percentage", pChange)

	url := fmt.Sprintf("%s/%s?%s", coinsURL, "markets", params.Encode())
	resp, err := c.MakeReq(url)
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

func (c *Client) GetCoinsId(id string, localization, tickers, marketData, communityData, developerData, sparkline bool) (*coins.CoinID, error) {
	params := url.Values{}

	params.Add("localization", strconv.FormatBool(localization))
	params.Add("tickers", strconv.FormatBool(tickers))
	params.Add("market_data", strconv.FormatBool(marketData))
	params.Add("community_data", strconv.FormatBool(communityData))
	params.Add("developer_data", strconv.FormatBool(developerData))
	params.Add("sparkline", strconv.FormatBool(sparkline))

	url := fmt.Sprintf("%s/%s?%s", coinsURL, id, params.Encode())
	resp, err := c.MakeReq(url)
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

func (c *Client) GetCoinsIdTickers(id, exchangeId, includeExchangeLogo, page, order, depth string) (*coins.Tickers, error) {
	params := url.Values{}

	params.Add("exchange_ids", exchangeId)
	params.Add("include_exchange_logo", includeExchangeLogo)
	params.Add("page", page)
	params.Add("order", order)
	params.Add("depth", depth)

	url := fmt.Sprintf("%s/%s/%s?%s", coinsURL, id, "tickers", params.Encode())
	resp, err := c.MakeReq(url)
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

func (c *Client) GetCoinsIdHistory(id, date string, localization bool) (*coins.History, error) {
	params := url.Values{}

	params.Add("date", date)
	params.Add("localization", strconv.FormatBool(localization))

	url := fmt.Sprintf("%s/%s/%s?%s", coinsURL, id, "history", params.Encode())
	resp, err := c.MakeReq(url)
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

func (c *Client) GetCoinsIdMarketChart(id, currency, days string) (*coins.MarketChart, error) {
	params := url.Values{}
	params.Add("vs_currency", currency)
	params.Add("days", days)
	params.Add("interval", "daily")

	url := fmt.Sprintf("%s/%s/%s?%s", coinsURL, id, "market_chart", params.Encode())
	resp, err := c.MakeReq(url)
	if err != nil {
		return nil, err
	}
	var data *coins.MarketChart
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) GetCoinsIdMarketChartRange(id, currency, from, to string) (*coins.MarketChart, error) {
	params := url.Values{}
	params.Add("vs_currency", currency)
	params.Add("from", from)
	params.Add("to", to)

	url := fmt.Sprintf("%s/%s/%s?%s", coinsURL, id, "market_chart/range", params.Encode())
	resp, err := c.MakeReq(url)
	if err != nil {
		return nil, err
	}
	var data *coins.MarketChart
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) GetCoinsOhlc(id, currency, days string) (*coins.Ohlc, error) {
	params := url.Values{}
	params.Add("vs_currency", currency)
	params.Add("days", days)

	url := fmt.Sprintf("%s/%s/%s?%s", coinsURL, id, "ohlc", params.Encode())
	resp, err := c.MakeReq(url)
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
