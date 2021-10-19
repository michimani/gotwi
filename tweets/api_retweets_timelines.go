package tweets

import (
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweets/types"
)

const (
	TweetRetweetsRetweetedByEndpoint = "https://api.twitter.com/2/tweets/:id/retweeted_by"
)

// Allows you to get information about who has Retweeted a Tweet.
// https://developer.twitter.com/en/docs/twitter-api/tweets/retweets/api-reference/get-tweets-id-retweeted_by
func TweetRetweetsRetweetedBy(c *gotwi.GotwiClient, p *types.TweetRetweetsRetweetedByParams) (*types.TweetRetweetsRetweetedByResponse, error) {
	res := &types.TweetRetweetsRetweetedByResponse{}
	if err := c.CallAPI(TweetRetweetsRetweetedByEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
