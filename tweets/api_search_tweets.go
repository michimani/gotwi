package tweets

import (
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweets/types"
)

const (
	SearchTweetsTweetsSearchRecentEndpoint = "https://api.twitter.com/2/tweets/search/recent"
)

func SearchTweetsTweetsSearchRecent(c *gotwi.TwitterClient, p *types.SearchTweetsTweetsSearchRecentParams) (*types.SearchTweetsTweetsSearchRecentResponse, error) {
	res := &types.SearchTweetsTweetsSearchRecentResponse{}
	if err := c.CallAPI(SearchTweetsTweetsSearchRecentEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
