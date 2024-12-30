package goingecko

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/JulianToledano/goingecko/v2/companies"
)

func (c *Client) PublicTreasuryCoinId(ctx context.Context, id string) (*companies.Treasury, error) {
	rUrl := fmt.Sprintf("%s/public_treasury/%s", c.getCompaniesURL(), id)
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data *companies.Treasury
	err = json.Unmarshal([]byte(resp), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
