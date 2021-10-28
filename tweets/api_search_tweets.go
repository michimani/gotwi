package tweets

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweets/types"
)

const (
	SearchTweetsRecentEndpoint = "https://api.twitter.com/2/tweets/search/recent"
	SearchTweetsAllEndpoint    = "https://api.twitter.com/2/tweets/search/all"
)

// The recent search endpoint returns Tweets from the last seven days that match a search query.
// https://developer.twitter.com/en/docs/twitter-api/tweets/search/api-reference/get-tweets-search-recent
func SearchTweetsRecent(ctx context.Context, c *gotwi.GotwiClient, p *types.SearchTweetsRecentParams) (*types.SearchTweetsRecentResponse, error) {
	res := &types.SearchTweetsRecentResponse{}
	if err := c.CallAPI(ctx, SearchTweetsRecentEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// This endpoint is only available to those users who have been approved for the Academic Research product track.
// The full-archive search endpoint returns the complete history of public Tweets matching a search query; since the first Tweet was created March 26, 2006.
// https://developer.twitter.com/en/docs/twitter-api/tweets/search/api-reference/get-tweets-search-all
func SearchTweetsAll(ctx context.Context, c *gotwi.GotwiClient, p *types.SearchTweetsAllParams) (*types.SearchTweetsAllResponse, error) {
	res := &types.SearchTweetsAllResponse{}
	if err := c.CallAPI(ctx, SearchTweetsAllEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
