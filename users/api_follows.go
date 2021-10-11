package users

import (
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/users/types"
)

const (
	FollowsFollowingGetEndpoint = "https://api.twitter.com/2/users/:id/following"
	FollowsFollowersEndpoint    = "https://api.twitter.com/2/users/:id/followers"
)

func FollowsFollowingGet(c *gotwi.TwitterClient, p *types.FollowsFollowingGetParams) (*types.FollowsFollowingGetResponse, error) {
	res := &types.FollowsFollowingGetResponse{}
	if err := c.CallAPI(FollowsFollowingGetEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

func FollowsFollowers(c *gotwi.TwitterClient, p *types.FollowsFollowersParams) (*types.FollowsFollowersResponse, error) {
	res := &types.FollowsFollowersResponse{}
	if err := c.CallAPI(FollowsFollowersEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
