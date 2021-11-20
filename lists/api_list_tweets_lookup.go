package lists

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/lists/types"
)

const ListTweetsLookupEndpoint = "https://api.twitter.com/2/lists/:id/tweets"

// Returns a list of Tweets from the specified List.
// https://developer.twitter.com/en/docs/twitter-api/lists/list-tweets/api-reference/get-lists-id-tweets
func ListTweetsLookup(ctx context.Context, c *gotwi.GotwiClient, p *types.ListTweetsLookupParams) (*types.ListTweetsLookupResponse, error) {
	res := &types.ListTweetsLookupResponse{}
	if err := c.CallAPI(ctx, ListTweetsLookupEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
