package users

import (
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/users/types"
)

const (
	UserLookupUsersEndpoint           = "https://api.twitter.com/2/users"
	UserLookupUsersIDEndpoint         = "https://api.twitter.com/2/users/:id"
	UserLookupUsersByEndpoint         = "https://api.twitter.com/2/users/by"
	UserLookupUsersByUsernameEndpoint = "https://api.twitter.com/2/users/by/username/:username"
)

// Returns a variety of information about one or more users specified by the requested IDs.
// https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users
func UserLookupUsers(c *gotwi.TwitterClient, p *types.UserLookupUsersParams) (*types.UserLookupUsersResponse, error) {
	res := &types.UserLookupUsersResponse{}
	if err := c.CallAPI(UserLookupUsersEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Returns a variety of information about a single user specified by the requested ID.
// https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-id
func UserLookupUsersID(c *gotwi.TwitterClient, p *types.UserLookupUsersIDParams) (*types.UserLookupUsersIDResponse, error) {
	res := &types.UserLookupUsersIDResponse{}
	if err := c.CallAPI(UserLookupUsersIDEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Returns a variety of information about one or more users specified by their usernames.
// https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-by
func UserLookupUsersBy(c *gotwi.TwitterClient, p *types.UserLookupUsersByParams) (*types.UserLookupUsersByResponse, error) {
	res := &types.UserLookupUsersByResponse{}
	if err := c.CallAPI(UserLookupUsersByEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Returns a variety of information about one or more users specified by their usernames.
// https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-by-username-username
func UserLookupUsersByUsername(c *gotwi.TwitterClient, p *types.UserLookupUsersByUsernameParams) (*types.UserLookupUsersByUsernameResponse, error) {
	res := &types.UserLookupUsersByUsernameResponse{}
	if err := c.CallAPI(UserLookupUsersByUsernameEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
