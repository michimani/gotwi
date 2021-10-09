package tweets

import (
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweets/types"
)

const (
	TweetLookupTweetsEndpoint = "https://api.twitter.com/2/tweets"
)

func TweetLookupTweets(c *gotwi.TwitterClient, p *types.TweetLookupTweetsParams) (*types.TweetLookupTweetsResponse, error) {
	res := &types.TweetLookupTweetsResponse{}
	if err := c.CallAPI(TweetLookupTweetsEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
