package tweets

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweets/types"
)

const (
	TweetTimelinesTweetsEndpoint   = "https://api.twitter.com/2/users/:id/tweets"
	TweetTimelinesMentionsEndpoint = "https://api.twitter.com/2/users/:id/mentions"
)

// Returns Tweets composed by a single user, specified by the requested user ID.
// By default, the most recent ten Tweets are returned per request. Using pagination, the most recent 3,200 Tweets can be retrieved.
// The Tweets returned by this endpoint count towards the Project-level Tweet cap.
// https://developer.twitter.com/en/docs/twitter-api/tweets/timelines/api-reference/get-users-id-tweets
func TweetTimelinesTweets(ctx context.Context, c *gotwi.GotwiClient, p *types.TweetTimelinesTweetsParams) (*types.TweetTimelinesTweetsResponse, error) {
	res := &types.TweetTimelinesTweetsResponse{}
	if err := c.CallAPI(ctx, TweetTimelinesTweetsEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Returns Tweets mentioning a single user specified by the requested user ID.
// By default, the most recent ten Tweets are returned per request. Using pagination, up to the most recent 800 Tweets can be retrieved.
// The Tweets returned by this endpoint count towards the Project-level Tweet cap.
// https://developer.twitter.com/en/docs/twitter-api/tweets/timelines/api-reference/get-users-id-mentions
func TweetTimelinesMentions(ctx context.Context, c *gotwi.GotwiClient, p *types.TweetTimelinesMentionsParams) (*types.TweetTimelinesMentionsResponse, error) {
	res := &types.TweetTimelinesMentionsResponse{}
	if err := c.CallAPI(ctx, TweetTimelinesMentionsEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
