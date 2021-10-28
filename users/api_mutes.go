package users

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/users/types"
)

const (
	MutesMutingGetEndpoint    = "https://api.twitter.com/2/users/:id/muting"
	MutesMutingPostEndpoint   = "https://api.twitter.com/2/users/:id/muting"
	MutesMutingDeleteEndpoint = "https://api.twitter.com/2/users/:source_user_id/muting/:target_user_id"
)

// Returns a list of users who are muted by the specified user ID.
// https://developer.twitter.com/en/docs/twitter-api/users/mutes/api-reference/get-users-muting
func MutesMutingGet(ctx context.Context, c *gotwi.GotwiClient, p *types.MutesMutingGetParams) (*types.MutesMutingGetResponse, error) {
	res := &types.MutesMutingGetResponse{}
	if err := c.CallAPI(ctx, MutesMutingGetEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Allows an authenticated user ID to mute the target user.
// https://developer.twitter.com/en/docs/twitter-api/users/mutes/api-reference/post-users-user_id-muting
func MutesMutingPost(ctx context.Context, c *gotwi.GotwiClient, p *types.MutesMutingPostParams) (*types.MutesMutingPostResponse, error) {
	res := &types.MutesMutingPostResponse{}
	if err := c.CallAPI(ctx, MutesMutingPostEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Allows an authenticated user ID to unmute the target user.
// The request succeeds with no action when the user sends a request to a user they're not muting or have already unmuted.
// https://developer.twitter.com/en/docs/twitter-api/users/mutes/api-reference/delete-users-user_id-muting
func MutesMutingDelete(ctx context.Context, c *gotwi.GotwiClient, p *types.MutesMutingDeleteParams) (*types.MutesMutingDeleteResponse, error) {
	res := &types.MutesMutingDeleteResponse{}
	if err := c.CallAPI(ctx, MutesMutingDeleteEndpoint, "DELETE", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
