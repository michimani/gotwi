package users

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/users/types"
)

const (
	BlocksBlockingGetEndpoint    = "https://api.twitter.com/2/users/:id/blocking"
	BlocksBlockingPostEndpoint   = "https://api.twitter.com/2/users/:id/blocking"
	BlocksBlockingDeleteEndpoint = "https://api.twitter.com/2/users/:source_user_id/blocking/:target_user_id"
)

// Returns a list of users who are blocked by the specified user ID.
// https://developer.twitter.com/en/docs/twitter-api/users/blocks/api-reference/get-users-blocking
func BlocksBlockingGet(ctx context.Context, c *gotwi.GotwiClient, p *types.BlocksBlockingGetParams) (*types.BlocksBlockingGetResponse, error) {
	res := &types.BlocksBlockingGetResponse{}
	if err := c.CallAPI(ctx, BlocksBlockingGetEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Causes the user (in the path) to block the target user. The user (in the path) must match the user context authorizing the request.
// https://developer.twitter.com/en/docs/twitter-api/users/blocks/api-reference/post-users-user_id-blocking
func BlocksBlockingPost(ctx context.Context, c *gotwi.GotwiClient, p *types.BlocksBlockingPostParams) (*types.BlocksBlockingPostResponse, error) {
	res := &types.BlocksBlockingPostResponse{}
	if err := c.CallAPI(ctx, BlocksBlockingPostEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Allows a user or authenticated user ID to unblock another user.
// The request succeeds with no action when the user sends a request to a user they're not blocking or have already unblocked.
// https://developer.twitter.com/en/docs/twitter-api/users/blocks/api-reference/delete-users-user_id-blocking
func BlocksBlockingDelete(ctx context.Context, c *gotwi.GotwiClient, p *types.BlocksBlockingDeleteParams) (*types.BlocksBlockingDeleteResponse, error) {
	res := &types.BlocksBlockingDeleteResponse{}
	if err := c.CallAPI(ctx, BlocksBlockingDeleteEndpoint, "DELETE", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
