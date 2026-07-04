package publicTreasury

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/JulianToledano/goingecko/v3/api/internal"
	"github.com/JulianToledano/goingecko/v3/api/publicTreasury/types"
)

// entityOption is an interface that extends internal.Option to provide options specific to the public treasury entity endpoint.
type entityOption interface {
	internal.Option

	IsEntityOption()
}

// WithHoldingAmountChange returns an entityOption that includes holding amount changes for the specified comma-separated timeframes.
func WithHoldingAmountChange(holdingAmountChange string) entityOption {
	return holdingAmountChangeOption{holdingAmountChange}
}

// WithHoldingChangePercentage returns an entityOption that includes holding change percentages for the specified comma-separated timeframes.
func WithHoldingChangePercentage(holdingChangePercentage string) entityOption {
	return holdingChangePercentageOption{holdingChangePercentage}
}

// PublicTreasuryEntity allows you to query public companies' or governments' cryptocurrency holdings by entity ID.
func (c *PublicTreasuryClient) PublicTreasuryEntity(ctx context.Context, entityID string, options ...entityOption) (*types.EntityTreasury, error) {
	params := url.Values{}
	for _, opt := range options {
		opt.Apply(&params)
	}

	rUrl := fmt.Sprintf("%s/public_treasury/%s?%s", c.URL, entityID, params.Encode())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data *types.EntityTreasury
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type holdingAmountChangeOption struct{ holdingAmountChange string }

func (o holdingAmountChangeOption) Apply(v *url.Values) {
	v.Set("holding_amount_change", o.holdingAmountChange)
}

func (o holdingAmountChangeOption) IsEntityOption() {}

type holdingChangePercentageOption struct{ holdingChangePercentage string }

func (o holdingChangePercentageOption) Apply(v *url.Values) {
	v.Set("holding_change_percentage", o.holdingChangePercentage)
}

func (o holdingChangePercentageOption) IsEntityOption() {}
