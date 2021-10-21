package tweets

import (
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweets/types"
)

const (
	TweetRetweetsRetweetedByEndpoint = "https://api.twitter.com/2/tweets/:id/retweeted_by"
	TweetRetweetsPostEndpoint        = "https://api.twitter.com/2/users/:id/retweets"
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

// Causes the user ID identified in the path parameter to Retweet the target Tweet.
// https://developer.twitter.com/en/docs/twitter-api/tweets/retweets/api-reference/post-users-id-retweets
func TweetRetweetsPost(c *gotwi.GotwiClient, p *types.TweetRetweetsPostParams) (*types.TweetRetweetsPostResponse, error) {
	res := &types.TweetRetweetsPostResponse{}
	if err := c.CallAPI(TweetRetweetsPostEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
