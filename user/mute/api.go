package mute

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/user/mute/types"
)

const (
	listEndpoint   = "https://api.twitter.com/2/users/:id/muting"
	createEndpoint = "https://api.twitter.com/2/users/:id/muting"
	deleteEndpoint = "https://api.twitter.com/2/users/:source_user_id/muting/:target_user_id"
)

// Returns a list of users who are muted by the specified user ID.
// https://developer.twitter.com/en/docs/twitter-api/users/mutes/api-reference/get-users-muting
func Lists(ctx context.Context, c *gotwi.Client, p *types.ListsInput) (*types.ListsOutput, error) {
	res := &types.ListsOutput{}
	if err := c.CallAPI(ctx, listEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Allows an authenticated user ID to mute the target user.
// https://developer.twitter.com/en/docs/twitter-api/users/mutes/api-reference/post-users-user_id-muting
func Create(ctx context.Context, c *gotwi.Client, p *types.CreateInput) (*types.CreateOutput, error) {
	res := &types.CreateOutput{}
	if err := c.CallAPI(ctx, createEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Allows an authenticated user ID to unmute the target user.
// The request succeeds with no action when the user sends a request to a user they're not muting or have already unmuted.
// https://developer.twitter.com/en/docs/twitter-api/users/mutes/api-reference/delete-users-user_id-muting
func Delete(ctx context.Context, c *gotwi.Client, p *types.DeleteInput) (*types.DeleteOutput, error) {
	res := &types.DeleteOutput{}
	if err := c.CallAPI(ctx, deleteEndpoint, "DELETE", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
