package like

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/like/types"
)

const (
	listUsersEndpoint = "https://api.twitter.com/2/tweets/:id/liking_users"
	listEndpoint      = "https://api.twitter.com/2/users/:id/liked_tweets"
	createEndpoint    = "https://api.twitter.com/2/users/:id/likes"
	deleteEndpoint    = "https://api.twitter.com/2/users/:id/likes/:tweet_id"
)

// Allows you to get information about a Tweet’s liking users.
// You will receive the most recent 100 users who liked the specified Tweet.
// https://developer.twitter.com/en/docs/twitter-api/tweets/likes/api-reference/get-tweets-id-liking_users
func ListUsers(ctx context.Context, c *gotwi.Client, p *types.ListUsersInput) (*types.ListUsersOutput, error) {
	res := &types.ListUsersOutput{}
	if err := c.CallAPI(ctx, listUsersEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Allows you to get information about a user’s liked Tweets.
// The Tweets returned by this endpoint count towards the Project-level Tweet cap.
// https://developer.twitter.com/en/docs/twitter-api/tweets/likes/api-reference/get-users-id-liked_tweets
func List(ctx context.Context, c *gotwi.Client, p *types.ListInput) (*types.ListOutput, error) {
	res := &types.ListOutput{}
	if err := c.CallAPI(ctx, listEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Causes the user ID identified in the path parameter to Like the target Tweet.
// https://developer.twitter.com/en/docs/twitter-api/tweets/likes/api-reference/post-users-id-likes
func Create(ctx context.Context, c *gotwi.Client, p *types.CreateInput) (*types.CreateOutput, error) {
	res := &types.CreateOutput{}
	if err := c.CallAPI(ctx, createEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Allows a user or authenticated user ID to unlike a Tweet.
// The request succeeds with no action when the user sends
//  a request to a user they're not liking the Tweet or have already unliked the Tweet.
// https://developer.twitter.com/en/docs/twitter-api/tweets/likes/api-reference/delete-users-id-likes-tweet_id
func Delete(ctx context.Context, c *gotwi.Client, p *types.DeleteInput) (*types.DeleteOutput, error) {
	res := &types.DeleteOutput{}
	if err := c.CallAPI(ctx, deleteEndpoint, "DELETE", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
