package goingecko

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/JulianToledano/goingecko/v2/global"
)

func (c *Client) Global(ctx context.Context) (*global.Global, error) {
	resp, err := c.MakeReq(ctx, c.getGlobalURL())
	if err != nil {
		return nil, err
	}

	var data *global.Global
	err = json.Unmarshal([]byte(resp), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) DecentralizedFinanceDEFI(ctx context.Context) (*global.Defi, error) {
	rUrl := fmt.Sprintf("%s/decentralized_finance_defi", c.getGlobalURL())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data *global.Defi
	err = json.Unmarshal([]byte(resp), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
