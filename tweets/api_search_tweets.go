package tweets

import (
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweets/types"
)

const (
	SearchTweetsTweetsSearchRecentEndpoint = "https://api.twitter.com/2/tweets/search/recent"
	SearchTweetsTweetsSearchAllEndpoint    = "https://api.twitter.com/2/tweets/search/all"
)

// The recent search endpoint returns Tweets from the last seven days that match a search query.
// https://developer.twitter.com/en/docs/twitter-api/tweets/search/api-reference/get-tweets-search-recent
func SearchTweetsTweetsSearchRecent(c *gotwi.TwitterClient, p *types.SearchTweetsTweetsSearchRecentParams) (*types.SearchTweetsTweetsSearchRecentResponse, error) {
	res := &types.SearchTweetsTweetsSearchRecentResponse{}
	if err := c.CallAPI(SearchTweetsTweetsSearchRecentEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// This endpoint is only available to those users who have been approved for the Academic Research product track.
// The full-archive search endpoint returns the complete history of public Tweets matching a search query; since the first Tweet was created March 26, 2006.
// https://developer.twitter.com/en/docs/twitter-api/tweets/search/api-reference/get-tweets-search-all
func SearchTweetsTweetsSearchAll(c *gotwi.TwitterClient, p *types.SearchTweetsTweetsSearchAllParams) (*types.SearchTweetsTweetsSearchAllResponse, error) {
	res := &types.SearchTweetsTweetsSearchAllResponse{}
	if err := c.CallAPI(SearchTweetsTweetsSearchAllEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
