package users

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/users/types"
)

const (
	listMutedUsersEndpoint  = "https://api.twitter.com/2/users/:id/muting"
	createMutedUserEndpoint = "https://api.twitter.com/2/users/:id/muting"
	deleteMutedUserEndpoint = "https://api.twitter.com/2/users/:source_user_id/muting/:target_user_id"
)

// Returns a list of users who are muted by the specified user ID.
// https://developer.twitter.com/en/docs/twitter-api/users/mutes/api-reference/get-users-muting
func ListMutedUsers(ctx context.Context, c *gotwi.Client, p *types.ListMutedUsersInput) (*types.ListMutedUsersOutput, error) {
	res := &types.ListMutedUsersOutput{}
	if err := c.CallAPI(ctx, listMutedUsersEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Allows an authenticated user ID to mute the target user.
// https://developer.twitter.com/en/docs/twitter-api/users/mutes/api-reference/post-users-user_id-muting
func CreateMutedUser(ctx context.Context, c *gotwi.Client, p *types.CreateMutedUserInput) (*types.CreateMutedUserOutput, error) {
	res := &types.CreateMutedUserOutput{}
	if err := c.CallAPI(ctx, createMutedUserEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Allows an authenticated user ID to unmute the target user.
// The request succeeds with no action when the user sends a request to a user they're not muting or have already unmuted.
// https://developer.twitter.com/en/docs/twitter-api/users/mutes/api-reference/delete-users-user_id-muting
func DeleteMutedUser(ctx context.Context, c *gotwi.Client, p *types.DeleteMutedUserInput) (*types.DeleteMutedUserOutput, error) {
	res := &types.DeleteMutedUserOutput{}
	if err := c.CallAPI(ctx, deleteMutedUserEndpoint, "DELETE", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
