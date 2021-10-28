package tweets

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweets/types"
)

const (
	TweetLikesLikingUsersEndpoint = "https://api.twitter.com/2/tweets/:id/liking_users"
	TweetLikesLikedTweetsEndpoint = "https://api.twitter.com/2/users/:id/liked_tweets"
	TweetLikesPostEndpoint        = "https://api.twitter.com/2/users/:id/likes"
	TweetLikesDeleteEndpoint      = "https://api.twitter.com/2/users/:id/likes/:tweet_id"
)

// Allows you to get information about a Tweet’s liking users.
// You will receive the most recent 100 users who liked the specified Tweet.
// https://developer.twitter.com/en/docs/twitter-api/tweets/likes/api-reference/get-tweets-id-liking_users
func TweetLikesLikingUsers(ctx context.Context, c *gotwi.GotwiClient, p *types.TweetLikesLikingUsersParams) (*types.TweetLikesLikingUsersResponse, error) {
	res := &types.TweetLikesLikingUsersResponse{}
	if err := c.CallAPI(ctx, TweetLikesLikingUsersEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Allows you to get information about a user’s liked Tweets.
// The Tweets returned by this endpoint count towards the Project-level Tweet cap.
// https://developer.twitter.com/en/docs/twitter-api/tweets/likes/api-reference/get-users-id-liked_tweets
func TweetLikesLikedTweets(ctx context.Context, c *gotwi.GotwiClient, p *types.TweetLikesLikedTweetsParams) (*types.TweetLikesLikedTweetsResponse, error) {
	res := &types.TweetLikesLikedTweetsResponse{}
	if err := c.CallAPI(ctx, TweetLikesLikedTweetsEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Causes the user ID identified in the path parameter to Like the target Tweet.
// https://developer.twitter.com/en/docs/twitter-api/tweets/likes/api-reference/post-users-id-likes
func TweetLikesPost(ctx context.Context, c *gotwi.GotwiClient, p *types.TweetLikesPostParams) (*types.TweetLikesPostResponse, error) {
	res := &types.TweetLikesPostResponse{}
	if err := c.CallAPI(ctx, TweetLikesPostEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Allows a user or authenticated user ID to unlike a Tweet.
// The request succeeds with no action when the user sends
//  a request to a user they're not liking the Tweet or have already unliked the Tweet.
// https://developer.twitter.com/en/docs/twitter-api/tweets/likes/api-reference/delete-users-id-likes-tweet_id
func TweetLikesDelete(ctx context.Context, c *gotwi.GotwiClient, p *types.TweetLikesDeleteParams) (*types.TweetLikesDeleteResponse, error) {
	res := &types.TweetLikesDeleteResponse{}
	if err := c.CallAPI(ctx, TweetLikesDeleteEndpoint, "DELETE", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
