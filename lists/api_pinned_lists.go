package lists

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/lists/types"
)

const (
	PinnedListsGetEndpoint    = "https://api.twitter.com/2/users/:id/pinned_lists"
	PinnedListsPostEndpoint   = "https://api.twitter.com/2/users/:id/pinned_lists"
	PinnedListsDeleteEndpoint = "https://api.twitter.com/2/users/:id/pinned_lists/:list_id"
)

// Returns the Lists pinned by a specified user.
// https://developer.twitter.com/en/docs/twitter-api/lists/pinned-lists/api-reference/get-users-id-pinned_lists
func PinnedListsGet(ctx context.Context, c *gotwi.GotwiClient, p *types.PinnedListsGetParams) (*types.PinnedListsGetResponse, error) {
	res := &types.PinnedListsGetResponse{}
	if err := c.CallAPI(ctx, PinnedListsGetEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Enables the authenticated user to pin a List.
// https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/post-users-id-pinned-lists
func PinnedListsPost(ctx context.Context, c *gotwi.GotwiClient, p *types.PinnedListsPostParams) (*types.PinnedListsPostResponse, error) {
	res := &types.PinnedListsPostResponse{}
	if err := c.CallAPI(ctx, PinnedListsPostEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Enables the authenticated user to unpin a List.
// https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/delete-users-id-pinned-lists-list_id
func PinnedListsDelete(ctx context.Context, c *gotwi.GotwiClient, p *types.PinnedListsDeleteParams) (*types.PinnedListsDeleteResponse, error) {
	res := &types.PinnedListsDeleteResponse{}
	if err := c.CallAPI(ctx, PinnedListsDeleteEndpoint, "DELETE", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
