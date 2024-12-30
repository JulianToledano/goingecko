package exchanges

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/JulianToledano/goingecko/api"

	"github.com/JulianToledano/goingecko/api/exchanges/types"
)

// listOption is an interface that extends api.Option to provide options specific to the exchanges list endpoint.
type listOption interface {
	api.Option

	isListOption()
}

// WithStatus returns a listOption that filters exchanges by their status.
// The status parameter can be used to filter exchanges based on their current operational status.
func WithStatus(status string) listOption {
	return statusOption{status}
}

// ExchangesList allows you to query all the supported exchanges with exchanges’ data (id, name, country, .... etc) that have active trading volumes on CoinGecko.
//
//	👍 Tips
//
//	    You may include values such as per_page and page to specify how many results you would like to show in the responses per page and which page of responses you would like to show.
//
//	📘 Notes
//
//	    All the exchanges in the responses are the exchanges with active trading volume on CoinGecko, any inactive or deactivated exchanges will be removed from the list
//	    Cache / Update Frequency: every 60 seconds for all the API plans
func (c *ExchangesClient) ExchangesList(ctx context.Context, options ...listOption) ([]types.ExchangeId, error) {
	params := url.Values{}
	for _, opt := range options {
		opt.Apply(&params)
	}

	rUrl := fmt.Sprintf("%s/list", c.exchangesUrl())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data []types.ExchangeId
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

type statusOption struct{ status string }

func (o statusOption) Apply(v *url.Values) {
	v.Set("status", o.status)
}

func (o statusOption) isListOption() {}
