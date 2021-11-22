package lists

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/lists/types"
)

const (
	ListFollowsFollowersEndpoint     = "https://api.twitter.com/2/lists/:id/followers"
	ListFollowsFollowedListsEndpoint = "https://api.twitter.com/2/users/:id/followed_lists"
	ListFollowsPostEndpoint          = "https://api.twitter.com/2/users/:id/followed_lists"
	ListFollowsDeleteEndpoint        = "https://api.twitter.com/2/users/:id/followed_lists/:list_id"
)

// Returns a list of users who are followers of the specified List.
// https://developer.twitter.com/en/docs/twitter-api/lists/list-follows/api-reference/get-lists-id-followers
func ListFollowsFollowers(ctx context.Context, c *gotwi.GotwiClient, p *types.ListFollowsFollowersParams) (*types.ListFollowsFollowersResponse, error) {
	res := &types.ListFollowsFollowersResponse{}
	if err := c.CallAPI(ctx, ListFollowsFollowersEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Returns all Lists a specified user follows.
// https://developer.twitter.com/en/docs/twitter-api/lists/list-follows/api-reference/get-users-id-followed_lists
func ListFollowsFollowedLists(ctx context.Context, c *gotwi.GotwiClient, p *types.ListFollowsFollowedListsParams) (*types.ListFollowsFollowedListsResponse, error) {
	res := &types.ListFollowsFollowedListsResponse{}
	if err := c.CallAPI(ctx, ListFollowsFollowedListsEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Enables the authenticated user to follow a List.
// https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/post-users-id-followed-lists
func ListFollowsPost(ctx context.Context, c *gotwi.GotwiClient, p *types.ListFollowsPostParams) (*types.ListFollowsPostResponse, error) {
	res := &types.ListFollowsPostResponse{}
	if err := c.CallAPI(ctx, ListFollowsPostEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Enables the authenticated user to unfollow a List.
// https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/delete-users-id-followed-lists-list_id
func ListFollowsDelete(ctx context.Context, c *gotwi.GotwiClient, p *types.ListFollowsDeleteParams) (*types.ListFollowsDeleteResponse, error) {
	res := &types.ListFollowsDeleteResponse{}
	if err := c.CallAPI(ctx, ListFollowsDeleteEndpoint, "DELETE", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
