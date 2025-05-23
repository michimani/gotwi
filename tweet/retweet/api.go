package retweet

import (
	"context"
	"errors"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/retweet/types"
)

const (
	listUsersEndpoint = "https://api.twitter.com/2/tweets/:id/retweeted_by"
	createEndpoint    = "https://api.twitter.com/2/users/:id/retweets"
	deleteEndpoint    = "https://api.twitter.com/2/users/:id/retweets/:source_tweet_id"
)

// Allows you to get information about who has Retweeted a Tweet.
// https://developer.twitter.com/en/docs/twitter-api/tweets/retweets/api-reference/get-tweets-id-retweeted_by
func ListUsers(ctx context.Context, c gotwi.IClient, p *types.ListUsersInput) (*types.ListUsersOutput, error) {
	if p == nil {
		return nil, errors.New("ListUsersInput is nil")
	}
	res := &types.ListUsersOutput{}
	if err := c.CallAPI(ctx, listUsersEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Causes the user ID identified in the path parameter to Retweet the target Tweet.
// https://developer.twitter.com/en/docs/twitter-api/tweets/retweets/api-reference/post-users-id-retweets
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

// Allows a user or authenticated user ID to remove the Retweet of a Tweet.
// The request succeeds with no action when the user sends a request to a user
// they're not Retweeting the Tweet or have already removed the Retweet of.
// https://developer.twitter.com/en/docs/twitter-api/tweets/retweets/api-reference/delete-users-id-retweets-tweet_id
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
