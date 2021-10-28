package lists

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/lists/types"
)

const (
	ManageListsPostEndpoint   = "https://api.twitter.com/2/lists"
	ManageListsPutEndpoint    = "https://api.twitter.com/2/lists/:id"
	ManageListsDeleteEndpoint = "https://api.twitter.com/2/lists/:id"
)

// Enables the authenticated user to create a List.
// https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/post-lists
func ManageListsPost(ctx context.Context, c *gotwi.GotwiClient, p *types.ManageListsPostParams) (*types.ManageListsPostResponse, error) {
	res := &types.ManageListsPostResponse{}
	if err := c.CallAPI(ctx, ManageListsPostEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Enables the authenticated user to update the meta data of a specified List that they own.
// https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/put-lists-id
func ManageListsPut(ctx context.Context, c *gotwi.GotwiClient, p *types.ManageListsPutParams) (*types.ManageListsPutResponse, error) {
	res := &types.ManageListsPutResponse{}
	if err := c.CallAPI(ctx, ManageListsPutEndpoint, "PUT", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Enables the authenticated user to delete a List that they own.
// https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/delete-lists-id
func ManageListsDelete(ctx context.Context, c *gotwi.GotwiClient, p *types.ManageListsDeleteParams) (*types.ManageListsDeleteResponse, error) {
	res := &types.ManageListsDeleteResponse{}
	if err := c.CallAPI(ctx, ManageListsDeleteEndpoint, "DELETE", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
