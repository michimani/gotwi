package lists

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/lists/types"
)

const (
	ListLookupIDEndpoint         = "https://api.twitter.com/2/lists/:id"
	ListLookupOwnedListsEndpoint = "https://api.twitter.com/2/users/:id/owned_lists"
)

// Returns the details of a specified List.
// https://developer.twitter.com/en/docs/twitter-api/lists/list-lookup/api-reference/get-lists-id
func ListLookupID(ctx context.Context, c *gotwi.GotwiClient, p *types.ListLookupIDParams) (*types.ListLookupIDResponse, error) {
	res := &types.ListLookupIDResponse{}
	if err := c.CallAPI(ctx, ListLookupIDEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Returns all Lists owned by the specified user.
// https://developer.twitter.com/en/docs/twitter-api/lists/list-lookup/api-reference/get-users-id-owned_lists
func ListLookupOwnedLists(ctx context.Context, c *gotwi.GotwiClient, p *types.ListLookupOwnedListsParams) (*types.ListLookupOwnedListsResponse, error) {
	res := &types.ListLookupOwnedListsResponse{}
	if err := c.CallAPI(ctx, ListLookupOwnedListsEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
