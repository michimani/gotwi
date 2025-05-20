package upload

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/media/upload/types"
)

const (
	initializeEndpoint = "https://api.twitter.com/2/media/upload/initialize"
)

func Initialize(ctx context.Context, c *gotwi.Client, p *types.InitializeInput) (*types.InitializeOutput, error) {
	res := &types.InitializeOutput{}
	if err := c.CallAPI(ctx, initializeEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
