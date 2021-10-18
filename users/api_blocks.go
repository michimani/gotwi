package users

import (
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/users/types"
)

const (
	BlocksBlockingGetEndpoint = "https://api.twitter.com/2/users/:id/blocking"
)

func BlocksBlockingGet(c *gotwi.GotwiClient, p *types.BlocksBlockingGetParams) (*types.BlocksBlockingGetResponse, error) {
	res := &types.BlocksBlockingGetResponse{}
	if err := c.CallAPI(BlocksBlockingGetEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
