package tweets

import (
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweets/types"
)

const (
	TweetCountsTweetsCountsRecentEndpoint = "https://api.twitter.com/2/tweets/counts/recent"
	TweetCountsTweetsCountsAllEndpoint    = "https://api.twitter.com/2/tweets/counts/all"
)

// The recent Tweet counts endpoint returns count of Tweets from the last seven days that match a search query.
// https://developer.twitter.com/en/docs/twitter-api/tweets/counts/api-reference/get-tweets-counts-recent
func TweetCountsTweetsCountsRecent(c *gotwi.TwitterClient, p *types.TweetCountsTweetsCountsRecentParams) (*types.TweetCountsTweetsCountsRecentResponse, error) {
	res := &types.TweetCountsTweetsCountsRecentResponse{}
	if err := c.CallAPI(TweetCountsTweetsCountsRecentEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// This endpoint is only available to those users who have been approved for the Academic Research product track.
// The full-archive search endpoint returns the complete history of public Tweets matching a search query; since the first Tweet was created March 26, 2006.
// https://developer.twitter.com/en/docs/twitter-api/tweets/counts/api-reference/get-tweets-counts-all
func TweetCountsTweetsCountsAll(c *gotwi.TwitterClient, p *types.TweetCountsTweetsCountsAllParams) (*types.TweetCountsTweetsCountsAllResponse, error) {
	res := &types.TweetCountsTweetsCountsAllResponse{}
	if err := c.CallAPI(TweetCountsTweetsCountsAllEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
