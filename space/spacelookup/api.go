package spacelookup

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/space/spacelookup/types"
)

const (
	getEndpoint              = "https://api.twitter.com/2/spaces/:id"
	listEndpoint             = "https://api.twitter.com/2/spaces"
	listByCreatorIDsEndpoint = "https://api.twitter.com/2/spaces/by/creator_ids"
	listBuyersEndpoint       = "https://api.twitter.com/2/spaces/:id/buyers"
	listTweetsEndpoint       = "https://api.twitter.com/2/spaces/:id/tweets"
)

// Returns a variety of information about a single Space specified by the requested ID.
// https://developer.twitter.com/en/docs/twitter-api/spaces/lookup/api-reference/get-spaces-id
func Get(ctx context.Context, c gotwi.IClient, p *types.GetInput) (*types.GetOutput, error) {
	res := &types.GetOutput{}
	if err := c.CallAPI(ctx, getEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Returns details about multiple Spaces. Up to 100 comma-separated Spaces IDs can be looked up using this endpoint
// https://developer.twitter.com/en/docs/twitter-api/spaces/lookup/api-reference/get-spaces
func List(ctx context.Context, c gotwi.IClient, p *types.ListInput) (*types.ListOutput, error) {
	res := &types.ListOutput{}
	if err := c.CallAPI(ctx, listEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Returns live or scheduled Spaces created by the specified user IDs.
// Up to 100 comma-separated IDs can be looked up using this endpoint.
// https://developer.twitter.com/en/docs/twitter-api/spaces/lookup/api-reference/get-spaces-by-creator-ids
func ListByCreatorIDs(ctx context.Context, c gotwi.IClient, p *types.ListByCreatorIDsInput) (*types.ListByCreatorIDsOutput, error) {
	res := &types.ListByCreatorIDsOutput{}
	if err := c.CallAPI(ctx, listByCreatorIDsEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Returns a list of user who purchased a ticket to the requested Space.
// You must authenticate the request using the access token of the creator of the requested Space.
// https://developer.twitter.com/en/docs/twitter-api/spaces/lookup/api-reference/get-spaces-id-buyers
func ListBuyers(ctx context.Context, c gotwi.IClient, p *types.ListBuyersInput) (*types.ListBuyersOutput, error) {
	res := &types.ListBuyersOutput{}
	if err := c.CallAPI(ctx, listBuyersEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Returns Tweets shared in the requested Spaces.
// https://developer.twitter.com/en/docs/twitter-api/spaces/lookup/api-reference/get-spaces-id-tweets
func ListTweets(ctx context.Context, c gotwi.IClient, p *types.ListTweetsInput) (*types.ListTweetsOutput, error) {
	res := &types.ListTweetsOutput{}
	if err := c.CallAPI(ctx, listTweetsEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
