package ping

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/JulianToledano/goingecko/v3/api/internal"
	"github.com/JulianToledano/goingecko/v3/api/ping/types"
	geckohttp "github.com/JulianToledano/goingecko/v3/http"
)

type PingClient struct {
	*internal.Client
}

func NewClient(c *geckohttp.Client, url string) *PingClient {
	return &PingClient{
		internal.NewClient(c, url),
	}
}

func (c *PingClient) Ping(ctx context.Context) (*types.Ping, error) {
	resp, err := c.MakeReq(ctx, fmt.Sprintf("%s/ping", c.URL))
	if err != nil {
		return nil, err
	}
	var data *types.Ping
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
