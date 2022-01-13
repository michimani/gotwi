package users

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/users/types"
)

const (
	UserLookupEndpoint           = "https://api.twitter.com/2/users"
	UserLookupIDEndpoint         = "https://api.twitter.com/2/users/:id"
	UserLookupByEndpoint         = "https://api.twitter.com/2/users/by"
	UserLookupByUsernameEndpoint = "https://api.twitter.com/2/users/by/username/:username"
	UserLookupMeEndpoint         = "https://api.twitter.com/2/users/me"
)

// Returns a variety of information about one or more users specified by the requested IDs.
// https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users
func UserLookup(ctx context.Context, c *gotwi.GotwiClient, p *types.UserLookupParams) (*types.UserLookupResponse, error) {
	res := &types.UserLookupResponse{}
	if err := c.CallAPI(ctx, UserLookupEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Returns a variety of information about a single user specified by the requested ID.
// https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-id
func UserLookupID(ctx context.Context, c *gotwi.GotwiClient, p *types.UserLookupIDParams) (*types.UserLookupIDResponse, error) {
	res := &types.UserLookupIDResponse{}
	if err := c.CallAPI(ctx, UserLookupIDEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Returns a variety of information about one or more users specified by their usernames.
// https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-by
func UserLookupBy(ctx context.Context, c *gotwi.GotwiClient, p *types.UserLookupByParams) (*types.UserLookupByResponse, error) {
	res := &types.UserLookupByResponse{}
	if err := c.CallAPI(ctx, UserLookupByEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Returns a variety of information about one or more users specified by their usernames.
// https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-by-username-username
func UserLookupByUsername(ctx context.Context, c *gotwi.GotwiClient, p *types.UserLookupByUsernameParams) (*types.UserLookupByUsernameResponse, error) {
	res := &types.UserLookupByUsernameResponse{}
	if err := c.CallAPI(ctx, UserLookupByUsernameEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Returns information about an authorized user.
// https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-me
func UserLookupMe(ctx context.Context, c *gotwi.GotwiClient, p *types.UserLookupMeParams) (*types.UserLookupMeResponse, error) {
	res := &types.UserLookupMeResponse{}
	if err := c.CallAPI(ctx, UserLookupMeEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
