package tweets

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweets/types"
)

const (
	TweetCountsRecentEndpoint = "https://api.twitter.com/2/tweets/counts/recent"
	TweetCountsAllEndpoint    = "https://api.twitter.com/2/tweets/counts/all"
)

// The recent Tweet counts endpoint returns count of Tweets from the last seven days that match a search query.
// https://developer.twitter.com/en/docs/twitter-api/tweets/counts/api-reference/get-tweets-counts-recent
func TweetCountsRecent(ctx context.Context, c *gotwi.GotwiClient, p *types.TweetCountsRecentParams) (*types.TweetCountsRecentResponse, error) {
	res := &types.TweetCountsRecentResponse{}
	if err := c.CallAPI(ctx, TweetCountsRecentEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// This endpoint is only available to those users who have been approved for the Academic Research product track.
// The full-archive search endpoint returns the complete history of public Tweets matching a search query; since the first Tweet was created March 26, 2006.
// https://developer.twitter.com/en/docs/twitter-api/tweets/counts/api-reference/get-tweets-counts-all
func TweetCountsAll(ctx context.Context, c *gotwi.GotwiClient, p *types.TweetCountsAllParams) (*types.TweetCountsAllResponse, error) {
	res := &types.TweetCountsAllResponse{}
	if err := c.CallAPI(ctx, TweetCountsAllEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
