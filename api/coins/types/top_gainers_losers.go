package types

import "encoding/json"

type TopGainersLosers struct {
	TopGainers []*TopGainersLosersCoin `json:"top_gainers"`
	TopLosers  []*TopGainersLosersCoin `json:"top_losers"`
}

type TopGainersLosersCoin struct {
	ID            string              `json:"id"`
	Symbol        string              `json:"symbol"`
	Name          string              `json:"name"`
	Image         string              `json:"image"`
	MarketCapRank int64               `json:"market_cap_rank"`
	CurrencyData  map[string]*float64 `json:"-"`
	Usd           float64             `json:"usd"`
	Usd24hVol     float64             `json:"usd_24h_vol"`
	Usd24hChange  float64             `json:"usd_24h_change"`
	Usd1hChange   *float64            `json:"usd_1h_change"`
	Usd7dChange   *float64            `json:"usd_7d_change"`
	Usd14dChange  *float64            `json:"usd_14d_change"`
	Usd30dChange  *float64            `json:"usd_30d_change"`
	Usd60dChange  *float64            `json:"usd_60d_change"`
	Usd200dChange *float64            `json:"usd_200d_change"`
	Usd1yChange   *float64            `json:"usd_1y_change"`
}

func (c *TopGainersLosersCoin) UnmarshalJSON(data []byte) error {
	type topGainersLosersCoin TopGainersLosersCoin

	var alias topGainersLosersCoin
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}

	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	for _, key := range []string{"id", "symbol", "name", "image", "market_cap_rank"} {
		delete(raw, key)
	}

	*c = TopGainersLosersCoin(alias)
	c.CurrencyData = make(map[string]*float64, len(raw))
	for key, value := range raw {
		if string(value) == "null" {
			c.CurrencyData[key] = nil
			continue
		}

		var number float64
		if err := json.Unmarshal(value, &number); err != nil {
			continue
		}
		c.CurrencyData[key] = &number
	}

	return nil
}
