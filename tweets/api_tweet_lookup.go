package tweets

import (
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweets/types"
)

const (
	TweetLookupEndpoint   = "https://api.twitter.com/2/tweets"
	TweetLookupIDEndpoint = "https://api.twitter.com/2/tweets/:id"
)

// Returns a variety of information about the Tweet specified by the requested ID or list of IDs.
// https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference/get-tweets
func TweetLookup(c *gotwi.GotwiClient, p *types.TweetLookupParams) (*types.TweetLookupResponse, error) {
	res := &types.TweetLookupResponse{}
	if err := c.CallAPI(TweetLookupEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Returns a variety of information about a single Tweet specified by the requested ID.
// https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference/get-tweets-id
func TweetLookupID(c *gotwi.GotwiClient, p *types.TweetLookupIDParams) (*types.TweetLookupIDResponse, error) {
	res := &types.TweetLookupIDResponse{}
	if err := c.CallAPI(TweetLookupIDEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
