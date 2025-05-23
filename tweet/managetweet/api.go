package managetweet

import (
	"context"
	"errors"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/managetweet/types"
)

const (
	createEndpoint = "https://api.twitter.com/2/tweets"
	deleteEndpoint = "https://api.twitter.com/2/tweets/:id"
)

// Creates a Tweet on behalf of an authenticated user.
// https://developer.twitter.com/en/docs/twitter-api/tweets/manage-tweets/api-reference/post-tweets
func Create(ctx context.Context, c gotwi.IClient, p *types.CreateInput) (*types.CreateOutput, error) {
	if p == nil {
		return nil, errors.New("parameters is required")
	}

	res := &types.CreateOutput{}
	if err := c.CallAPI(ctx, createEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Allows a user or authenticated user ID to delete a Tweet.
// https://developer.twitter.com/en/docs/twitter-api/tweets/manage-tweets/api-reference/delete-tweets-id
func Delete(ctx context.Context, c gotwi.IClient, p *types.DeleteInput) (*types.DeleteOutput, error) {
	if p == nil {
		return nil, errors.New("parameters is required")
	}

	res := &types.DeleteOutput{}
	if err := c.CallAPI(ctx, deleteEndpoint, "DELETE", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
