package lists

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/lists/types"
)

const (
	ListMembersListMembershipsEndpoint = "https://api.twitter.com/2/users/:id/list_memberships"
	ListMembersGetEndpoint             = "https://api.twitter.com/2/lists/:id/members"
	ListMembersPostEndpoint            = "https://api.twitter.com/2/lists/:id/members"
	ListMembersDeleteEndpoint          = "https://api.twitter.com/2/lists/:id/members/:user_id"
)

// Returns all Lists a specified user is a member of.
// https://developer.twitter.com/en/docs/twitter-api/lists/list-members/api-reference/get-users-id-list_memberships
func ListMembersListMemberships(ctx context.Context, c *gotwi.GotwiClient, p *types.ListMembersListMembershipsParams) (*types.ListMembersListMembershipsResponse, error) {
	res := &types.ListMembersListMembershipsResponse{}
	if err := c.CallAPI(ctx, ListMembersListMembershipsEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Returns a list of users who are members of the specified List.
// https://developer.twitter.com/en/docs/twitter-api/lists/list-members/api-reference/get-lists-id-members
func ListMembersGet(ctx context.Context, c *gotwi.GotwiClient, p *types.ListMembersGetParams) (*types.ListMembersGetResponse, error) {
	res := &types.ListMembersGetResponse{}
	if err := c.CallAPI(ctx, ListMembersGetEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Enables the authenticated user to add a member to a List they own.
// https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/post-lists-id-members
func ListMembersPost(ctx context.Context, c *gotwi.GotwiClient, p *types.ListMembersPostParams) (*types.ListMembersPostResponse, error) {
	res := &types.ListMembersPostResponse{}
	if err := c.CallAPI(ctx, ListMembersPostEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Enables the authenticated user to remove a member from a List they own.
// https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/delete-lists-id-members-user_id
func ListMembersDelete(ctx context.Context, c *gotwi.GotwiClient, p *types.ListMembersDeleteParams) (*types.ListMembersDeleteResponse, error) {
	res := &types.ListMembersDeleteResponse{}
	if err := c.CallAPI(ctx, ListMembersDeleteEndpoint, "DELETE", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
