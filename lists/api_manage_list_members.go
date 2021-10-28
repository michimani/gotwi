package lists

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/lists/types"
)

const (
	ManageListMembersPostEndpoint   = "https://api.twitter.com/2/lists/:id/members"
	ManageListMembersDeleteEndpoint = "https://api.twitter.com/2/lists/:id/members/:user_id"
)

// Enables the authenticated user to add a member to a List they own.
// https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/post-lists-id-members
func ManageListMembersPost(ctx context.Context, c *gotwi.GotwiClient, p *types.ManageListMembersPostParams) (*types.ManageListMembersPostResponse, error) {
	res := &types.ManageListMembersPostResponse{}
	if err := c.CallAPI(ctx, ManageListMembersPostEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Enables the authenticated user to remove a member from a List they own.
// https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/delete-lists-id-members-user_id
func ManageListMembersDelete(ctx context.Context, c *gotwi.GotwiClient, p *types.ManageListMembersDeleteParams) (*types.ManageListMembersDeleteResponse, error) {
	res := &types.ManageListMembersDeleteResponse{}
	if err := c.CallAPI(ctx, ManageListMembersDeleteEndpoint, "DELETE", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
