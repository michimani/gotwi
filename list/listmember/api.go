package listmember

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/list/listmember/types"
)

const (
	listMembershipsEndpoint = "https://api.twitter.com/2/users/:id/list_memberships"
	listEndpoint            = "https://api.twitter.com/2/lists/:id/members"
	createEndpoint          = "https://api.twitter.com/2/lists/:id/members"
	deleteEndpoint          = "https://api.twitter.com/2/lists/:id/members/:user_id"
)

// Returns all Lists a specified user is a member of.
// https://developer.twitter.com/en/docs/twitter-api/lists/list-members/api-reference/get-users-id-list_memberships
func ListMemberships(ctx context.Context, c *gotwi.Client, p *types.ListMembershipsInput) (*types.ListMembershipsOutput, error) {
	res := &types.ListMembershipsOutput{}
	if err := c.CallAPI(ctx, listMembershipsEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Returns a list of users who are members of the specified List.
// https://developer.twitter.com/en/docs/twitter-api/lists/list-members/api-reference/get-lists-id-members
func List(ctx context.Context, c *gotwi.Client, p *types.ListInput) (*types.ListOutput, error) {
	res := &types.ListOutput{}
	if err := c.CallAPI(ctx, listEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Enables the authenticated user to add a member to a List they own.
// https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/post-lists-id-members
func Create(ctx context.Context, c *gotwi.Client, p *types.CreateInput) (*types.CreateOutput, error) {
	res := &types.CreateOutput{}
	if err := c.CallAPI(ctx, createEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Enables the authenticated user to remove a member from a List they own.
// https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/delete-lists-id-members-user_id
func Delete(ctx context.Context, c *gotwi.Client, p *types.DeleteInput) (*types.DeleteOutput, error) {
	res := &types.DeleteOutput{}
	if err := c.CallAPI(ctx, deleteEndpoint, "DELETE", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
