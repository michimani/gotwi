package lists

import (
	"context"
	"errors"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/list/listlookup/types"
)

const (
	getEndpoint       = "https://api.twitter.com/2/lists/:id"
	listOwnedEndpoint = "https://api.twitter.com/2/users/:id/owned_lists"
)

// Returns the details of a specified List.
// https://developer.twitter.com/en/docs/twitter-api/lists/list-lookup/api-reference/get-lists-id
func Get(ctx context.Context, c gotwi.IClient, p *types.GetInput) (*types.GetOutput, error) {
	if p == nil {
		return nil, errors.New("GetInput is nil")
	}
	res := &types.GetOutput{}
	if err := c.CallAPI(ctx, getEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Returns all Lists owned by the specified user.
// https://developer.twitter.com/en/docs/twitter-api/lists/list-lookup/api-reference/get-users-id-owned_lists
func ListOwned(ctx context.Context, c gotwi.IClient, p *types.ListOwnedInput) (*types.ListOwnedOutput, error) {
	if p == nil {
		return nil, errors.New("ListOwnedInput is nil")
	}
	res := &types.ListOwnedOutput{}
	if err := c.CallAPI(ctx, listOwnedEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
