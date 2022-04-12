package tweetlookup

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/tweetlookup/types"
)

const (
	listEndpoint = "https://api.twitter.com/2/tweets"
	getEndpoint  = "https://api.twitter.com/2/tweets/:id"
)

// Returns a variety of information about the Tweet specified by the requested ID or list of IDs.
// https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference/get-tweets
func List(ctx context.Context, c *gotwi.Client, p *types.ListInput) (*types.ListOutput, error) {
	res := &types.ListOutput{}
	if err := c.CallAPI(ctx, listEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Returns a variety of information about a single Tweet specified by the requested ID.
// https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference/get-tweets-id
func Get(ctx context.Context, c *gotwi.Client, p *types.GetInput) (*types.GetOutput, error) {
	res := &types.GetOutput{}
	if err := c.CallAPI(ctx, getEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
