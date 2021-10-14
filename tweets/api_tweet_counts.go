package tweets

import (
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweets/types"
)

const (
	TweetCountsTweetsCountsRecentEndpoint = "https://api.twitter.com/2/tweets/counts/recent"
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
