package publicTreasury

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/JulianToledano/goingecko/v3/api/internal"
	"github.com/JulianToledano/goingecko/v3/api/publicTreasury/types"
)

// publicTreasuryCoinIdOption is an interface that extends internal.Option to provide options specific to the public treasury coin ID endpoint.
type publicTreasuryCoinIdOption interface {
	internal.Option

	isPublicTreasuryCoinIdOption()
}

// WithPublicTreasuryPerPage returns a publicTreasuryCoinIdOption that sets the number of treasury holdings returned per page.
func WithPublicTreasuryPerPage(perPage int) publicTreasuryCoinIdOption {
	return publicTreasuryPerPageOption{perPage}
}

// WithPublicTreasuryPage returns a publicTreasuryCoinIdOption that sets the page of treasury holdings to return.
func WithPublicTreasuryPage(page int) publicTreasuryCoinIdOption {
	return publicTreasuryPageOption{page}
}

// WithPublicTreasuryOrder returns a publicTreasuryCoinIdOption that sets the sort order for treasury holdings.
// Supported values are "total_holdings_usd_desc" and "total_holdings_usd_asc".
func WithPublicTreasuryOrder(order string) publicTreasuryCoinIdOption {
	return publicTreasuryOrderOption{order}
}

// PublicTreasuryCoinId allows you to query public companies' cryptocurrency holdings by coin ID.
func (c *PublicTreasuryClient) PublicTreasuryCoinId(ctx context.Context, id string, options ...publicTreasuryCoinIdOption) (*types.Treasury, error) {
	return c.PublicTreasuryCoinIdByEntity(ctx, "companies", id, options...)
}

// PublicTreasuryCoinIdByEntity allows you to query public companies' or governments' cryptocurrency holdings by coin ID.
// Supported entity values are "companies" and "governments".
func (c *PublicTreasuryClient) PublicTreasuryCoinIdByEntity(ctx context.Context, entity, coinID string, options ...publicTreasuryCoinIdOption) (*types.Treasury, error) {
	params := url.Values{}
	for _, opt := range options {
		opt.Apply(&params)
	}

	rUrl := fmt.Sprintf("%s/%s/public_treasury/%s?%s", c.URL, entity, coinID, params.Encode())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data *types.Treasury
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type publicTreasuryPerPageOption struct{ perPage int }

func (o publicTreasuryPerPageOption) Apply(v *url.Values) {
	v.Set("per_page", strconv.Itoa(o.perPage))
}

func (o publicTreasuryPerPageOption) isPublicTreasuryCoinIdOption() {}

type publicTreasuryPageOption struct{ page int }

func (o publicTreasuryPageOption) Apply(v *url.Values) {
	v.Set("page", strconv.Itoa(o.page))
}

func (o publicTreasuryPageOption) isPublicTreasuryCoinIdOption() {}

type publicTreasuryOrderOption struct{ order string }

func (o publicTreasuryOrderOption) Apply(v *url.Values) {
	v.Set("order", o.order)
}

func (o publicTreasuryOrderOption) isPublicTreasuryCoinIdOption() {}
