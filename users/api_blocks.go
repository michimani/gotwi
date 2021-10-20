package users

import (
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/users/types"
)

const (
	BlocksBlockingGetEndpoint  = "https://api.twitter.com/2/users/:id/blocking"
	BlocksBlockingPostEndpoint = "https://api.twitter.com/2/users/:id/blocking"
)

// Returns a list of users who are blocked by the specified user ID.
// https://developer.twitter.com/en/docs/twitter-api/users/blocks/api-reference/get-users-blocking
func BlocksBlockingGet(c *gotwi.GotwiClient, p *types.BlocksBlockingGetParams) (*types.BlocksBlockingGetResponse, error) {
	res := &types.BlocksBlockingGetResponse{}
	if err := c.CallAPI(BlocksBlockingGetEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Causes the user (in the path) to block the target user. The user (in the path) must match the user context authorizing the request.
// https://developer.twitter.com/en/docs/twitter-api/users/blocks/api-reference/post-users-user_id-blocking
func BlocksBlockingPost(c *gotwi.GotwiClient, p *types.BlocksBlockingPostParams) (*types.BlocksBlockingPostResponse, error) {
	res := &types.BlocksBlockingPostResponse{}
	if err := c.CallAPI(BlocksBlockingPostEndpoint, "Post", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
