package exchanges

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/JulianToledano/goingecko/v3/api/exchanges/types"
	"github.com/JulianToledano/goingecko/v3/api/internal"
)

// exchangesOption is an interface that extends internal.Option to provide options specific to the exchanges endpoint.
type exchangesOption interface {
	internal.Option

	isExchangesOptions()
}

// WithPerPageExchagesOption returns an exchangesOption that sets the number of results per page.
// The perPage parameter specifies how many exchanges to return per page in the response.
func WithPerPageExchagesOption(perPage int64) exchangesOption {
	return perPageExchangesOption{perPage: perPage}
}

// WithPageExchagesOption returns an exchangesOption that sets which page of results to return.
// The page parameter specifies which page of exchanges to return in the response.
func WithPageExchagesOption(page int64) exchangesOption {
	return pageExchangesOption{page: page}
}

// Exchanges allows you to query all the supported exchanges with exchanges’ data (id, name, country, .... etc) that have active trading volumes on CoinGecko.
//
//	👍 Tips
//
//	    You may include values such as per_page and page to specify how many results you would like to show in the responses per page and which page of responses you would like to show.
//
//	📘 Notes
//
//	    All the exchanges in the responses are the exchanges with active trading volume on CoinGecko, any inactive or deactivated exchanges will be removed from the list
//	    Cache / Update Frequency: every 60 seconds for all the API plans
func (c *ExchangesClient) Exchanges(ctx context.Context, options ...exchangesOption) (*types.ExchangesList, error) {
	params := url.Values{}
	for _, opt := range options {
		opt.Apply(&params)
	}

	rUrl := fmt.Sprintf("%s?%s", c.exchangesUrl(), params.Encode())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data *types.ExchangesList
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

type perPageExchangesOption struct{ perPage int64 }
type pageExchangesOption struct{ page int64 }

func (o perPageExchangesOption) Apply(v *url.Values) {
	v.Set("per_page", strconv.FormatInt(o.perPage, 10))
}
func (o pageExchangesOption) Apply(v *url.Values) {
	v.Set("page", strconv.FormatInt(o.page, 10))
}

func (o perPageExchangesOption) isExchangesOptions() {}
func (o pageExchangesOption) isExchangesOptions()    {}
