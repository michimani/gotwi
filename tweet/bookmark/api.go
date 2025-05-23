package bookmark

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/bookmark/types"
)

const (
	listEndpoint   = "https://api.twitter.com/2/users/:id/bookmarks"
	createEndpoint = "https://api.twitter.com/2/users/:id/bookmarks"
	deleteEndpoint = "https://api.twitter.com/2/users/:id/bookmarks/:tweet_id"
)

// Allows you to get information about a authenticated user’s 800 most recent bookmarked Tweets
// https://developer.twitter.com/en/docs/twitter-api/tweets/bookmarks/api-reference/get-users-id-bookmarks
func List(ctx context.Context, c gotwi.IClient, p *types.ListInput) (*types.ListOutput, error) {
	res := &types.ListOutput{}
	if err := c.CallAPI(ctx, listEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Causes the user ID of an authenticated user identified in the path parameter
// to Bookmark the target Tweet provided in the request body.
// https://developer.twitter.com/en/docs/twitter-api/tweets/bookmarks/api-reference/post-users-id-bookmarks
func Create(ctx context.Context, c gotwi.IClient, p *types.CreateInput) (*types.CreateOutput, error) {
	res := &types.CreateOutput{}
	if err := c.CallAPI(ctx, createEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Allows a user or authenticated user ID to remove a Bookmark of a Tweet.
// https://developer.twitter.com/en/docs/twitter-api/tweets/bookmarks/api-reference/delete-users-id-bookmarks-tweet_id
func Delete(ctx context.Context, c gotwi.IClient, p *types.DeleteInput) (*types.DeleteOutput, error) {
	res := &types.DeleteOutput{}
	if err := c.CallAPI(ctx, deleteEndpoint, "DELETE", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
