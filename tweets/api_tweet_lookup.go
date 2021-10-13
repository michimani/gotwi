package tweets

import (
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweets/types"
)

const (
	TweetLookupTweetsEndpoint   = "https://api.twitter.com/2/tweets"
	TweetLookupTweetsIDEndpoint = "https://api.twitter.com/2/tweets/:id"
)

// Returns a variety of information about the Tweet specified by the requested ID or list of IDs.
// https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference/get-tweets
func TweetLookupTweets(c *gotwi.TwitterClient, p *types.TweetLookupTweetsParams) (*types.TweetLookupTweetsResponse, error) {
	res := &types.TweetLookupTweetsResponse{}
	if err := c.CallAPI(TweetLookupTweetsEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Returns a variety of information about a single Tweet specified by the requested ID.
// https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference/get-tweets-id
func TweetLookupTweetsID(c *gotwi.TwitterClient, p *types.TweetLookupTweetsIDParams) (*types.TweetLookupTweetsIDResponse, error) {
	res := &types.TweetLookupTweetsIDResponse{}
	if err := c.CallAPI(TweetLookupTweetsIDEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
