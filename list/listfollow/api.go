package listfollow

import (
	"context"
	"errors"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/list/listfollow/types"
)

const (
	listFollowersEndpoint = "https://api.twitter.com/2/lists/:id/followers"
	listFollowedEndpoint  = "https://api.twitter.com/2/users/:id/followed_lists"
	createEndpoint        = "https://api.twitter.com/2/users/:id/followed_lists"
	deleteEndpoint        = "https://api.twitter.com/2/users/:id/followed_lists/:list_id"
)

// Returns a list of users who are followers of the specified List.
// https://developer.twitter.com/en/docs/twitter-api/lists/list-follows/api-reference/get-lists-id-followers
func ListFollowers(ctx context.Context, c gotwi.IClient, p *types.ListFollowersInput) (*types.ListFollowersOutput, error) {
	if p == nil {
		return nil, errors.New("ListFollowersInput is nil")
	}
	res := &types.ListFollowersOutput{}
	if err := c.CallAPI(ctx, listFollowersEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Returns all Lists a specified user follows.
// https://developer.twitter.com/en/docs/twitter-api/lists/list-follows/api-reference/get-users-id-followed_lists
func ListFollowed(ctx context.Context, c gotwi.IClient, p *types.ListFollowedInput) (*types.ListFollowedOutput, error) {
	if p == nil {
		return nil, errors.New("ListFollowedInput is nil")
	}
	res := &types.ListFollowedOutput{}
	if err := c.CallAPI(ctx, listFollowedEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Enables the authenticated user to follow a List.
// https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/post-users-id-followed-lists
func Create(ctx context.Context, c gotwi.IClient, p *types.CreateInput) (*types.CreateOutput, error) {
	if p == nil {
		return nil, errors.New("CreateInput is nil")
	}
	res := &types.CreateOutput{}
	if err := c.CallAPI(ctx, createEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Enables the authenticated user to unfollow a List.
// https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/delete-users-id-followed-lists-list_id
func Delete(ctx context.Context, c gotwi.IClient, p *types.DeleteInput) (*types.DeleteOutput, error) {
	if p == nil {
		return nil, errors.New("DeleteInput is nil")
	}
	res := &types.DeleteOutput{}
	if err := c.CallAPI(ctx, deleteEndpoint, "DELETE", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
