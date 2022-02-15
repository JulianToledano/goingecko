package goingecko

import (
	"encoding/json"
	"fmt"

	"github.com/JulianToledano/goingecko/global"
)

func (c *Client) Global() (*global.Global, error) {
	rUrl := globalURL
	resp, err := c.MakeReq(rUrl)
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

func (c *Client) DecentrilizedFinanceDEFI() (*global.Defi, error) {
	rUrl := fmt.Sprintf("%s/decentralized_finance_defi", globalURL)
	resp, err := c.MakeReq(rUrl)
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
