package users

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/users/types"
)

const (
	FollowsFollowingGetEndpoint    = "https://api.twitter.com/2/users/:id/following"
	FollowsFollowersEndpoint       = "https://api.twitter.com/2/users/:id/followers"
	FollowsFollowingPostEndpoint   = "https://api.twitter.com/2/users/:id/following"
	FollowsFollowingDeleteEndpoint = "https://api.twitter.com/2/users/:source_user_id/following/:target_user_id"
)

// Returns a list of users the specified user ID is following.
// https://developer.twitter.com/en/docs/twitter-api/users/follows/api-reference/get-users-id-following
func FollowsFollowingGet(ctx context.Context, c *gotwi.GotwiClient, p *types.FollowsFollowingGetParams) (*types.FollowsFollowingGetResponse, error) {
	res := &types.FollowsFollowingGetResponse{}
	if err := c.CallAPI(ctx, FollowsFollowingGetEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Returns a list of users who are followers of the specified user ID.
// https://developer.twitter.com/en/docs/twitter-api/users/follows/api-reference/get-users-id-followers
func FollowsFollowers(ctx context.Context, c *gotwi.GotwiClient, p *types.FollowsFollowersParams) (*types.FollowsFollowersResponse, error) {
	res := &types.FollowsFollowersResponse{}
	if err := c.CallAPI(ctx, FollowsFollowersEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Allows a user ID to follow another user.
// If the target user does not have public Tweets, this endpoint will send a follow request.
// The request succeeds with no action when the authenticated user sends a request to a user
// they're already following, or if they're sending a follower request to a user that does not have public Tweets.
// https://developer.twitter.com/en/docs/twitter-api/users/follows/api-reference/post-users-source_user_id-following
func FollowsFollowingPost(ctx context.Context, c *gotwi.GotwiClient, p *types.FollowsFollowingPostParams) (*types.FollowsFollowingPostResponse, error) {
	res := &types.FollowsFollowingPostResponse{}
	if err := c.CallAPI(ctx, FollowsFollowingPostEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Allows a user ID to unfollow another user.
// The request succeeds with no action when the authenticated user sends a request to a user they're not following or have already unfollowed.
// https://developer.twitter.com/en/docs/twitter-api/users/follows/api-reference/delete-users-source_id-following
func FollowsFollowingDelete(ctx context.Context, c *gotwi.GotwiClient, p *types.FollowsFollowingDeleteParams) (*types.FollowsFollowingDeleteResponse, error) {
	res := &types.FollowsFollowingDeleteResponse{}
	if err := c.CallAPI(ctx, FollowsFollowingDeleteEndpoint, "DELETE", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
