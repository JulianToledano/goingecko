package companies

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/JulianToledano/goingecko/api/companies/types"
)

func (c *CompaniesClient) PublicTreasuryCoinId(ctx context.Context, id string) (*types.Treasury, error) {
	rUrl := fmt.Sprintf("%s/public_treasury/%s", c.companiesUrl(), id)
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
