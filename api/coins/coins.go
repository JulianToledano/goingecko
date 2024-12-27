package coins

//
//import (
//	"context"
//	"encoding/json"
//	"fmt"
//	"net/url"
//	"strconv"
//	"strings"
//
//	"github.com/JulianToledano/goingecko"
//
//	cointypes "github.com/JulianToledano/goingecko/api/coins/types"
//	"github.com/JulianToledano/goingecko/types"
//)
//

//func (c *Client) CoinsIdHistory(ctx context.Context, id, date string, localization bool) (*cointypes.History, error) {
//	params := url.Values{}
//
//	params.Add("date", date)
//	params.Add("localization", strconv.FormatBool(localization))
//
//	rUrl := fmt.Sprintf("%s/%s/%s?%s", c.getCoinsURL(), id, "history", params.Encode())
//	resp, err := c.MakeReq(ctx, rUrl)
//	if err != nil {
//		return nil, err
//	}
//	var data *cointypes.History
//	err = json.Unmarshal(resp, &data)
//	if err != nil {
//		return nil, err
//	}
//	return data, nil
//}
//
//func (c *Client) CoinsIdMarketChart(ctx context.Context, id, currency, days string) (*types.MarketChart, error) {
//	params := url.Values{}
//	params.Add("vs_currency", currency)
//	params.Add("days", days)
//	params.Add("interval", "daily")
//
//	rUrl := fmt.Sprintf("%s/%s/%s?%s", c.getCoinsURL(), id, "market_chart", params.Encode())
//	resp, err := c.MakeReq(ctx, rUrl)
//	if err != nil {
//		return nil, err
//	}
//	var data *types.MarketChart
//	err = json.Unmarshal(resp, &data)
//	if err != nil {
//		return nil, err
//	}
//	return data, nil
//}
//
//func (c *Client) CoinsIdMarketChartRange(ctx context.Context, id, currency, from, to string) (*types.MarketChart, error) {
//	params := url.Values{}
//	params.Add("vs_currency", currency)
//	params.Add("from", from)
//	params.Add("to", to)
//
//	rUrl := fmt.Sprintf("%s/%s/%s?%s", c.getCoinsURL(), id, "market_chart/range", params.Encode())
//	resp, err := c.MakeReq(ctx, rUrl)
//	if err != nil {
//		return nil, err
//	}
//	var data *types.MarketChart
//	err = json.Unmarshal(resp, &data)
//	if err != nil {
//		return nil, err
//	}
//	return data, nil
//}
//
//func (c *Client) CoinsOhlc(ctx context.Context, id, currency, days string) (*cointypes.Ohlc, error) {
//	params := url.Values{}
//	params.Add("vs_currency", currency)
//	params.Add("days", days)
//
//	rUrl := fmt.Sprintf("%s/%s/%s?%s", c.getCoinsURL(), id, "ohlc", params.Encode())
//	resp, err := c.MakeReq(ctx, rUrl)
//	if err != nil {
//		return nil, err
//	}
//	var data *cointypes.Ohlc
//	err = json.Unmarshal(resp, &data)
//	if err != nil {
//		return nil, err
//	}
//	return data, nil
//}
