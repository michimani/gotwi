package tweets

import (
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweets/types"
)

const (
	SearchTweetsTweetsSearchRecentEndpoint = "https://api.twitter.com/2/tweets/search/recent"
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
