package users

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/users/types"
)

const (
	listUsersEndpoint            = "https://api.twitter.com/2/users"
	getUserEndpoint              = "https://api.twitter.com/2/users/:id"
	listUsersByUsernamesEndpoint = "https://api.twitter.com/2/users/by"
	getUserByUsernameEndpoint    = "https://api.twitter.com/2/users/by/username/:username"
	getMeEndpoint                = "https://api.twitter.com/2/users/me"
)

// GET /2/users
// Returns a variety of information about one or more users specified by the requested IDs.
// https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users
func ListUsers(ctx context.Context, c *gotwi.GotwiClient, p *types.ListUsersInput) (*types.ListUsersOutput, error) {
	res := &types.ListUsersOutput{}
	if err := c.CallAPI(ctx, listUsersEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GET /2/users/:id
// Returns a variety of information about a single user specified by the requested ID.
// https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-id
func GetUser(ctx context.Context, c *gotwi.GotwiClient, p *types.GetUserInput) (*types.GetUserOutput, error) {
	res := &types.GetUserOutput{}
	if err := c.CallAPI(ctx, getUserEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GET /2/users/by
// Returns a variety of information about one or more users specified by their usernames.
// https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-by
func ListUsersByUsernames(ctx context.Context, c *gotwi.GotwiClient, p *types.ListUsersByUsernamesInput) (*types.ListUsersByUsernamesOutput, error) {
	res := &types.ListUsersByUsernamesOutput{}
	if err := c.CallAPI(ctx, listUsersByUsernamesEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GET /2/users/by/username/:username
// Returns a variety of information about a single user specified by their usernames.
// https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-by-username-username
func GetUserByUsername(ctx context.Context, c *gotwi.GotwiClient, p *types.GetUserByUsernameInput) (*types.GetUserByUsernameOutput, error) {
	res := &types.GetUserByUsernameOutput{}
	if err := c.CallAPI(ctx, getUserByUsernameEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GET /2/users/me
// Returns information about an authorized user.
// https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-me
func GetMe(ctx context.Context, c *gotwi.GotwiClient, p *types.GetMeInput) (*types.GetMeOutput, error) {
	res := &types.GetMeOutput{}
	if err := c.CallAPI(ctx, getMeEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
