package lists

import (
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/lists/types"
)

const (
	ManagePinnedListPostEndpoint   = "https://api.twitter.com/2/users/:id/pinned_lists"
	ManagePinnedListDeleteEndpoint = "https://api.twitter.com/2/users/:id/pinned_lists/:list_id"
)

// Enables the authenticated user to pin a List.
// https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/post-users-id-pinned-lists
func ManagePinnedListPost(c *gotwi.GotwiClient, p *types.ManagePinnedListPostParams) (*types.ManagePinnedListPostResponse, error) {
	res := &types.ManagePinnedListPostResponse{}
	if err := c.CallAPI(ManagePinnedListPostEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Enables the authenticated user to unpin a List.
// https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/delete-users-id-pinned-lists-list_id
func ManagePinnedListDelete(c *gotwi.GotwiClient, p *types.ManagePinnedListDeleteParams) (*types.ManagePinnedListDeleteResponse, error) {
	res := &types.ManagePinnedListDeleteResponse{}
	if err := c.CallAPI(ManagePinnedListDeleteEndpoint, "DELETE", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
