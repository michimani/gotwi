package lists

import (
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/lists/types"
)

const (
	ManageListFollowsPostEndpoint   = "https://api.twitter.com/2/users/:id/followed_lists"
	ManageListFollowsDeleteEndpoint = "https://api.twitter.com/2/users/:id/followed_lists/:list_id"
)

// Enables the authenticated user to follow a List.
// https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/post-users-id-followed-lists
func ManageListFollowsPost(c *gotwi.GotwiClient, p *types.ManageListFollowsPostParams) (*types.ManageListFollowsPostResponse, error) {
	res := &types.ManageListFollowsPostResponse{}
	if err := c.CallAPI(ManageListFollowsPostEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Enables the authenticated user to unfollow a List.
// https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/delete-users-id-followed-lists-list_id
func ManageListFollowsDelete(c *gotwi.GotwiClient, p *types.ManageListFollowsDeleteParams) (*types.ManageListFollowsDeleteResponse, error) {
	res := &types.ManageListFollowsDeleteResponse{}
	if err := c.CallAPI(ManageListFollowsDeleteEndpoint, "DELETE", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
