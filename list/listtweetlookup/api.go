package listtweetlookup

import (
	"context"
	"errors"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/list/listtweetlookup/types"
)

const listEndpoint = "https://api.twitter.com/2/lists/:id/tweets"

// Returns a list of Tweets from the specified List.
// https://developer.twitter.com/en/docs/twitter-api/lists/list-tweets/api-reference/get-lists-id-tweets
func List(ctx context.Context, c gotwi.IClient, p *types.ListInput) (*types.ListOutput, error) {
	if p == nil {
		return nil, errors.New("ListInput is nil")
	}
	res := &types.ListOutput{}
	if err := c.CallAPI(ctx, listEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
