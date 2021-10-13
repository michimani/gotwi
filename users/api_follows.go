package users

import (
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/users/types"
)

const (
	FollowsFollowingGetEndpoint = "https://api.twitter.com/2/users/:id/following"
	FollowsFollowersEndpoint    = "https://api.twitter.com/2/users/:id/followers"
)

// Returns a list of users the specified user ID is following.
// https://developer.twitter.com/en/docs/twitter-api/users/follows/api-reference/get-users-id-following
func FollowsFollowingGet(c *gotwi.TwitterClient, p *types.FollowsFollowingGetParams) (*types.FollowsFollowingGetResponse, error) {
	res := &types.FollowsFollowingGetResponse{}
	if err := c.CallAPI(FollowsFollowingGetEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Returns a list of users who are followers of the specified user ID.
// https://developer.twitter.com/en/docs/twitter-api/users/follows/api-reference/get-users-id-followers
func FollowsFollowers(c *gotwi.TwitterClient, p *types.FollowsFollowersParams) (*types.FollowsFollowersResponse, error) {
	res := &types.FollowsFollowersResponse{}
	if err := c.CallAPI(FollowsFollowersEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
