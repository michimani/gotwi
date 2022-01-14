package spaces

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/spaces/types"
)

const (
	SpacesLookupIDEndpoint           = "https://api.twitter.com/2/spaces/:id"
	SpacesLookupEndpoint             = "https://api.twitter.com/2/spaces"
	SpacesLookupByCreatorIDsEndpoint = "https://api.twitter.com/2/spaces/by/creator_ids"
	SpacesLookupBuyersEndpoint       = "https://api.twitter.com/2/spaces/:id/buyers"
	SpacesLookupTweetsEndpoint       = "https://api.twitter.com/2/spaces/:id/tweets"
)

// Returns a variety of information about a single Space specified by the requested ID.
// https://developer.twitter.com/en/docs/twitter-api/spaces/lookup/api-reference/get-spaces-id
func SpacesLookupID(ctx context.Context, c *gotwi.GotwiClient, p *types.SpacesLookupIDParams) (*types.SpacesLookupIDResponse, error) {
	res := &types.SpacesLookupIDResponse{}
	if err := c.CallAPI(ctx, SpacesLookupIDEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Returns details about multiple Spaces. Up to 100 comma-separated Spaces IDs can be looked up using this endpoint
// https://developer.twitter.com/en/docs/twitter-api/spaces/lookup/api-reference/get-spaces
func SpacesLookup(ctx context.Context, c *gotwi.GotwiClient, p *types.SpacesLookupParams) (*types.SpacesLookupResponse, error) {
	res := &types.SpacesLookupResponse{}
	if err := c.CallAPI(ctx, SpacesLookupEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Returns live or scheduled Spaces created by the specified user IDs.
// Up to 100 comma-separated IDs can be looked up using this endpoint.
// https://developer.twitter.com/en/docs/twitter-api/spaces/lookup/api-reference/get-spaces-by-creator-ids
func SpacesLookupByCreatorIDs(ctx context.Context, c *gotwi.GotwiClient, p *types.SpacesLookupByCreatorIDsParams) (*types.SpacesLookupByCreatorIDsResponse, error) {
	res := &types.SpacesLookupByCreatorIDsResponse{}
	if err := c.CallAPI(ctx, SpacesLookupByCreatorIDsEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Returns a list of user who purchased a ticket to the requested Space.
// You must authenticate the request using the access token of the creator of the requested Space.
// https://developer.twitter.com/en/docs/twitter-api/spaces/lookup/api-reference/get-spaces-id-buyers
func SpacesLookupBuyers(ctx context.Context, c *gotwi.GotwiClient, p *types.SpacesLookupBuyersParams) (*types.SpacesLookupBuyersResponse, error) {
	res := &types.SpacesLookupBuyersResponse{}
	if err := c.CallAPI(ctx, SpacesLookupBuyersEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Returns Tweets shared in the requested Spaces.
// https://developer.twitter.com/en/docs/twitter-api/spaces/lookup/api-reference/get-spaces-id-tweets
func SpacesLookupTweets(ctx context.Context, c *gotwi.GotwiClient, p *types.SpacesLookupTweetsParams) (*types.SpacesLookupTweetsResponse, error) {
	res := &types.SpacesLookupTweetsResponse{}
	if err := c.CallAPI(ctx, SpacesLookupTweetsEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
