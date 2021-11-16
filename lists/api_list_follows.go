package lists

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/lists/types"
)

const (
	ListFollowsPostEndpoint   = "https://api.twitter.com/2/users/:id/followed_lists"
	ListFollowsDeleteEndpoint = "https://api.twitter.com/2/users/:id/followed_lists/:list_id"
)

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
