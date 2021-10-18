package users

import (
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/users/types"
)

const (
	UserLookupEndpoint           = "https://api.twitter.com/2/users"
	UserLookupIDEndpoint         = "https://api.twitter.com/2/users/:id"
	UserLookupByEndpoint         = "https://api.twitter.com/2/users/by"
	UserLookupByUsernameEndpoint = "https://api.twitter.com/2/users/by/username/:username"
)

// Returns a variety of information about one or more users specified by the requested IDs.
// https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users
func UserLookup(c *gotwi.GotwiClient, p *types.UserLookupParams) (*types.UserLookupResponse, error) {
	res := &types.UserLookupResponse{}
	if err := c.CallAPI(UserLookupEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Returns a variety of information about a single user specified by the requested ID.
// https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-id
func UserLookupID(c *gotwi.GotwiClient, p *types.UserLookupIDParams) (*types.UserLookupIDResponse, error) {
	res := &types.UserLookupIDResponse{}
	if err := c.CallAPI(UserLookupIDEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Returns a variety of information about one or more users specified by their usernames.
// https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-by
func UserLookupBy(c *gotwi.GotwiClient, p *types.UserLookupByParams) (*types.UserLookupByResponse, error) {
	res := &types.UserLookupByResponse{}
	if err := c.CallAPI(UserLookupByEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Returns a variety of information about one or more users specified by their usernames.
// https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-by-username-username
func UserLookupByUsername(c *gotwi.GotwiClient, p *types.UserLookupByUsernameParams) (*types.UserLookupByUsernameResponse, error) {
	res := &types.UserLookupByUsernameResponse{}
	if err := c.CallAPI(UserLookupByUsernameEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
