package lists

import (
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
func ManageListsPost(c *gotwi.GotwiClient, p *types.ManageListsPostParams) (*types.ManageListsPostResponse, error) {
	res := &types.ManageListsPostResponse{}
	if err := c.CallAPI(ManageListsPostEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Enables the authenticated user to update the meta data of a specified List that they own.
// https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/put-lists-id
func ManageListsPut(c *gotwi.GotwiClient, p *types.ManageListsPutParams) (*types.ManageListsPutResponse, error) {
	res := &types.ManageListsPutResponse{}
	if err := c.CallAPI(ManageListsPutEndpoint, "PUT", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Enables the authenticated user to delete a List that they own.
// https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/delete-lists-id
func ManageListsDelete(c *gotwi.GotwiClient, p *types.ManageListsDeleteParams) (*types.ManageListsDeleteResponse, error) {
	res := &types.ManageListsDeleteResponse{}
	if err := c.CallAPI(ManageListsDeleteEndpoint, "DELETE", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
