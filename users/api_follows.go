package users

import (
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/users/types"
)

const (
	FollowsFollowingGetEndpoint = "https://api.twitter.com/2/users/:id/following"
)

func FollowsFollowingGet(c *gotwi.TwitterClient, p *types.FollowsFollowingGetParams) (*types.FollowsFollowingGetResponse, error) {
	res := &types.FollowsFollowingGetResponse{}
	if err := c.CallAPI(FollowsFollowingGetEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
