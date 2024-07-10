package goingecko

import (
	"encoding/json"
	"fmt"

	"github.com/JulianToledano/goingecko/companies"
)

func (c *Client) PublicTreasuryCoinId(id string) (*companies.Treasury, error) {
	rUrl := fmt.Sprintf("%s/public_treasury/%s", c.getCompaniesURL(), id)
	resp, err := c.MakeReq(rUrl)
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
