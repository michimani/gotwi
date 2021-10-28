package tweets

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweets/types"
)

const (
	TweetRetweetsRetweetedByEndpoint = "https://api.twitter.com/2/tweets/:id/retweeted_by"
	TweetRetweetsPostEndpoint        = "https://api.twitter.com/2/users/:id/retweets"
	TweetRetweetsDeleteEndpoint      = "https://api.twitter.com/2/users/:id/retweets/:source_tweet_id"
)

// Allows you to get information about who has Retweeted a Tweet.
// https://developer.twitter.com/en/docs/twitter-api/tweets/retweets/api-reference/get-tweets-id-retweeted_by
func TweetRetweetsRetweetedBy(ctx context.Context, c *gotwi.GotwiClient, p *types.TweetRetweetsRetweetedByParams) (*types.TweetRetweetsRetweetedByResponse, error) {
	res := &types.TweetRetweetsRetweetedByResponse{}
	if err := c.CallAPI(ctx, TweetRetweetsRetweetedByEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Causes the user ID identified in the path parameter to Retweet the target Tweet.
// https://developer.twitter.com/en/docs/twitter-api/tweets/retweets/api-reference/post-users-id-retweets
func TweetRetweetsPost(ctx context.Context, c *gotwi.GotwiClient, p *types.TweetRetweetsPostParams) (*types.TweetRetweetsPostResponse, error) {
	res := &types.TweetRetweetsPostResponse{}
	if err := c.CallAPI(ctx, TweetRetweetsPostEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Allows a user or authenticated user ID to remove the Retweet of a Tweet.
// The request succeeds with no action when the user sends a request to a user
// they're not Retweeting the Tweet or have already removed the Retweet of.
// https://developer.twitter.com/en/docs/twitter-api/tweets/retweets/api-reference/delete-users-id-retweets-tweet_id
func TweetRetweetsDelete(ctx context.Context, c *gotwi.GotwiClient, p *types.TweetRetweetsDeleteParams) (*types.TweetRetweetsDeleteResponse, error) {
	res := &types.TweetRetweetsDeleteResponse{}
	if err := c.CallAPI(ctx, TweetRetweetsDeleteEndpoint, "DELETE", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
