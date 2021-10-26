package lists

import (
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/lists/types"
)

const (
	ManageListsPostEndpoint = "https://api.twitter.com/2/lists"
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
