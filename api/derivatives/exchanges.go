package derivatives

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/JulianToledano/goingecko/v3/api"

	"github.com/JulianToledano/goingecko/v3/api/derivatives/types"
)

// exchangesOption is an interface that extends api.Option to provide options specific to the derivatives exchanges endpoint.
type exchangesOption interface {
	api.Option

	isExchangesOption()
}

// WithOrderDerivativesExchanges returns an exchangesOption that sets the ordering of results.
// The order parameter specifies how to sort the derivatives exchanges in the response.
func WithOrderDerivativesExchanges(order string) exchangesOption {
	return orderExchangeOption{order}
}

// WithPerPageDerivativesExchanges returns an exchangesOption that sets the number of results per page.
// The perPage parameter specifies how many derivatives exchanges to return per page in the response.
func WithPerPageDerivativesExchanges(perPage int64) exchangesOption {
	return perPageExchangeOption{perPage}
}

// WithPageDerivativesExchanges returns an exchangesOption that sets which page of results to return.
// The page parameter specifies which page of derivatives exchanges to return in the response.
func WithPageDerivativesExchanges(page int64) exchangesOption {
	return pageExchangeOption{page}
}

// DerivativesExchanges allows you to query all the derivatives exchanges with related data (id, name, open interest, .... etc) on CoinGecko.
//
//	👍 Tips
//
//	    You may include values such as per_page and page to specify how many results you would like to show in the responses per page and which page of responses you would like to show
//
//	📘Notes
//
//	    Cache / Update Frequency: every 60 seconds for all the API plans
func (c *DerivativesClient) DerivativesExchanges(ctx context.Context, options ...exchangesOption) ([]types.Exchange, error) {
	params := url.Values{}
	for _, opt := range options {
		opt.Apply(&params)
	}

	rUrl := fmt.Sprintf("%s/exchanges?%s", c.derivativesUrl(), params.Encode())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data []types.Exchange
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

type orderExchangeOption struct{ order string }
type perPageExchangeOption struct{ perPage int64 }
type pageExchangeOption struct{ page int64 }

func (o orderExchangeOption) Apply(v *url.Values) {
	v.Set("order", o.order)
}
func (o perPageExchangeOption) Apply(v *url.Values) {
	v.Set("per_page", strconv.FormatInt(o.perPage, 10))
}
func (o pageExchangeOption) Apply(v *url.Values) {
	v.Set("page", strconv.FormatInt(o.page, 10))
}

func (o orderExchangeOption) isExchangesOption()   {}
func (o perPageExchangeOption) isExchangesOption() {}
func (o pageExchangeOption) isExchangesOption()    {}
