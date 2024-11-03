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

// CoinsMarketOption is specific to the CoinsMarket function
type CoinsMarketOption interface {
	Option
	isCoinsMarketOption()
}

// Define option types
type idsOption struct{ ids []string }
type categoryOption struct{ category string }
type orderOption struct{ order string }
type perPageOption struct{ perPage int }
type pageOption struct{ page int }
type sparklineOption struct{ sparkline bool }
type priceChangePercentageOption struct{ intervals []string }

// Implement Option interface
func (o idsOption) apply(v *url.Values)       { v.Set("ids", strings.Join(o.ids, ",")) }
func (o categoryOption) apply(v *url.Values)  { v.Set("category", o.category) }
func (o orderOption) apply(v *url.Values)     { v.Set("order", o.order) }
func (o perPageOption) apply(v *url.Values)   { v.Set("per_page", strconv.Itoa(o.perPage)) }
func (o pageOption) apply(v *url.Values)      { v.Set("page", strconv.Itoa(o.page)) }
func (o sparklineOption) apply(v *url.Values) { v.Set("sparkline", strconv.FormatBool(o.sparkline)) }
func (o priceChangePercentageOption) apply(v *url.Values) {
	v.Set("price_change_percentage", strings.Join(o.intervals, ","))
}

// Implement CoinsMarketOption interface
func (idsOption) isCoinsMarketOption()                   {}
func (categoryOption) isCoinsMarketOption()              {}
func (orderOption) isCoinsMarketOption()                 {}
func (perPageOption) isCoinsMarketOption()               {}
func (pageOption) isCoinsMarketOption()                  {}
func (sparklineOption) isCoinsMarketOption()             {}
func (priceChangePercentageOption) isCoinsMarketOption() {}

// Option constructors
func WithIDs(ids []string) CoinsMarketOption         { return idsOption{ids} }
func WithCategory(category string) CoinsMarketOption { return categoryOption{category} }
func WithOrder(order string) CoinsMarketOption       { return orderOption{order} }
func WithPerPage(perPage int) CoinsMarketOption      { return perPageOption{perPage} }
func WithPage(page int) CoinsMarketOption            { return pageOption{page} }
func WithSparkline(sparkline bool) CoinsMarketOption { return sparklineOption{sparkline} }
func WithPriceChangePercentage(intervals []string) CoinsMarketOption {
	return priceChangePercentageOption{intervals}
}

func (c *Client) CoinsMarket(currency string, options ...CoinsMarketOption) ([]*coins.Market, error) {
	params := url.Values{}
	params.Add("vs_currency", currency)

	// Apply all the options
	for _, opt := range options {
		opt.apply(&params)
	}

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

// CoinsIdOption is specific to the CoinsId function
type CoinsIdOption interface {
	Option
	isCoinsIdOption()
}

// Define option types
type localizationOption struct{ localization bool }
type tickersOption struct{ tickers bool }
type marketDataOption struct{ marketData bool }
type communityDataOption struct{ communityData bool }
type developerDataOption struct{ developerData bool }
type coinSparklineOption struct{ sparkline bool }

// Implement Option interface
func (o localizationOption) apply(v *url.Values) {
	v.Set("localization", strconv.FormatBool(o.localization))
}
func (o tickersOption) apply(v *url.Values) { v.Set("tickers", strconv.FormatBool(o.tickers)) }
func (o marketDataOption) apply(v *url.Values) {
	v.Set("market_data", strconv.FormatBool(o.marketData))
}
func (o communityDataOption) apply(v *url.Values) {
	v.Set("community_data", strconv.FormatBool(o.communityData))
}
func (o developerDataOption) apply(v *url.Values) {
	v.Set("developer_data", strconv.FormatBool(o.developerData))
}
func (o coinSparklineOption) apply(v *url.Values) {
	v.Set("sparkline", strconv.FormatBool(o.sparkline))
}

// Implement CoinsIdOption interface
func (localizationOption) isCoinsIdOption()  {}
func (tickersOption) isCoinsIdOption()       {}
func (marketDataOption) isCoinsIdOption()    {}
func (communityDataOption) isCoinsIdOption() {}
func (developerDataOption) isCoinsIdOption() {}
func (coinSparklineOption) isCoinsIdOption() {}

// Option constructors
func WithLocalization(localization bool) CoinsIdOption { return localizationOption{localization} }
func WithTickers(tickers bool) CoinsIdOption           { return tickersOption{tickers} }
func WithMarketData(marketData bool) CoinsIdOption     { return marketDataOption{marketData} }
func WithCommunityData(communityData bool) CoinsIdOption {
	return communityDataOption{communityData}
}
func WithDeveloperData(developerData bool) CoinsIdOption {
	return developerDataOption{developerData}
}
func WithCoinSparkline(sparkline bool) CoinsIdOption { return coinSparklineOption{sparkline} }

func (c *Client) CoinsId(id string, options ...CoinsIdOption) (*coins.CoinID, error) {
	params := url.Values{}

	// Apply all the options
	for _, opt := range options {
		opt.apply(&params)
	}

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

// CoinsIdTickersOption is specific to the CoinsIdTickers function
type CoinsIdTickersOption interface {
	Option
	isCoinsIdTickersOption()
}

// Define option types
type includeExchangeLogoOption struct{ includeLogo bool }
type depthOption struct{ depth string }
type exchangeIdOption struct{ exchangeId string }

// Implement Option interface
func (o includeExchangeLogoOption) apply(v *url.Values) {
	v.Set("include_exchange_logo", strconv.FormatBool(o.includeLogo))
}
func (o depthOption) apply(v *url.Values)      { v.Set("depth", o.depth) }
func (o exchangeIdOption) apply(v *url.Values) { v.Set("exchange_ids", o.exchangeId) }

// Implement CoinsIdTickersOption interface
func (includeExchangeLogoOption) isCoinsIdTickersOption() {}
func (orderOption) isCoinsIdTickersOption()               {}
func (depthOption) isCoinsIdTickersOption()               {}
func (pageOption) isCoinsIdTickersOption()                {}
func (exchangeIdOption) isCoinsIdTickersOption()          {}

// Option constructors
func WithIncludeExchangeLogo(includeLogo bool) CoinsIdTickersOption {
	return includeExchangeLogoOption{includeLogo}
}
func WithDepth(depth string) CoinsIdTickersOption           { return depthOption{depth} }
func WithExchangeId(exchangeId string) CoinsIdTickersOption { return exchangeIdOption{exchangeId} }

func (c *Client) CoinsIdTickers(id string, options ...CoinsIdTickersOption) (*coins.Tickers, error) {
	params := url.Values{}

	// Apply all the options
	for _, opt := range options {
		opt.apply(&params)
	}

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

func (c *Client) CoinsIdMarketChart(id, currency, days string) (*types.MarketChart, error) {
	params := url.Values{}
	params.Add("vs_currency", currency)
	params.Add("days", days)
	params.Add("interval", "daily")

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
