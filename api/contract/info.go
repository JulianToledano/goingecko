package contract

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/JulianToledano/goingecko/api/contract/types"
)

func (c *ContractClient) ContractInfo(ctx context.Context, id, contractAddress string) (*types.ContractAddressInfo, error) {
	rUrl := fmt.Sprintf("%s/%s/%s/%s", c.contractUrl(), id, "contract", contractAddress)
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data *types.ContractAddressInfo
	err = json.Unmarshal([]byte(resp), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
